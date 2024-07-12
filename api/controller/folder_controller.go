package controller

import (
	"github.com/neJok/StonTactics/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FolderController struct {
	FolderUsecase    domain.FolderUsecase
	StrategyUsecase  domain.StrategyUsecase
	SpreadingUsecase domain.SpreadingUsecase
}

// Create		godoc
// @Summary		Создать папку
// @Tags        Folder
// @Router      /api/folder [post]
// @Success		200		{object}	domain.SuccessResponse
// @Failure		400		{object}	domain.ErrorResponse
// @Param       folder	body		domain.FolderCreateRequest	true	"folder"
// @Produce		json
// @Security 	Bearer
func (fc *FolderController) Create(c *gin.Context) {
	var folderRequest domain.FolderCreateRequest

	err := c.ShouldBindBodyWith(&folderRequest, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	userID := c.GetString("x-user-id")
	folder := domain.Folder{
		ID:         primitive.NewObjectID(),
		Name:       folderRequest.Name,
		Strategies: make([]string, 0),
		Spreadouts: make([]string, 0),
		UserID:     userID,
	}

	err = fc.FolderUsecase.Create(c, &folder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Folder created successfully",
	})
}

// FetchAll		godoc
// @Summary		Получить все папки пользователя
// @Tags        Folder
// @Router      /api/folder [get]
// @Success		200		{array}		[]domain.Folder
// @Failure		400		{object}	domain.ErrorResponse
// @Produce		json
// @Security 	Bearer
func (fc *FolderController) FetchAll(c *gin.Context) {
	userID := c.GetString("x-user-id")

	folders, err := fc.FolderUsecase.FetchByUserID(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, folders)
}

// AddStrategies		godoc
// @Summary		Добавить стратегии в папку
// @Tags        Folder
// @Router      /api/folder/strategy [put]
// @Success		200		{object}	domain.SuccessResponse
// @Failure		400		{object}	domain.ErrorResponse
// @Param       request	body		domain.FolderAddStrategiesRequest	true	"request"
// @Produce		json
// @Security 	Bearer
func (fc *FolderController) AddStrategies(c *gin.Context) {
	userID := c.GetString("x-user-id")

	var request domain.FolderAddStrategiesRequest
	err := c.ShouldBindBodyWith(&request, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = fc.FolderUsecase.AddStrategies(c, userID, request.FolderID, request.StrategiesIDS)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Folder updated successfully",
	})
}

// RemoveStrategies		godoc
// @Summary		Удалить стратегии из папки
// @Tags        Folder
// @Router      /api/folder/strategy [delete]
// @Success		200		{object}	domain.SuccessResponse
// @Failure		400		{object}	domain.ErrorResponse
// @Param       request	body		domain.FolderRemoveStrategiesRequest	true	"request"
// @Produce		json
// @Security 	Bearer
func (fc *FolderController) RemoveStrategies(c *gin.Context) {
	userID := c.GetString("x-user-id")

	var request domain.FolderRemoveStrategiesRequest
	err := c.ShouldBindBodyWith(&request, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	_, err = fc.FolderUsecase.FetchOneByID(c, userID, request.FolderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = fc.FolderUsecase.RemoveStrategies(c, userID, request.FolderID, request.StrategiesIDS)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Folder updated successfully",
	})
}

// RemoveSpreadouts		godoc
// @Summary		Удалить раскидки из папки
// @Tags        Folder
// @Router      /api/folder/spreading [delete]
// @Success		200		{object}	domain.SuccessResponse
// @Failure		400		{object}	domain.ErrorResponse
// @Param       request	body		domain.FolderRemoveSpreadoutsRequest	true	"request"
// @Produce		json
// @Security 	Bearer
func (fc *FolderController) RemoveSpreadouts(c *gin.Context) {
	userID := c.GetString("x-user-id")

	var request domain.FolderRemoveSpreadoutsRequest
	err := c.ShouldBindBodyWith(&request, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	_, err = fc.FolderUsecase.FetchOneByID(c, userID, request.FolderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = fc.FolderUsecase.RemoveSpreadouts(c, userID, request.FolderID, request.SpreadoutsIDS)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Folder updated successfully",
	})
}

// AddSpreadouts godoc
// @Summary		Добавить раскидки в папку
// @Tags        Folder
// @Router      /api/folder/spreading [put]
// @Success		200		{object}	domain.SuccessResponse
// @Failure		400		{object}	domain.ErrorResponse
// @Param       request	body		domain.FolderAddSpreadoutsRequest	true	"request"
// @Produce		json
// @Security 	Bearer
func (fc *FolderController) AddSpreadouts(c *gin.Context) {
	userID := c.GetString("x-user-id")

	var request domain.FolderAddSpreadoutsRequest
	err := c.ShouldBindBodyWith(&request, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = fc.FolderUsecase.AddSpreadouts(c, userID, request.FolderID, request.SpreadoutsIDS)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Folder updated successfully",
	})
}

// DeleteFolder godoc
// @Summary		Удалить папку
// @Tags        Folder
// @Router      /api/folder/{id} [delete]
// @Success		200		{object}	domain.SuccessResponse
// @Failure		400		{object}	domain.ErrorResponse
// @Param       id	path	string	true	"id"
// @Produce		json
// @Security 	Bearer
func (fc *FolderController) DeleteFolder(c *gin.Context) {
	userID := c.GetString("x-user-id")
	folderID := c.Param("id")

	_, err := fc.FolderUsecase.FetchOneByID(c, userID, folderID)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "Folder not found"})
		return
	}

	err = fc.FolderUsecase.DeleteOneByID(c, userID, folderID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Folder deleted successfully",
	})
}
