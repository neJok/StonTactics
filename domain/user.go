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

type EmailAuth struct {
	Email    string `bson:"email"`
	Password []byte `bson:"password"`
}

type SocialAuth struct {
	ID string `bson:"id"`
}

type UserAuth struct {
	Email  EmailAuth  `bson:"email"`
	VK     SocialAuth `bson:"vk"`
	Google SocialAuth `bson:"google"`
}

type User struct {
	ID        string    `bson:"_id"`
	Name      string    `bson:"name"`
	AvatarURL string    `bson:"avatar_url"`
	Pro       UserPro   `bson:"pro"`
	Auth      UserAuth  `bson:"auth"`
	CreatedAt *time.Time `bson:"created_at"`
}

type UserRepository interface {
	Create(c context.Context, user *User) (string, error)
	UpdateMetaData(c context.Context, id string, name string, avatarUrl string) error
	ActivatePro(c context.Context, id string, until *time.Time) error
	GetByID(c context.Context, id string) (User, error)

	GetUserByGoogleID(c context.Context, id string) (User, error)
	GetUserByVKID(c context.Context, id string) (User, error)
	GetUserByEmail(c context.Context, email string) (User, error)
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
