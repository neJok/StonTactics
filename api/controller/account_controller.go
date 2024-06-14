package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stontactics/bootstrap"
	"stontactics/domain"
)

type AccountController struct {
	AccountUsecase domain.AccountUsecase
	Env            *bootstrap.Env
}

// DeleteAccount	godoc
// @Summary	    Удалить аккаунт
// @Tags        Account
// @Router      /api/account [delete]
// @Success		200		{object}	domain.SuccessResponse
// @Produce		json
// @Security 	Bearer
func (ac *AccountController) DeleteAccount(c *gin.Context) {
	userID := c.GetString("x-user-id")
	ac.AccountUsecase.DeleteByID(c, userID)
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Account deleted"})
}

// Fetch		godoc
// @Summary		Получить информацию о пользователе
// @Tags        Account
// @Success     200  {object}  domain.Account
// @Router      /api/account [get]
// @Produce		json
// @Security 	Bearer
func (pc *AccountController) Fetch(c *gin.Context) {
	userID := c.GetString("x-user-id")

	account, err := pc.AccountUsecase.GetByAccountByID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, account)
}
