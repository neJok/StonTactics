package domain

import (
	"context"
	"time"
)

const (
	CollectionRegisterCode = "register_codes"
)

type RegisterCode struct {
	Email     string     `bson:"email"`
	Password  []byte     `bson:"password"`
	Code      int        `bson:"code"`
	CreatedAt *time.Time `bson:"created_at"`
	Attempts  int        `bson:"attempts"`
}

type ConfirmCodeRequest struct {
	Name  string `bson:"name" form:"name" binding:"required,max=30,min=2" json:"name"`
	Email string `bson:"email" form:"email" binding:"required,max=256,min=3" json:"email"`
	Code  int    `bson:"code" form:"code" binding:"required" json:"code"`
}

type RegisterCodesRepository interface {
	CreateRegisterCode(c context.Context, code *RegisterCode) error
	GetRegisterCode(c context.Context, email string) (RegisterCode, error)
	IncAttemptsRegisterCode(c context.Context, email string)
}
