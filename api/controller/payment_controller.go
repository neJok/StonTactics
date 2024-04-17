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

	prices := map[string]uint64 {
		"30": 9900,
		"90": 14900,
	}

	if paymentRequest.Days == "" {
		paymentRequest.Days = "30"
	} else if _, ok := prices[paymentRequest.Days]; !ok {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "The days parameter was incorrectly passed, possible values: 30, 90"})
		return
	}

	req := &tinkoff.InitRequest{
		Amount:      prices[paymentRequest.Days],
		OrderID:     uuid.New().String(),
		RedirectDueDate: tinkoff.Time(time.Now().Add(4 * 24 * time.Hour)),
		Receipt: &tinkoff.Receipt{
			Email: paymentRequest.Email,
			Items: []*tinkoff.ReceiptItem{
				{
					Price:    prices[paymentRequest.Days],
					Quantity: "1",
					Amount:   prices[paymentRequest.Days],
					Name:     "Подпсика Pro",
					Tax:      tinkoff.VAT20,
				},
			},
			Taxation: tinkoff.TaxationUSNIncome,
			Payments: &tinkoff.ReceiptPayments{
				Electronic: prices[paymentRequest.Days],
			},
		},
		Data: map[string]string{
			"days": paymentRequest.Days,
			"user_id": c.GetString("x-user-id"),
		},
	}
	res, err := pc.TinkoffClient.Init(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, domain.PaymentCreated{Url: res.PaymentURL})
}

func (pc *PaymentController) CallbackTinkoff(c *gin.Context) {
	notification, err := pc.TinkoffClient.ParseNotification(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	days, err := strconv.Atoi(notification.Data["days"])
    if err != nil {
        c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
    }
	until := time.Now().Add(time.Hour * 24 * time.Duration(days))
	pc.PaymentUsecase.ActivatePro(c, notification.Data["user_id"], &until)

	c.String(http.StatusOK, pc.TinkoffClient.GetNotificationSuccessResponse())
}