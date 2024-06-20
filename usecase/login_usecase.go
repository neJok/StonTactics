package usecase

import (
	"context"
	"time"

	"github.com/neJok/StonTactics/domain"
	"github.com/neJok/StonTactics/internal/tokenutil"

	"go.mongodb.org/mongo-driver/bson"
)

type loginUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewLoginUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (lu *loginUsecase) Create(c context.Context, user *domain.User) (string, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.Create(ctx, user)
}

func (lu *loginUsecase) GetUserByID(c context.Context, id string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.GetByID(ctx, id)
}

func (lu *loginUsecase) UpdateUser(c context.Context, id string, data bson.M) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	lu.userRepository.Update(ctx, id, data)
}

func (lu *loginUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (lu *loginUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}

func (lu *loginUsecase) GetUserByGoogleID(c context.Context, id string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.GetUserByGoogleID(ctx, id)
}

func (lu *loginUsecase) GetUserByVKID(c context.Context, id string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.GetUserByVKID(ctx, id)
}

func (su *loginUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.GetUserByEmail(ctx, email)
}

func (su *loginUsecase) DeleteByID(c context.Context, id string) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	su.userRepository.DeleteByID(ctx, id)
}
