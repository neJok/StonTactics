package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"stontactics/domain"
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

// AddStrategy		godoc
// @Summary		Добавить стратегию в папку
// @Tags        Folder
// @Router      /api/folder/strategy [put]
// @Success		200		{object}	domain.SuccessResponse
// @Failure		400		{object}	domain.ErrorResponse
// @Param       request	body		domain.FolderAddStrategyRequest	true	"request"
// @Produce		json
// @Security 	Bearer
func (fc *FolderController) AddStrategy(c *gin.Context) {
	userID := c.GetString("x-user-id")

	var request domain.FolderAddStrategyRequest
	err := c.ShouldBindBodyWith(&request, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	_, err = fc.StrategyUsecase.FetchByID(c, request.StrategyID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = fc.FolderUsecase.AddStrategy(c, userID, request.FolderID, request.StrategyID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Folder updated successfully",
	})
}

// AddSpreading	godoc
// @Summary		Добавить раскидку в папку
// @Tags        Folder
// @Router      /api/folder/spreading [put]
// @Success		200		{object}	domain.SuccessResponse
// @Failure		400		{object}	domain.ErrorResponse
// @Param       request	body		domain.FolderAddSpreadingRequest	true	"request"
// @Produce		json
// @Security 	Bearer
func (fc *FolderController) AddSpreading(c *gin.Context) {
	userID := c.GetString("x-user-id")

	var request domain.FolderAddSpreadingRequest
	err := c.ShouldBindBodyWith(&request, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	_, err = fc.SpreadingUsecase.FetchByID(c, request.SpreadingID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = fc.FolderUsecase.AddSpreading(c, userID, request.FolderID, request.SpreadingID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Folder updated successfully",
	})
}
