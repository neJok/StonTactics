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

type ComfirmCodeRequest struct {
	Email string `bson:"email"`
	Code  int    `bson:"code"`
}

type RegisterCodesRepository interface {
	CreateRegisterCode(c context.Context, code *RegisterCode) error
	GetRegisterCode(c context.Context, email string) (RegisterCode, error)
	IncAttemptsRegisterCode(c context.Context, email string)
}
