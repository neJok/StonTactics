package domain

import (
	"context"
	"time"
)

const (
	CollectionUser = "users"
)

type UserPro struct {
	Active bool       `bson:"active"`
	Until  *time.Time `bson:"until"`
}

type User struct {
	ID        string  `bson:"_id"`
	Name      string  `bson:"name"`
	Email     string  `bson:"email"`
	AvatarURL string  `bson:"avatar_url"`
	Pro       UserPro `bson:"pro"`
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	UpdateMetaData(c context.Context, id string, name string, avatarUrl string) error
	GetByID(c context.Context, id string) (User, error)
}

type GoogleUserResponse struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
	Name          string `json:"name"`
}

type VKUserResponse struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Photo200  string `json:"photo_200"`
}
