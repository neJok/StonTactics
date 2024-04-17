package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"stontactics/domain"
	"stontactics/mongo"
)

type paymentRepository struct {
	database   mongo.Database
	collection string
}

func NewPaymentRepository(db mongo.Database, collection string) domain.PaymentRepository {
	return &paymentRepository{
		database:   db,
		collection: collection,
	}
}

func (pr *paymentRepository) Create(c context.Context, payment *domain.Payment) error {
	collection := pr.database.Collection(pr.collection)

	_, err := collection.InsertOne(c, payment)
	return err
}

func (pr *paymentRepository) GetByID(c context.Context, id string) (domain.Payment, error) {
	collection := pr.database.Collection(pr.collection)

	var payment domain.Payment
	err := collection.FindOne(c, bson.M{"payment_id": id}).Decode(&payment)

	if err != nil {
		return domain.Payment{}, err
	}

	return payment, err
}

func (pr *paymentRepository) SetPaid(c context.Context, id string) error {
	collection := pr.database.Collection(pr.collection)

	_, err := collection.UpdateOne(c, bson.M{"payment_id": id}, bson.M{"$set": bson.M{"paid": true}})
	return err
}