package controller

import (
	"net/http"
	"stontactics/bootstrap"
	"stontactics/domain"
	"stontactics/internal/mail"
	"stontactics/internal/random"
	"stontactics/internal/tokenutil"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
)

type ResetPassowrdController struct {
	ResetPassowrdUsecase domain.ResetPassowrdUsecase
	Env                  *bootstrap.Env
}

// FetchOne	godoc
// @Summary	    Отправить запрос на смену пароля
// @Tags        ResetPassword
// @Router      /reset/password [post]
// @Success		200		{object}	domain.SuccessResponse
// @Failure		400		{object}	domain.ErrorResponse
// @Param       createResetPasswordRequest	body	domain.ResetPassowrdCreate	true	"create code request"
// @Produce		json
func (rc *ResetPassowrdController) CreateResetPasswordCode(c *gin.Context) {
	var createResetPasswordRequest domain.ResetPassowrdCreate
	err := c.ShouldBindBodyWith(&createResetPasswordRequest, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	email := strings.ToLower(createResetPasswordRequest.Email)
	_, err = rc.ResetPassowrdUsecase.GetUserByEmail(c, email)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "User not found"})
		return
	}

	now := time.Now()

	lastCode, err := rc.ResetPassowrdUsecase.GetCodeByEmail(c, email)
	if err == nil {
		codeWorkUntil := lastCode.CreatedAt.Add(time.Minute)
		if now.Before(codeWorkUntil) {
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Wait to create a new code"})
			return
		}
	}

	code := domain.ResetPassword{
		Email:     email,
		Code:      random.RandRange(100000, 999999),
		Token:     tokenutil.GenerateSecureToken(16),
		CreatedAt: &now,
		Attempts:  10,
	}
	rc.ResetPassowrdUsecase.CreateCode(c, &code)

	data := make(map[string]interface{}, 0)
	subject := "Восстановление пароля Ston Tactics"

	data["Code"] = code.Code
	data["Subject"] = subject
	go mail.SendEmail(email, "registerLetter", data, subject, rc.Env)

	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "the code has been sent by email"})
}

// ConfirmCode	godoc
// @Summary		Подтверждение почты по коду
// @Tags        ResetPassword
// @Router      /reset/password/confirm [post]
// @Success		200		{object}	domain.ResetPasswordResponse
// @Failure		400		{object}	domain.ErrorResponse
// @Param       codeRequest	body	domain.ResetPasswordConfirmRequest	true	"code request"
// @Produce		json
func (rc *ResetPassowrdController) ConfirmResetCode(c *gin.Context) {
	var confirmCodeRequest domain.ResetPasswordConfirmRequest
	err := c.ShouldBindBodyWith(&confirmCodeRequest, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	email := strings.ToLower(confirmCodeRequest.Email)
	_, err = rc.ResetPassowrdUsecase.GetUserByEmail(c, email)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "User not found"})
		return
	}

	code, err := rc.ResetPassowrdUsecase.GetCodeByEmail(c, email)
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
		rc.ResetPassowrdUsecase.IncCodeAttempts(c, email)
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "The code does not match"})
		return
	}

	c.JSON(http.StatusOK, domain.ResetPasswordResponse{
		Token: code.Token,
	})
}

// ConfirmCode	godoc
// @Summary		Смена пароля
// @Tags        ResetPassword
// @Router      /reset/password [put]
// @Success		200		{object}	domain.SuccessResponse
// @Failure		400		{object}	domain.ErrorResponse
// @Param       tokenRequest	body	domain.ResetPasswordRequest	true	"token and name request"
// @Produce		json
func (rc *ResetPassowrdController) ResetPasswordToken(c *gin.Context) {
	var resetPassowordRequest domain.ResetPasswordRequest
	err := c.ShouldBindBodyWith(&resetPassowordRequest, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	code, err := rc.ResetPassowrdUsecase.GetCodeByToken(c, resetPassowordRequest.Token)
	if err != nil || resetPassowordRequest.Token != code.Token {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid token"})
		return
	}

	now := time.Now()
	codeWorkUntil := code.CreatedAt.Add(15 * time.Minute)
	if now.After(codeWorkUntil) {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "The token has expired"})
		return
	}

	user, err := rc.ResetPassowrdUsecase.GetUserByEmail(c, code.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "User not found"})
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(resetPassowordRequest.Password), bcrypt.DefaultCost)
	rc.ResetPassowrdUsecase.UpdatePassword(c, user.ID, password)

	go rc.ResetPassowrdUsecase.DeleteCodeByEmail(c, code.Email)

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Password updated",
	})
}
