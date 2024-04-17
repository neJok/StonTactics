package domain

import (
	"context"
	"time"
)

type PaymentCreateRequest struct {
	Days      string `bson:"days" form:"days" json:"days"`
	Email     string `bson:"email" form:"email" binding:"required,max=256,min=3" json:"email"`
}

type PaymentCreated struct {
	Url     string `json:"url"`
}

type PaymentUsecase interface {
	ActivatePro(c context.Context, id string, until *time.Time) error
}
