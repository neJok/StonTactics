package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"stontactics/domain"
)

type ProfileController struct {
	ProfileUsecase domain.ProfileUsecase
}

// Fetch		godoc
// @Summary		Получить информацию о пользователе
// @Tags        User
// @Success     200  {object}  domain.Profile
// @Router      /api/profile [get]
// @Produce		json
// @Security 	Bearer
func (pc *ProfileController) Fetch(c *gin.Context) {
	userID := c.GetString("x-user-id")

	profile, err := pc.ProfileUsecase.GetProfileByID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}
