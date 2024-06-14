package domain

import (
	"context"
	"time"
)

type Account struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	AvatarURl string     `json:"avatar_url"`
	Pro       UserPro    `json:"pro"`
	CreatedAt *time.Time `json:"created_at"`
}

type AccountUsecase interface {
	GetByAccountByID(c context.Context, id string) (*Account, error)
	DeleteByID(c context.Context, id string)
}
