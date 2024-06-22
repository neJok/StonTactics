package usecase

import (
	"context"
	"time"

	"github.com/neJok/StonTactics/domain"
	"go.mongodb.org/mongo-driver/bson"
)

type accountUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewAccountUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.AccountUsecase {
	return &accountUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (au *accountUsecase) DeleteByID(c context.Context, id string) {
	ctx, cancel := context.WithTimeout(c, au.contextTimeout)
	defer cancel()
	au.userRepository.DeleteByID(ctx, id)
}

func (au *accountUsecase) GetByAccountByID(c context.Context, id string) (*domain.Account, error) {
	ctx, cancel := context.WithTimeout(c, au.contextTimeout)
	defer cancel()

	user, err := au.userRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &domain.Account{
		ID: id, Name: user.Name,
		Email:     user.Auth.Email.Email,
		AvatarURl: user.AvatarURL,
		Pro:       user.Pro,
		CreatedAt: user.CreatedAt,
		VK:        user.Auth.VK,
	}, nil
}

func (au *accountUsecase) UpdateByID(c context.Context, id string, data bson.M) error {
	ctx, cancel := context.WithTimeout(c, au.contextTimeout)
	defer cancel()
	return au.userRepository.Update(ctx, id, data)
}