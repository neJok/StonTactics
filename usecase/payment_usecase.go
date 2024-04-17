package usecase

import (
	"context"
	"time"

	"stontactics/domain"
)

type paymentUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func (su *paymentUsecase) ActivatePro(c context.Context, id string, until *time.Time) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	su.userRepository.ActivatePro(ctx, id, until)
}
