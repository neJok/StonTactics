package domain

import "context"

type SingUpRequest struct {
	Email    string `bson:"email" form:"email" binding:"required,max=100,min=1" json:"email"`
	Password string `bson:"password" form:"password" binding:"required,max=100,min=8" json:"password"`
}

type SingUpUsecase interface {
	CreateRegisterCode(c context.Context, code *RegisterCode) error
	GetRegisterCode(c context.Context, email string) (RegisterCode, error)
	IncAttemptsRegisterCode(c context.Context, email string)
	CreateUser(c context.Context, user *User) (string, error)
	GetUserByEmail(c context.Context, email string) (User, error)
	
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
}
