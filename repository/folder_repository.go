package repository

import (
	"context"

	"github.com/neJok/StonTactics/domain"
	"github.com/neJok/StonTactics/mongo"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
)

type folderRepository struct {
	database   mongo.Database
	collection string
}

func NewFolderRepository(db mongo.Database, collection string) domain.FolderRepository {
	return &folderRepository{
		database:   db,
		collection: collection,
	}
}

func (sr *folderRepository) Create(c context.Context, folder *domain.Folder) error {
	collection := sr.database.Collection(sr.collection)

	_, err := collection.InsertOne(c, folder)

	return err
}

func (sr *folderRepository) FetchByUserID(c context.Context, userID string) ([]domain.Folder, error) {
	collection := sr.database.Collection(sr.collection)

	var folders []domain.Folder

	opts := options.Find().SetProjection(bson.D{{Key: "elements", Value: 0}})
	cursor, err := collection.Find(c, bson.M{"user_id": userID}, opts)
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &folders)
	if folders == nil {
		return []domain.Folder{}, err
	}

	return folders, err
}

func (sr *folderRepository) FetchOneByID(c context.Context, userID string, folderID string) (domain.Folder, error) {
	collection := sr.database.Collection(sr.collection)

	var folder domain.Folder
	
	folderIDHex, err := primitive.ObjectIDFromHex(folderID)
	if err != nil {
		return domain.Folder{}, err
	}

	opts := options.FindOne().SetProjection(bson.D{{Key: "elements", Value: 0}})
	err = collection.FindOne(c, bson.M{"_id": folderIDHex, "user_id": userID}, opts).Decode(&folder)
	if err != nil {
		return domain.Folder{}, err
	}

	return folder, nil
}

func (sr *folderRepository) AddStrategy(c context.Context, userID string, folderID string, strategyID string) error {
	collection := sr.database.Collection(sr.collection)

	folderIDHex, err := primitive.ObjectIDFromHex(folderID)
	if err != nil {
		return err
	}

	_, err = collection.UpdateOne(c, bson.M{"_id": folderIDHex, "user_id": userID}, bson.M{"$addToSet": bson.M{"strategies": strategyID}})
	return err
}

func (sr *folderRepository) AddSpreading(c context.Context, userID string, folderID string, spreadingID string) error {
	collection := sr.database.Collection(sr.collection)

	folderIDHex, err := primitive.ObjectIDFromHex(folderID)
	if err != nil {
		return err
	}

	_, err = collection.UpdateOne(c, bson.M{"_id": folderIDHex, "user_id": userID}, bson.M{"$addToSet": bson.M{"spreadouts": spreadingID}})
	return err
}

func (sr *folderRepository) RemoveStrategy(c context.Context, userID string, folderID string, strategyID string) error {
	collection := sr.database.Collection(sr.collection)

	folderIDHex, err := primitive.ObjectIDFromHex(folderID)
	if err != nil {
		return err
	}

	_, err = collection.UpdateOne(c, bson.M{"_id": folderIDHex, "user_id": userID}, bson.M{"$pull": bson.M{"strategies": strategyID}})
	return err
}

func (sr *folderRepository) RemoveSpreading(c context.Context, userID string, folderID string, spreadingID string) error {
	collection := sr.database.Collection(sr.collection)

	folderIDHex, err := primitive.ObjectIDFromHex(folderID)
	if err != nil {
		return err
	}

	_, err = collection.UpdateOne(c, bson.M{"_id": folderIDHex, "user_id": userID}, bson.M{"$pull": bson.M{"spreadouts": spreadingID}})
	return err
}