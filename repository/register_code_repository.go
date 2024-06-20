package repository

import (
	"context"

	"github.com/neJok/StonTactics/domain"
	"github.com/neJok/StonTactics/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

type registerCodeRepository struct {
	database   mongo.Database
	collection string
}

func NewRegisterCodeRepository(db mongo.Database, collection string) domain.RegisterCodesRepository {
	return &registerCodeRepository{
		database:   db,
		collection: collection,
	}
}

func (pr *registerCodeRepository) CreateRegisterCode(c context.Context, code *domain.RegisterCode) error {
	collection := pr.database.Collection(pr.collection)

	collection.DeleteMany(c, bson.M{"email": code.Email})
	_, err := collection.InsertOne(c, code)
	return err
}

func (pr *registerCodeRepository) GetRegisterCode(c context.Context, email string) (domain.RegisterCode, error) {
	collection := pr.database.Collection(pr.collection)

	var registerCode domain.RegisterCode

	err := collection.FindOne(c, bson.M{"email": email}).Decode(&registerCode)
	return registerCode, err
}

func (pr *registerCodeRepository) IncAttemptsRegisterCode(c context.Context, email string) {
	collection := pr.database.Collection(pr.collection)

	collection.UpdateOne(c, bson.M{"email": email}, bson.M{"$inc": bson.M{"attempts": -1}})
}
