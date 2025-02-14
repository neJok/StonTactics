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
	"golang.org/x/crypto/bcrypt"
)

type SignUpController struct {
	SignUpUsecase domain.SignUpUsecase
	Env           *bootstrap.Env
}

// SignUp	godoc
// @Summary		Регистрация по почте и паролю
// @Tags        Signup
// @Router      /signup/register [post]
// @Success		200		{object}	domain.SuccessResponse
// @Failure		400		{object}	domain.ErrorResponse
// @Param       signUpRequest	body	domain.SignUpRequest	true	"sign up request"
// @Produce		json
func (sc *SignUpController) SignUp(c *gin.Context) {
	var signUpRequest domain.SignUpRequest
	err := c.ShouldBindBodyWith(&signUpRequest, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	email := strings.ToLower(signUpRequest.Email)
	_, err = sc.SignUpUsecase.GetUserByEmail(c, email)
	if err == nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "User already exist"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signUpRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	now := time.Now()

	lastRegisterCode, err := sc.SignUpUsecase.GetRegisterCode(c, email)
	if err == nil {
		codeWorkUntil := lastRegisterCode.CreatedAt.Add(time.Minute)
		if now.Before(codeWorkUntil) {
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Wait to create a new code"})
			return
		}
	}

	code := domain.RegisterCode{
		Email:     email,
		Password:  hashedPassword,
		Code:      random.RandRange(100000, 999999),
		CreatedAt: &now,
		Attempts:  10,
	}
	sc.SignUpUsecase.CreateRegisterCode(c, &code)

	data := make(map[string]interface{}, 0)
	subject := "Регистрация"

	data["Code"] = code.Code
	data["Subject"] = subject
	go mail.SendEmail(email, "letter", data, subject, sc.Env)

	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "the code has been sent by email"})
}

// ConfirmCode	godoc
// @Summary		Подтверждение почты
// @Tags        Signup
// @Router      /signup/confirm [post]
// @Success		200		{object}	domain.RefreshTokenResponse
// @Failure		400		{object}	domain.ErrorResponse
// @Param       codeRequest	body	domain.ConfirmRegisterCodeRequest	true	"code request"
// @Produce		json
func (sc *SignUpController) ConfirmSingUpCode(c *gin.Context) {
	var confirmCodeRequest domain.ConfirmRegisterCodeRequest
	err := c.ShouldBindBodyWith(&confirmCodeRequest, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	email := strings.ToLower(confirmCodeRequest.Email)
	_, err = sc.SignUpUsecase.GetUserByEmail(c, email)
	if err == nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "User already exist"})
		return
	}

	registerCode, err := sc.SignUpUsecase.GetRegisterCode(c, email)
	if err != nil || registerCode.Attempts == 0 {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "The code was not found"})
		return
	}

	now := time.Now()
	codeWorkUntil := registerCode.CreatedAt.Add(15 * time.Minute)
	if now.After(codeWorkUntil) {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "The code has expired"})
		return
	}

	if confirmCodeRequest.Code != registerCode.Code {
		sc.SignUpUsecase.IncAttemptsRegisterCode(c, email)
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "The code does not match"})
		return
	}

	userEntry := domain.User{
		ID:        "",
		Name:      confirmCodeRequest.Name,
		AvatarURL: "",
		Pro: domain.UserPro{
			Active: false,
			Until:  nil,
		},
		Auth: domain.UserAuth{
			Email: domain.EmailAuth{
				Email:    registerCode.Email,
				Password: registerCode.Password,
			},
			Google: domain.SocialAuth{},
			VK:     domain.SocialAuth{},
		},
		CreatedAt: &now,
	}

	userEntry.ID, err = sc.SignUpUsecase.CreateUser(c, &userEntry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	accessToken, err := sc.SignUpUsecase.CreateAccessToken(&userEntry, sc.Env.AccessTokenSecret, sc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := sc.SignUpUsecase.CreateRefreshToken(&userEntry, sc.Env.RefreshTokenSecret, sc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
