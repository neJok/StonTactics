package domain

import (
	"context"
	"time"
)

const (
	CollectionResetPasswordCode = "reset_password_codes"
)

type ResetPassword struct {
	Email     string     `bson:"email"`
	Code      int        `bson:"code"`
	Token     string     `bson:"token"` // второй код для фронтенда (костыль из-за смены пароля в два этапа)
	CreatedAt *time.Time `bson:"created_at"`
	Attempts  int        `bson:"attempts"`
}

type ResetPassowrdCreate struct {
	Email string `bson:"email" form:"email" binding:"required,min=1,max=100" json:"email"`
}

type ResetPasswordConfirmRequest struct {
	Email string `bson:"email" form:"email" binding:"required,min=1,max=100" json:"email"`
	Code  int    `bson:"code" form:"code" binding:"required" json:"code"`
}

type ResetPasswordRequest struct {
	Token    string `bson:"token" form:"token" binding:"required" json:"token"`
	Password string `bson:"password" form:"password" binding:"required,max=100,min=8" json:"password"`
}

type ResetPasswordResponse struct {
	Token string `bson:"token" form:"token" binding:"required" json:"token"`
}

type ResetPasswordCodesRepository interface {
	Create(c context.Context, code *ResetPassword) error
	GetByEmail(c context.Context, email string) (ResetPassword, error)
	GetByToken(c context.Context, token string) (ResetPassword, error)
	DeleteByEmail(c context.Context, email string)
	IncAttempts(c context.Context, email string)
}

type ResetPassowrdUsecase interface {
	CreateCode(c context.Context, code *ResetPassword) error
	GetCodeByEmail(c context.Context, email string) (ResetPassword, error)
	GetCodeByToken(c context.Context, token string) (ResetPassword, error)
	DeleteCodeByEmail(c context.Context, email string)
	IncCodeAttempts(c context.Context, email string)

	GetUserByEmail(c context.Context, email string) (User, error)
	UpdatePassword(c context.Context, id string, password []byte) error
}
