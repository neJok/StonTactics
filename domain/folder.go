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
	FetchOneByID(c context.Context, userID string, folderID string) (Folder, error)
	DeleteOneByID(c context.Context, userID string, folderID string) error
	AddStrategies(c context.Context, userID string, folderID string, strategiesIDS []string) error
	RemoveStrategies(c context.Context, userID string, folderID string, strategiesIDS []string) error
	AddSpreadouts(c context.Context, userID string, folderID string, spreadoutsIDS []string) error
	RemoveSpreadouts(c context.Context, userID string, folderID string, spreadoutsIDS []string) error
}

type FolderUsecase interface {
	Create(c context.Context, Folder *Folder) error
	FetchByUserID(c context.Context, userID string) ([]Folder, error)
	FetchOneByID(c context.Context, userID string, folderID string) (Folder, error)
	DeleteOneByID(c context.Context, userID string, folderID string) error
	AddStrategies(c context.Context, userID string, folderID string, strategiesIDS []string) error
	RemoveStrategies(c context.Context, userID string, folderID string, strategiesIDS []string) error
	AddSpreadouts(c context.Context, userID string, folderID string, spreadoutsIDS []string) error
	RemoveSpreadouts(c context.Context, userID string, folderID string, spreadoutsIDS []string) error
}

type FolderCreateRequest struct {
	Name string `bson:"name" form:"name" binding:"required,max=100" json:"name"`
}

type FolderAddStrategiesRequest struct {
	FolderID      string   `json:"folder_id" binding:"required"`
	StrategiesIDS []string `json:"strategies_ids" binding:"required"`
}

type FolderRemoveStrategiesRequest struct {
	FolderID      string   `json:"folder_id" binding:"required"`
	StrategiesIDS []string `json:"strategies_ids" binding:"required"`
}

type FolderAddSpreadoutsRequest struct {
	FolderID      string   `json:"folder_id" binding:"required"`
	SpreadoutsIDS []string `json:"spreadouts_ids" binding:"required"`
}

type FolderRemoveSpreadoutsRequest struct {
	FolderID      string   `json:"folder_id" binding:"required"`
	SpreadoutsIDS []string `json:"spreadouts_ids" binding:"required"`
}
