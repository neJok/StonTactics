package repository

import (
	"context"

	"github.com/neJok/StonTactics/domain"
	"github.com/neJok/StonTactics/mongo"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
)

type spreadingRepository struct {
	database   mongo.Database
	collection string
}

func NewSpreadingRepository(db mongo.Database, collection string) domain.SpreadingRepository {
	return &spreadingRepository{
		database:   db,
		collection: collection,
	}
}

func (sr *spreadingRepository) Create(c context.Context, spreading *domain.Spreading) error {
	collection := sr.database.Collection(sr.collection)

	_, err := collection.InsertOne(c, spreading)
	return err
}

func (sr *spreadingRepository) FetchMany(c context.Context, userID string, ids []primitive.ObjectID) ([]domain.SpreadingResponse, error) {
	collection := sr.database.Collection(sr.collection)

	var spreadouts []domain.SpreadingResponse

	opts := options.Find().SetProjection(bson.D{{Key: "_id", Value: 1}, {Key: "name", Value: 1}, {Key: "map_name", Value: 1}})

	filter := bson.M{"user_id": userID}
	if len(ids) > 0 {
		filter["_id"] = bson.M{"$in": ids}
	}

	cursor, err := collection.Find(c, filter, opts)
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &spreadouts)
	if spreadouts == nil {
		return []domain.SpreadingResponse{}, err
	}

	return spreadouts, err
}

func (sr *spreadingRepository) FetchByID(c context.Context, id string) (domain.Spreading, error) {
	collection := sr.database.Collection(sr.collection)

	var spreading domain.Spreading

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return spreading, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&spreading)
	if err != nil {
		return domain.Spreading{}, err
	}

	return spreading, err
}

func (sr *spreadingRepository) Update(c context.Context, id string, elements []map[string]interface{}, mapName string) error {
	collection := sr.database.Collection(sr.collection)

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.UpdateOne(c, bson.M{"_id": idHex}, bson.M{"$set": bson.M{"elements": elements, "map_name": mapName}})
	return err
}

func (sr *spreadingRepository) GetCount(c context.Context, userID string) int64 {
	collection := sr.database.Collection(sr.collection)

	count, _ := collection.CountDocuments(c, bson.M{"user_id": userID})
	return count
}

func (sr *spreadingRepository) DeleteByID(c context.Context, userID, spreadingID string) error {
	collection := sr.database.Collection(sr.collection)

	idHex, err := primitive.ObjectIDFromHex(spreadingID)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(c, bson.M{"user_id": userID, "_id": idHex})
	return err
}