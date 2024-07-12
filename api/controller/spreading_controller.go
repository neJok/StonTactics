package controller

import (
	"net/http"
	"strings"

	"github.com/neJok/StonTactics/domain"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SpreadingController struct {
	SpreadingUsecase domain.SpreadingUsecase
	AccountUsecase   domain.AccountUsecase
}

// Create		godoc
// @Summary		Создать раскидку
// @Tags        Spreading
// @Router      /api/spreading [post]
// @Success		201		{object}	domain.Spreading
// @Failure		400		{object}	domain.ErrorResponse
// @Param       spreading	body	domain.SpreadingCreateRequest	true	"spreading"
// @Produce		json
// @Security 	Bearer
func (sc *SpreadingController) Create(c *gin.Context) {
	var spreadingRequest domain.SpreadingCreateRequest
	err := c.ShouldBindBodyWith(&spreadingRequest, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	userID := c.GetString("x-user-id")
	user, err := sc.AccountUsecase.GetByAccountByID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	countDocs := sc.SpreadingUsecase.GetCount(c, userID)
	if !user.Pro.Active && countDocs >= 5 {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "You can't upload new spreading without the pro version"})
		return
	}

	spreading := domain.Spreading{
		ID:       primitive.NewObjectID(),
		Name:     spreadingRequest.Name,
		Elements: spreadingRequest.Elements,
		MapName:  spreadingRequest.MapName,
		UserID:   userID,
	}

	err = sc.SpreadingUsecase.Create(c, &spreading)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, spreading)
}

// FetchMany	godoc
// @Summary		Получить несколько раскидок
// @Tags        Spreading
// @Router      /api/spreading [get]
// @Success		200		{array}		[]domain.Spreading
// @Failure		400		{object}	domain.ErrorResponse
// @Param       ids	query	string	false	"ids"
// @Produce		json
// @Security 	Bearer
func (sc *SpreadingController) FetchMany(c *gin.Context) {
	idsStr := c.Query("ids")

	var ids []primitive.ObjectID
	if idsStr != "" {
		idStrings := strings.Split(idsStr, ",")
		for _, idStr := range idStrings {
			id, err := primitive.ObjectIDFromHex(idStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id provided"})
				return
			}
			ids = append(ids, id)
		}
	} else {
		ids = make([]primitive.ObjectID, 0)
	}

	userID := c.GetString("x-user-id")

	spreadouts, err := sc.SpreadingUsecase.FetchMany(c, userID, ids)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, spreadouts)
}

// FetchOne	godoc
// @Summary		Получить одну раскидку по айди
// @Tags        Spreading
// @Router      /api/spreading/{id} [get]
// @Success		200		{object}	domain.Spreading
// @Failure		400		{object}	domain.ErrorResponse
// @Param       id	path	string	true	"id"
// @Produce		json
// @Security 	Bearer
func (sc *SpreadingController) FetchOne(c *gin.Context) {
	id := c.Param("id")
	spreading, err := sc.SpreadingUsecase.FetchByID(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, spreading)
}

// Update	godoc
// @Summary		Обновить раскидку
// @Tags        Spreading
// @Router      /api/spreading/{id} [put]
// @Success		200		{object}	nil
// @Failure		400		{object}	domain.ErrorResponse
// @Param       id		path	string	true	"id"
// @Param       update	body	domain.SpreadingUpdateRequest	true	"update"
// @Produce		json
// @Security 	Bearer
func (sc *SpreadingController) Update(c *gin.Context) {
	var spreadingUpdateRequest domain.SpreadingUpdateRequest
	err := c.ShouldBindBodyWith(&spreadingUpdateRequest, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	id := c.Param("id")

	err = sc.SpreadingUsecase.Update(c, id, spreadingUpdateRequest.Elements, spreadingUpdateRequest.MapName)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}


// DeleteMany	godoc
// @Summary		Удалить раскидки
// @Tags        Spreading
// @Router      /api/spreading [delete]
// @Success		200		{object}	nil
// @Failure		400		{object}	domain.ErrorResponse
// @Param       delete	body	domain.SpreadoutsDeleteRequest	true	"spreadouts ids"
// @Produce		json
// @Security 	Bearer
func (sc *SpreadingController) DeleteMany(c *gin.Context) {
	userID := c.GetString("x-user-id")

	var spreadoutsDeleteRequest domain.SpreadoutsDeleteRequest
	err := c.ShouldBindBodyWith(&spreadoutsDeleteRequest, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = sc.SpreadingUsecase.DeleteByIDS(c, userID, spreadoutsDeleteRequest.SpreadoutsIDS)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}