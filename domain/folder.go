package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionFolders = "folders"
)

type Folder struct {
	ID         primitive.ObjectID `bson:"_id" json:"id" binding:"-"`
	Name       string             `bson:"name" form:"name" binding:"required,max=100" json:"name"`
	Strategies []string           `bson:"strategies" form:"strategies" binding:"-" json:"strategies"`
	Spreadouts []string           `bson:"spreadouts" form:"spreadouts" binding:"-" json:"spreadouts"`
	UserID     string             `bson:"user_id" json:"-"`
}

type FolderRepository interface {
	Create(c context.Context, Folder *Folder) error
	FetchByUserID(c context.Context, userID string) ([]Folder, error)
	AddStrategy(c context.Context, userID string, folderID string, strategyID string) error
	AddSpreading(c context.Context, userID string, folderID string, spreadingID string) error
}

type FolderUsecase interface {
	Create(c context.Context, Folder *Folder) error
	FetchByUserID(c context.Context, userID string) ([]Folder, error)
	AddStrategy(c context.Context, userID string, folderID string, strategyID string) error
	AddSpreading(c context.Context, userID string, folderID string, spreadingID string) error
}

type FolderCreateRequest struct {
	Name string `bson:"name" form:"name" binding:"required,max=100" json:"name"`
}

type FolderAddStrategyRequest struct {
	FolderID   string `json:"folder_id" binding:"required"`
	StrategyID string `json:"strategy_id" binding:"required"`
}

type FolderAddSpreadingRequest struct {
	FolderID    string `json:"folder_id" binding:"required"`
	SpreadingID string `json:"spreading_id" binding:"required"`
}
