package usecase

import (
	"context"
	"time"

	"github.com/neJok/StonTactics/domain"
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
