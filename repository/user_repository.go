package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"stontactics/domain"
	"stontactics/mongo"
)

type userRepository struct {
	database   mongo.Database
	collection string
}

func NewUserRepository(db mongo.Database, collection string) domain.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *userRepository) Create(c context.Context, user *domain.User) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(c, user)

	return err
}

func (ur *userRepository) GetByID(c context.Context, id string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)

	var user domain.User
	err := collection.FindOne(c, bson.M{"_id": id}).Decode(&user)
	return user, err
}

func (ur *userRepository) UpdateMetaData(c context.Context, id string, name string, avatarUrl string) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.UpdateOne(c, bson.M{"_id": id}, bson.M{"$set": bson.M{"name": name, "avatar_url": avatarUrl}})
	return err
}
