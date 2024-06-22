package controller

import (
	"fmt"
	"net/http"
	"os"

	"image"
	"image/jpeg"
	"image/png"

	"github.com/neJok/StonTactics/bootstrap"
	"github.com/neJok/StonTactics/domain"
	"github.com/nickalie/go-webpbin"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/nfnt/resize"

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
// @Success     200  {object}  domain.SuccessResponse
// @Success     400  {object}  domain.ErrorResponse
// @Success     500  {object}  domain.ErrorResponse
// @Router      /api/account [put]
// @Produce		json
// @Security 	Bearer
func (ac *AccountController) Update(c *gin.Context) {
	userID := c.GetString("x-user-id")

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Error receiving the file"})
		return
	}
	defer file.Close()

	_, format, err := image.DecodeConfig(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if _, err := file.Seek(0, 0); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	newFilename := fmt.Sprintf("%s.webp", userID)
	savePath := filepath.Join("media/avatars", newFilename)

	if err := os.MkdirAll("media/avatars", os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if _, err := os.Stat(savePath); err == nil {
		if err := os.Remove(savePath); err != nil {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
			return
		}
	}

	var img image.Image
	switch format {
	case "jpg", "jpeg":
		img, err = jpeg.Decode(file)
	case "png":
		img, err = png.Decode(file)
	default:
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid file format"})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Error decoding the file"})
		return
	}

	resizedImg := resize.Resize(500, 500, img, resize.Lanczos3)

	f, err := os.Create(savePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if err := webpbin.Encode(f, resizedImg); err != nil {
		f.Close()
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if err := f.Close(); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	ac.AccountUsecase.UpdateByID(c, userID, bson.M{"avatar_url": savePath})

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: savePath,
	})
}
