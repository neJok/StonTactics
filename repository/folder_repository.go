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

func (sr *folderRepository) DeleteOneByID(c context.Context, userID string, folderID string) error {
	collection := sr.database.Collection(sr.collection)

	folderIDHex, err := primitive.ObjectIDFromHex(folderID)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(c, bson.M{"_id": folderIDHex, "user_id": userID})
	return err
}

func (sr *folderRepository) AddStrategies(c context.Context, userID string, folderID string, strategiesIDS []string) error {
	collection := sr.database.Collection(sr.collection)

	folderIDHex, err := primitive.ObjectIDFromHex(folderID)
	if err != nil {
		return err
	}

	_, err = collection.UpdateOne(c, bson.M{"_id": folderIDHex, "user_id": userID}, bson.M{"$addToSet": bson.M{"strategies": bson.M{"$each": strategiesIDS}}})
	return err
}

func (sr *folderRepository) AddSpreadouts(c context.Context, userID string, folderID string, spreadoutsIDS []string) error {
	collection := sr.database.Collection(sr.collection)

	folderIDHex, err := primitive.ObjectIDFromHex(folderID)
	if err != nil {
		return err
	}

	_, err = collection.UpdateOne(c, bson.M{"_id": folderIDHex, "user_id": userID}, bson.M{"$addToSet": bson.M{"spreadouts": bson.M{"$each": spreadoutsIDS}}})
	return err
}

func (sr *folderRepository) RemoveStrategies(c context.Context, userID string, folderID string, strategiesIDS []string) error {
	collection := sr.database.Collection(sr.collection)

	folderIDHex, err := primitive.ObjectIDFromHex(folderID)
	if err != nil {
		return err
	}

	_, err = collection.UpdateOne(c, bson.M{"_id": folderIDHex, "user_id": userID}, bson.M{"$pull": bson.M{"strategies": bson.M{"$in": strategiesIDS}}})
	return err
}

func (sr *folderRepository) RemoveSpreadouts(c context.Context, userID string, folderID string, spreadoutsIDS []string) error {
	collection := sr.database.Collection(sr.collection)

	folderIDHex, err := primitive.ObjectIDFromHex(folderID)
	if err != nil {
		return err
	}

	_, err = collection.UpdateOne(c, bson.M{"_id": folderIDHex, "user_id": userID}, bson.M{"$pull": bson.M{"spreadouts": bson.M{"$in": spreadoutsIDS}}})
	return err
}