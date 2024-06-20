package repository

import (
	"context"

	"github.com/neJok/StonTactics/domain"
	"github.com/neJok/StonTactics/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

type changeEmailRepository struct {
	database   mongo.Database
	collection string
}

func NewChangeEmailRepository(db mongo.Database, collection string) domain.ChangeEmailCodesRepository {
	return &changeEmailRepository{
		database:   db,
		collection: collection,
	}
}

func (cr *changeEmailRepository) Create(c context.Context, code *domain.ChangeEmail) error {
	collection := cr.database.Collection(cr.collection)

	collection.DeleteMany(c, bson.M{"user_id": code.UserID})
	_, err := collection.InsertOne(c, code)
	return err
}

func (pr *changeEmailRepository) GetByID(c context.Context, id string) (domain.ChangeEmail, error) {
	collection := pr.database.Collection(pr.collection)

	var code domain.ChangeEmail

	err := collection.FindOne(c, bson.M{"user_id": id}).Decode(&code)
	return code, err
}

func (pr *changeEmailRepository) IncAttempts(c context.Context, id string) {
	collection := pr.database.Collection(pr.collection)

	collection.UpdateOne(c, bson.M{"user_id": id}, bson.M{"$inc": bson.M{"attempts": -1}})
}

func (pr *changeEmailRepository) DeleteByID(c context.Context, id string) {
	collection := pr.database.Collection(pr.collection)

	collection.DeleteMany(c, bson.M{"user_id": id})
}
