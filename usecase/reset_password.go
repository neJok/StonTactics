package usecase

import (
	"context"
	"stontactics/domain"
	"time"
)

type resetPassowrdUsecase struct {
	userRepository               domain.UserRepository
	resetPasswordCodesRepository domain.ResetPasswordCodesRepository
	contextTimeout               time.Duration
}

func NewResetPasswordUsecase(userRepository domain.UserRepository, resetPasswordCodesRepository domain.ResetPasswordCodesRepository, timeout time.Duration) domain.ResetPassowrdUsecase {
	return &resetPassowrdUsecase{
		userRepository:               userRepository,
		resetPasswordCodesRepository: resetPasswordCodesRepository,
		contextTimeout:               timeout,
	}
}

func (ru *resetPassowrdUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, ru.contextTimeout)
	defer cancel()
	return ru.userRepository.GetUserByEmail(ctx, email)
}

func (ru *resetPassowrdUsecase) UpdatePassword(c context.Context, id string, password []byte) error {
	ctx, cancel := context.WithTimeout(c, ru.contextTimeout)
	defer cancel()
	return ru.userRepository.UpdatePassword(ctx, id, password)
}

func (ru *resetPassowrdUsecase) CreateCode(c context.Context, code *domain.ResetPassword) error {
	ctx, cancel := context.WithTimeout(c, ru.contextTimeout)
	defer cancel()
	return ru.resetPasswordCodesRepository.Create(ctx, code)
}

func (ru *resetPassowrdUsecase) GetCodeByEmail(c context.Context, email string) (domain.ResetPassword, error) {
	ctx, cancel := context.WithTimeout(c, ru.contextTimeout)
	defer cancel()
	return ru.resetPasswordCodesRepository.GetByEmail(ctx, email)
}

func (ru *resetPassowrdUsecase) GetCodeByToken(c context.Context, token string) (domain.ResetPassword, error) {
	ctx, cancel := context.WithTimeout(c, ru.contextTimeout)
	defer cancel()
	return ru.resetPasswordCodesRepository.GetByToken(ctx, token)
}

func (su *resetPassowrdUsecase) IncCodeAttempts(c context.Context, email string) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	su.resetPasswordCodesRepository.IncAttempts(ctx, email)
}

func (su *resetPassowrdUsecase) DeleteCodeByEmail(c context.Context, email string) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	su.resetPasswordCodesRepository.DeleteByEmail(ctx, email)
}
