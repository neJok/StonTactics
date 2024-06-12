package repository

import (
	"context"
	"stontactics/domain"
	"stontactics/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

type resetPasswordRepository struct {
	database   mongo.Database
	collection string
}

func NewResetPasswordRepository(db mongo.Database, collection string) domain.ResetPasswordCodesRepository {
	return &resetPasswordRepository{
		database:   db,
		collection: collection,
	}
}

func (pr *resetPasswordRepository) Create(c context.Context, code *domain.ResetPassword) error {
	collection := pr.database.Collection(pr.collection)

	collection.DeleteMany(c, bson.M{"email": code.Email})
	_, err := collection.InsertOne(c, code)
	return err
}

func (pr *resetPasswordRepository) GetByEmail(c context.Context, email string) (domain.ResetPassword, error) {
	collection := pr.database.Collection(pr.collection)

	var code domain.ResetPassword

	err := collection.FindOne(c, bson.M{"email": email}).Decode(&code)
	return code, err
}

func (pr *resetPasswordRepository) GetByToken(c context.Context, token string) (domain.ResetPassword, error) {
	collection := pr.database.Collection(pr.collection)

	var code domain.ResetPassword

	err := collection.FindOne(c, bson.M{"token": token}).Decode(&code)
	return code, err
}

func (pr *resetPasswordRepository) IncAttempts(c context.Context, email string) {
	collection := pr.database.Collection(pr.collection)

	collection.UpdateOne(c, bson.M{"email": email}, bson.M{"$inc": bson.M{"attempts": -1}})
}

func (pr *resetPasswordRepository) DeleteByEmail(c context.Context, email string) {
	collection := pr.database.Collection(pr.collection)

	collection.DeleteMany(c, bson.M{"email": email})
}