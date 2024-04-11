package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"stontactics/domain"
	"stontactics/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

type strategyRepository struct {
	database   mongo.Database
	collection string
}

func NewStrategyRepository(db mongo.Database, collection string) domain.StrategyRepository {
	return &strategyRepository{
		database:   db,
		collection: collection,
	}
}

func (sr *strategyRepository) Create(c context.Context, strategy *domain.Strategy) error {
	collection := sr.database.Collection(sr.collection)

	_, err := collection.InsertOne(c, strategy)
	return err
}

func (sr *strategyRepository) FetchMany(c context.Context, userID string, ids []primitive.ObjectID) ([]domain.StrategyResponse, error) {
	collection := sr.database.Collection(sr.collection)

	var strategies []domain.StrategyResponse

	opts := options.Find().SetProjection(bson.D{{Key: "_id", Value: 1}, {Key: "name", Value: 1}, {Key: "map_name", Value: 1}})

	filter := bson.M{"user_id": userID}
	if len(ids) > 0 {
		filter["_id"] = bson.M{"$in": ids}
	}

	cursor, err := collection.Find(c, filter, opts)
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &strategies)
	if strategies == nil {
		return []domain.StrategyResponse{}, err
	}

	return strategies, err
}

func (sr *strategyRepository) FetchByID(c context.Context, id string) (domain.Strategy, error) {
	collection := sr.database.Collection(sr.collection)

	var strategy domain.Strategy

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return strategy, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&strategy)
	if err != nil {
		return domain.Strategy{}, err
	}

	return strategy, err
}

func (sr *strategyRepository) Update(c context.Context, id string, parts map[string]interface{}, mapName string) error {
	collection := sr.database.Collection(sr.collection)

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.UpdateOne(c, bson.M{"_id": idHex}, bson.M{"$set": bson.M{"parts": parts, "map_name": mapName}})
	return err
}
