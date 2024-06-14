package domain

import (
	"context"
	"time"
)

const (
	CollectionChangeEmailCode = "change_email_codes"
)

type ChangeEmail struct {
	UserID    string     `bson:"user_id"`
	Email     string     `bson:"email"`
	Code      int        `bson:"code"`
	CreatedAt *time.Time `bson:"created_at"`
	Attempts  int        `bson:"attempts"`
}

type ChangeEmailCreate struct {
	Email string `bson:"email" form:"email" binding:"required,min=1,max=100" json:"email"`
}

type ChangeEmailConfirmRequest struct {
	Code  int    `bson:"code" form:"code" binding:"required" json:"code"`
}

type ChangeEmailCodesRepository interface {
	Create(c context.Context, code *ChangeEmail) error
	GetByID(c context.Context, id string) (ChangeEmail, error)
	DeleteByID(c context.Context, id string)
	IncAttempts(c context.Context, id string)
}

type ChangeEmailUsecase interface {
	CreateCode(c context.Context, code *ChangeEmail) error
	GetCodeByID(c context.Context, id string) (ChangeEmail, error)
	DeleteCodeByID(c context.Context, id string)
	IncCodeAttempts(c context.Context, id string)

	GetUserByID(c context.Context, id string) (User, error)
	GetUserByEmail(c context.Context, email string) (User, error)
	UpdateEmail(c context.Context, id string, email string) error
}
