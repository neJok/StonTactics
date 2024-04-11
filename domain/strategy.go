package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionStrategies = "strategies"
)

type Strategy struct {
	ID      primitive.ObjectID     `bson:"_id" json:"id" binding:"-" form:"-"`
	Name    string                 `bson:"name" form:"name" binding:"required,max=100,min=1" json:"name"`
	Parts   map[string]interface{} `bson:"parts" form:"parts" binding:"required" json:"parts"`
	MapName string                 `bson:"map_name" form:"map_name" binding:"required,max=100,min=1" json:"map_name"`
	UserID  string                 `bson:"user_id" json:"-" form:"-"`
}

type StrategyRepository interface {
	Create(c context.Context, strategy *Strategy) error
	FetchMany(c context.Context, userID string, ids []primitive.ObjectID) ([]StrategyResponse, error)
	FetchByID(c context.Context, id string) (Strategy, error)
	Update(c context.Context, id string, parts map[string]interface{}, mapName string) error
	GetCount(c context.Context, userID string) int64
}

type StrategyUsecase interface {
	Create(c context.Context, strategy *Strategy) error
	FetchMany(c context.Context, userID string, ids []primitive.ObjectID) ([]StrategyResponse, error)
	FetchByID(c context.Context, id string) (Strategy, error)
	Update(c context.Context, id string, parts map[string]interface{}, mapName string) error
	GetCount(c context.Context, userID string) int64
}

type StrategyResponse struct {
	ID      primitive.ObjectID `bson:"_id" json:"id"`
	Name    string             `bson:"name" json:"name"`
	MapName string             `bson:"map_name" json:"map_name"`
}

type StrategyCreateRequest struct {
	Name    string                 `bson:"name" form:"name" binding:"required,max=100,min=1" json:"name"`
	Parts   map[string]interface{} `bson:"parts" form:"parts" binding:"required" json:"parts"`
	MapName string                 `bson:"map_name" form:"map_name" binding:"required,max=100,min=1" json:"map_name"`
}

type StrategyUpdateRequest struct {
	Parts map[string]interface{} `bson:"parts" form:"parts" binding:"required" json:"parts"`
	MapName string               `bson:"map_name" form:"map_name" binding:"required,max=100,min=1" json:"map_name"`
}
