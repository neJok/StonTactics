package domain

import (
	"context"
	"time"
)

const (
	CollectionPayment = "payments"
)

type Payment struct {
	PaymentID string `bson:"payment_id"`
	UserID 	  string `bson:"user_id"`
	Days	  int16  `bson:"days"`
	Paid	  bool   `bson:"paid"`
}


type PaymentCreateRequest struct {
	Days      int16  `bson:"days" form:"days" binding:"required" json:"days"`
	Email     string `bson:"email" form:"email" binding:"required,max=256,min=3" json:"email"`
}

type PaymentCreated struct {
	Url     string `json:"url"`
}

type PaymentRepository interface {
	SetPaid(c context.Context, id string) error
	GetByID(c context.Context, id string) (Payment, error)
	Create(c context.Context, payment *Payment) error
}

type PaymentUsecase interface {
	SetPaid(c context.Context, id string) error
	GetByID(c context.Context, id string) (Payment, error)
	Create(c context.Context, payment *Payment) error
	ActivatePro(c context.Context, id string, until *time.Time) error
	GetUser(c context.Context, id string) (User, error)
}