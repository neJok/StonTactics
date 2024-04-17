package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"

	"stontactics/domain"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/nikita-vanyasin/tinkoff"
)

type PaymentController struct {
	PaymentUsecase 	  domain.PaymentUsecase
	TinkoffClient	  *tinkoff.Client
}


// Create		godoc
// @Summary		Создать платеж
// @Tags        Payment
// @Router      /api/payment/create/tinkoff [post]
// @Success		201		{object}	domain.PaymentCreated
// @Failure		400		{object}	domain.ErrorResponse
// @Param       paymentInfo	body	domain.PaymentCreateRequest	true	"paymentInfo"
// @Produce		json
// @Security 	Bearer
func (pc *PaymentController) CreateTinkoff(c *gin.Context) {
	var paymentRequest domain.PaymentCreateRequest
	err := c.ShouldBindBodyWith(&paymentRequest, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	
	prices := map[int16]uint64 {
		30: 9900,
		90: 14900,
	}

	if _, ok := prices[paymentRequest.Days]; !ok {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "The days parameter was incorrectly passed, possible values: 30, 90"})
		return
	}

	req := &tinkoff.InitRequest{
		Amount:      prices[paymentRequest.Days],
		OrderID:     uuid.New().String(),
		Receipt: &tinkoff.Receipt{
			Email: paymentRequest.Email,
			Taxation: tinkoff.TaxationUSNIncome,
			Items: []*tinkoff.ReceiptItem{
				{
					Price:    prices[paymentRequest.Days],
					Quantity: "1",
					Amount:   prices[paymentRequest.Days],
					Name:     "Подпсика Pro",
					Tax:      tinkoff.VAT20,
				},
			},
		},
		Data: map[string]string{"": "",},
	}
	res, err := pc.TinkoffClient.Init(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	
	pc.PaymentUsecase.Create(c, &domain.Payment{
		PaymentID: res.PaymentID,
		UserID: c.GetString("x-user-id"),
		Days: paymentRequest.Days,
		Paid: false,
	})

	c.JSON(http.StatusCreated, domain.PaymentCreated{Url: res.PaymentURL})
}

func (pc *PaymentController) CallbackTinkoff(c *gin.Context) {
	notification, err := pc.TinkoffClient.ParseNotification(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if notification.Success && notification.Status == "CONFIRMED" {
		paymentID := strconv.Itoa(int(notification.PaymentID))
		payment, err := pc.PaymentUsecase.GetByID(c, paymentID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
			return
		}

		user, err := pc.PaymentUsecase.GetUser(c, payment.UserID)
		if !payment.Paid && err == nil {
			additional := time.Hour * 24 * time.Duration(payment.Days)
			
			var until time.Time
			if user.Pro.Active {
				until = user.Pro.Until.Add(additional)
			} else {
				until = time.Now().Add(additional)
			}

			pc.PaymentUsecase.ActivatePro(c, payment.UserID, &until)
			pc.PaymentUsecase.SetPaid(c, paymentID)
		}
	}

	c.String(http.StatusOK, pc.TinkoffClient.GetNotificationSuccessResponse())
}