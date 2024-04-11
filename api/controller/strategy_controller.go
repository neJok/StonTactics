package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"stontactics/domain"
	"strings"
)

type StrategyController struct {
	StrategyUsecase domain.StrategyUsecase
}

// Create		godoc
// @Summary		Создать стратегию
// @Tags        Strategy
// @Router      /api/strategy [post]
// @Success		201		{object}	domain.StrategyCreateRequest
// @Failure		400		{object}	domain.ErrorResponse
// @Param       strategy	body		domain.StrategyCreateRequest	true	"strategy"
// @Produce		json
// @Security 	Bearer
func (sc *StrategyController) Create(c *gin.Context) {
	var strategyRequest domain.StrategyCreateRequest
	err := c.ShouldBindBodyWith(&strategyRequest, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	userID := c.GetString("x-user-id")

	strategy := domain.Strategy{
		ID:      primitive.NewObjectID(),
		Name:    strategyRequest.Name,
		Parts:   strategyRequest.Parts,
		MapName: strategyRequest.MapName,
		UserID:  userID,
	}

	err = sc.StrategyUsecase.Create(c, &strategy)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, strategy)
}

// FetchMany	godoc
// @Summary		Получить несколько стратегий
// @Tags        Strategy
// @Router      /api/strategy [get]
// @Success		200		{array}		[]domain.Strategy
// @Failure		400		{object}	domain.ErrorResponse
// @Param       ids	query	string	false	"ids"
// @Produce		json
// @Security 	Bearer
func (sc *StrategyController) FetchMany(c *gin.Context) {
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

	strategies, err := sc.StrategyUsecase.FetchMany(c, userID, ids)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, strategies)
}

// FetchOne	godoc
// @Summary		Получить одну стратегию по айди
// @Tags        Strategy
// @Router      /api/strategy/{id} [get]
// @Success		200		{object}	domain.Strategy
// @Failure		400		{object}	domain.ErrorResponse
// @Param       id	path	string	true	"id"
// @Produce		json
// @Security 	Bearer
func (sc *StrategyController) FetchOne(c *gin.Context) {
	id := c.Param("id")
	strategy, err := sc.StrategyUsecase.FetchByID(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, strategy)
}

// Update	godoc
// @Summary		Обновить стратегию
// @Tags        Strategy
// @Router      /api/strategy/{id} [put]
// @Success		200		{object}	nil
// @Failure		400		{object}	domain.ErrorResponse
// @Param       id		path	string	true	"id"
// @Param       update	body	domain.StrategyUpdateRequest	true	"update"
// @Produce		json
// @Security 	Bearer
func (sc *StrategyController) Update(c *gin.Context) {
	var strategyUpdateRequest domain.StrategyUpdateRequest
	err := c.ShouldBindBodyWith(&strategyUpdateRequest, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	id := c.Param("id")

	err = sc.StrategyUsecase.Update(c, id, strategyUpdateRequest.Parts, strategyUpdateRequest.MapName)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}
