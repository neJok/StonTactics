package controller

import (
	"net/http"
	"os"

	"github.com/neJok/StonTactics/bootstrap"
	"github.com/neJok/StonTactics/domain"
	"github.com/neJok/StonTactics/internal/imageutil"
	"go.mongodb.org/mongo-driver/bson"

	"path/filepath"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
	AccountUsecase domain.AccountUsecase
	Env            *bootstrap.Env
}

// Delete	godoc
// @Summary	    Удалить аккаунт
// @Tags        Account
// @Router      /api/account [delete]
// @Success		200		{object}	domain.SuccessResponse
// @Produce		json
// @Security 	Bearer
func (ac *AccountController) Delete(c *gin.Context) {
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
func (ac *AccountController) Fetch(c *gin.Context) {
	userID := c.GetString("x-user-id")

	account, err := ac.AccountUsecase.GetByAccountByID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, account)
}

// Update		godoc
// @Summary		Обновить информацию о пользователе
// @Tags        Account
// @Accept multipart/form-data
// @Param file formData file true "Файл для загрузки"
// @Success     200  {object}  domain.Account
// @Success     400  {object}  domain.ErrorResponse
// @Success     500  {object}  domain.ErrorResponse
// @Router      /api/account [put]
// @Produce		json
// @Security 	Bearer
func (ac *AccountController) Update(c *gin.Context) {
	userID := c.GetString("x-user-id")

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Error receiving the file"})
		return
	}
	defer file.Close()

	if !imageutil.IsValidImageFile(header) {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid file format. Please upload an image file."})
		return
	}

	user, err := ac.AccountUsecase.GetByAccountByID(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "User not found"})
		return
	}

	newFilename := user.ID + filepath.Ext(header.Filename)
	savePath := filepath.Join("media/avatars", newFilename)

	if err := os.MkdirAll("media/avatars", os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if _, err := os.Stat(user.AvatarURL); err == nil {
		os.Remove(user.AvatarURL)
	}
	err = c.SaveUploadedFile(header, savePath)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to save file"})
		return
	}

	ac.AccountUsecase.UpdateByID(c, user.ID, bson.M{"avatar_url": savePath})
	user.AvatarURL = savePath

	c.JSON(http.StatusOK, user)
}
