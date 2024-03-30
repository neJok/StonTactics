package domain

import (
	"context"
)

type LoginUsecase interface {
	Create(c context.Context, user *User) error
	GetUserByID(c context.Context, id string) (User, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
	UpdateUser(c context.Context, id string, name string, avatarURL string)
}
