package domain

import (
	"context"
	"time"
)

type Profile struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	AvatarURl string    `json:"avatar_url"`
	Pro       UserPro   `json:"pro"`
	CreatedAt *time.Time `json:"created_at"`
}

type ProfileUsecase interface {
	GetProfileByID(c context.Context, userID string) (*Profile, error)
}
