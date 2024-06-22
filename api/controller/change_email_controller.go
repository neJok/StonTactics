package controller

import (
	"net/http"
	"strings"
	"time"

	"github.com/neJok/StonTactics/bootstrap"
	"github.com/neJok/StonTactics/domain"
	"github.com/neJok/StonTactics/internal/mail"
	"github.com/neJok/StonTactics/internal/random"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type ChangeEmailController struct {
	ChangeEmailUsecase domain.ChangeEmailUsecase
	Env                *bootstrap.Env
}

// CreateChangeEmailCode	godoc
// @Summary	    Отправить код на новую почту
// @Tags        ChangeEmail
// @Router      /api/reset/email [post]
// @Success		200		{object}	domain.Account
// @Failure		400		{object}	domain.ErrorResponse
// @Param       createChangeEmailRequest	body	domain.ChangeEmailCreate	true	"create code request"
// @Produce		json
// @Security 	Bearer
func (cc *ChangeEmailController) CreateChangeEmailCode(c *gin.Context) {
	var createChangeEmailRequest domain.ChangeEmailCreate
	err := c.ShouldBindBodyWith(&createChangeEmailRequest, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	email := strings.ToLower(createChangeEmailRequest.Email)
	user, err := cc.ChangeEmailUsecase.GetUserByEmail(c, email)
	if err == nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "This mail already used"})
		return
	}

	userID := c.GetString("x-user-id")

	now := time.Now()
	lastCode, err := cc.ChangeEmailUsecase.GetCodeByID(c, userID)
	if err == nil {
		codeWorkUntil := lastCode.CreatedAt.Add(time.Minute)
		if now.Before(codeWorkUntil) {
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Wait to create a new code"})
			return
		}
	}

	code := domain.ChangeEmail{
		UserID:    userID,
		Email:     email,
		Code:      random.RandRange(100000, 999999),
		CreatedAt: &now,
		Attempts:  10,
	}
	cc.ChangeEmailUsecase.CreateCode(c, &code)

	data := make(map[string]interface{}, 0)
	subject := "Смена почты"

	data["Code"] = code.Code
	data["Subject"] = subject
	go mail.SendEmail(email, "letter", data, subject, cc.Env)

	response := domain.Account{
		ID:        userID,
		Name:      user.Name,
		Email:     user.Auth.Email.Email,
		AvatarURL: user.AvatarURL,
		Pro:       user.Pro,
		CreatedAt: user.CreatedAt,
		VK:        user.Auth.VK,
	}

	c.JSON(http.StatusOK, response)
}

// ConfirmResetCode	godoc
// @Summary		Подтверждение новой почты
// @Tags        ChangeEmail
// @Router      /api/reset/email/confirm [post]
// @Success		200		{object}	domain.SuccessResponse
// @Failure		400		{object}	domain.ErrorResponse
// @Param       codeRequest	body	domain.ChangeEmailConfirmRequest	true	"code request"
// @Produce		json
// @Security 	Bearer
func (cc *ChangeEmailController) ConfirmResetCode(c *gin.Context) {
	var confirmCodeRequest domain.ChangeEmailConfirmRequest
	err := c.ShouldBindBodyWith(&confirmCodeRequest, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	userID := c.GetString("x-user-id")

	code, err := cc.ChangeEmailUsecase.GetCodeByID(c, userID)
	if err != nil || code.Attempts == 0 {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "The code was not found"})
		return
	}

	now := time.Now()
	codeWorkUntil := code.CreatedAt.Add(15 * time.Minute)
	if now.After(codeWorkUntil) {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "The code has expired"})
		return
	}

	if confirmCodeRequest.Code != code.Code {
		cc.ChangeEmailUsecase.IncCodeAttempts(c, userID)
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "The code does not match"})
		return
	}

	cc.ChangeEmailUsecase.UpdateEmail(c, userID, code.Email)
	go cc.ChangeEmailUsecase.DeleteCodeByID(c, userID)

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Email has been successfully updated",
	})
}
