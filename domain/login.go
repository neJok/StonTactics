package domain

import (
	"context"
)

type LoginUsecase interface {
	Create(c context.Context, user *User) (string, error)
	UpdateUser(c context.Context, id string, name string, avatarURL string)
	
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
	
	GetUserByID(c context.Context, id string) (User, error)
	GetUserByGoogleID(c context.Context, id string) (User, error)
	GetUserByVKID(c context.Context, id string) (User, error)
	GetUserByEmail(c context.Context, email string) (User, error)
}

type LoginRequest struct {
	Email    string `bson:"email" form:"email" binding:"required,max=100,min=1" json:"email"`
	Password string `bson:"password" form:"password" binding:"required,max=100" json:"password"`
}