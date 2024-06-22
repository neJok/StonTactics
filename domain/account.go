package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Account struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	AvatarURL string     `json:"avatar_url"`
	Pro       UserPro    `json:"pro"`
	CreatedAt *time.Time `json:"created_at"`
	VK        SocialAuth `json:"vk"`
}

type AccountUsecase interface {
	GetByAccountByID(c context.Context, id string) (*Account, error)
	DeleteByID(c context.Context, id string)
	UpdateByID(c context.Context, id string, data bson.M) error
}
