package usecase

import (
	"context"
	"stontactics/domain"
	"stontactics/internal/tokenutil"
	"time"
)

type singUpUsecase struct {
	userRepository          domain.UserRepository
	registerCodesRepository domain.RegisterCodesRepository
	contextTimeout          time.Duration
}

func NewSignUpUsecase(userRepository domain.UserRepository, registerCodesRepository domain.RegisterCodesRepository, timeout time.Duration) domain.SignUpUsecase {
	return &singUpUsecase{
		userRepository:          userRepository,
		registerCodesRepository: registerCodesRepository,
		contextTimeout:          timeout,
	}
}

func (su *singUpUsecase) CreateUser(c context.Context, user *domain.User) (string, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.Create(ctx, user)
}

func (su *singUpUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.GetUserByEmail(ctx, email)
}

func (su *singUpUsecase) CreateRegisterCode(c context.Context, code *domain.RegisterCode) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.registerCodesRepository.CreateRegisterCode(ctx, code)
}

func (su *singUpUsecase) GetRegisterCode(c context.Context, email string) (domain.RegisterCode, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.registerCodesRepository.GetRegisterCode(ctx, email)
}

func (su *singUpUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (su *singUpUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}

func (su *singUpUsecase) IncAttemptsRegisterCode(c context.Context, email string) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	su.registerCodesRepository.IncAttemptsRegisterCode(ctx, email)
}
