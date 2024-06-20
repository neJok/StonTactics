package repository

import (
	"context"
	"strconv"
	"time"

	"github.com/neJok/StonTactics/domain"
	"github.com/neJok/StonTactics/mongo"

	"go.mongodb.org/mongo-driver/bson"
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

func (ur *userRepository) Create(c context.Context, user *domain.User) (string, error) {
	collection := ur.database.Collection(ur.collection)

	countDocuments, _ := collection.CountDocuments(c, bson.M{})
	id := strconv.Itoa(int(countDocuments) + 1)
	user.ID = id
	_, err := collection.InsertOne(c, user)
	return id, err
}

func CheckPro(c context.Context, collection mongo.Collection, user *domain.User) domain.UserPro {
	if user.Pro.Active && user.Pro.Until != nil {
		now := time.Now()
		if now.After(*user.Pro.Until) {
			user.Pro = domain.UserPro{
				Active: false,
				Until:  nil,
			}
			defer collection.UpdateOne(c, bson.M{"_id": user.ID}, bson.M{"$set": bson.M{"pro": user.Pro}})
		}
	}
	return user.Pro
}

func (ur *userRepository) GetByID(c context.Context, id string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)

	var user domain.User
	err := collection.FindOne(c, bson.M{"_id": id}).Decode(&user)

	if err != nil {
		return user, err
	}

	user.Pro = CheckPro(c, collection, &user)
	return user, nil
}

func (ur *userRepository) GetUserByVKID(c context.Context, id string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)

	var user domain.User
	err := collection.FindOne(c, bson.M{"auth.vk.id": id}).Decode(&user)
	if err != nil {
		return user, err
	}

	user.Pro = CheckPro(c, collection, &user)
	return user, nil
}

func (ur *userRepository) GetUserByGoogleID(c context.Context, id string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)

	var user domain.User
	err := collection.FindOne(c, bson.M{"auth.google.id": id}).Decode(&user)
	if err != nil {
		return user, err
	}

	user.Pro = CheckPro(c, collection, &user)
	return user, nil
}

func (ur *userRepository) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)

	var user domain.User
	err := collection.FindOne(c, bson.M{"auth.email.email": email}).Decode(&user)
	if err != nil {
		return user, err
	}

	user.Pro = CheckPro(c, collection, &user)
	return user, nil
}

func (ur *userRepository) Update(c context.Context, id string, data bson.M) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.UpdateOne(c, bson.M{"_id": id}, bson.M{"$set": data})
	return err
}

func (ur *userRepository) ActivatePro(c context.Context, id string, until *time.Time) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.UpdateOne(c, bson.M{"_id": id}, bson.M{"$set": bson.M{"pro.active": true, "pro.until": until}})
	return err
}

func (ur *userRepository) UpdatePassword(c context.Context, id string, password []byte) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.UpdateOne(c, bson.M{"_id": id}, bson.M{"$set": bson.M{"auth.email.password": password}})
	return err
}

func (ur *userRepository) DeleteByID(c context.Context, id string) {
	collection := ur.database.Collection(ur.collection)

	collection.DeleteOne(c, bson.M{"_id": id})
}
