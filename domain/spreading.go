package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionSpreadouts = "spreadouts"
)

type Spreading struct {
	ID       primitive.ObjectID       `bson:"_id" json:"id" binding:"-"`
	Name     string                   `bson:"name" form:"name" binding:"required,max=100,min=1" json:"name"`
	Elements []map[string]interface{} `bson:"elements" form:"elements" binding:"required" json:"elements"`
	MapName  string                   `bson:"map_name" form:"map_name" binding:"required,max=100,min=1" json:"map_name"`
	UserID   string                   `bson:"user_id" json:"-"`
}

type SpreadingRepository interface {
	Create(c context.Context, spreading *Spreading) error
	FetchMany(c context.Context, userID string, ids []primitive.ObjectID) ([]SpreadingResponse, error)
	FetchByID(c context.Context, id string) (Spreading, error)
	Update(c context.Context, id string, elements []map[string]interface{}, mapName string) error
	GetCount(c context.Context, userID string) int64
}

type SpreadingUsecase interface {
	Create(c context.Context, spreading *Spreading) error
	FetchMany(c context.Context, userID string, ids []primitive.ObjectID) ([]SpreadingResponse, error)
	FetchByID(c context.Context, id string) (Spreading, error)
	Update(c context.Context, id string, elements []map[string]interface{}, mapName string) error
	GetCount(c context.Context, userID string) int64
}

type SpreadingResponse struct {
	ID      primitive.ObjectID `bson:"_id" json:"id"`
	Name    string             `bson:"name" json:"name"`
	MapName string             `bson:"map_name" json:"map_name"`
}

type SpreadingCreateRequest struct {
	Name     string                   `bson:"name" form:"name" binding:"required,max=100,min=1" json:"name"`
	Elements []map[string]interface{} `bson:"elements" form:"elements" binding:"required" json:"elements"`
	MapName  string                   `bson:"map_name" form:"map_name" binding:"required,max=100,min=1" json:"map_name"`
}

type SpreadingUpdateRequest struct {
	Elements []map[string]interface{} `bson:"elements" form:"elements" binding:"required" json:"elements"`
	MapName string               `bson:"map_name" form:"map_name" binding:"required,max=100,min=1" json:"map_name"`
}
