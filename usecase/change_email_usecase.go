package usecase

import (
	"context"
	"stontactics/domain"
	"time"
)

type changeEmailUsecase struct {
	userRepository             domain.UserRepository
	changeEmailCodesRepository domain.ChangeEmailCodesRepository
	contextTimeout             time.Duration
}

func NewChangeEmailUsecase(userRepository domain.UserRepository, changeEmailCodesRepository domain.ChangeEmailCodesRepository, timeout time.Duration) domain.ChangeEmailUsecase {
	return &changeEmailUsecase{
		userRepository:             userRepository,
		changeEmailCodesRepository: changeEmailCodesRepository,
		contextTimeout:             timeout,
	}
}

func (cu *changeEmailUsecase) GetUserByID(c context.Context, id string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.userRepository.GetByID(ctx, id)
}

func (cu *changeEmailUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.userRepository.GetUserByEmail(ctx, email)
}

func (cu *changeEmailUsecase) UpdateEmail(c context.Context, id string, email string) error {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.userRepository.UpdateEmail(ctx, id, email)
}

func (cu *changeEmailUsecase) CreateCode(c context.Context, code *domain.ChangeEmail) error {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.changeEmailCodesRepository.Create(ctx, code)
}

func (cu *changeEmailUsecase) GetCodeByID(c context.Context, id string) (domain.ChangeEmail, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.changeEmailCodesRepository.GetByID(ctx, id)
}

func (cu *changeEmailUsecase) IncCodeAttempts(c context.Context, id string) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	cu.changeEmailCodesRepository.IncAttempts(ctx, id)
}

func (cu *changeEmailUsecase) DeleteCodeByID(c context.Context, id string) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	cu.changeEmailCodesRepository.DeleteByID(ctx, id)
}
