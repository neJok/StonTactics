package usecase

import (
	"context"
	"time"

	"stontactics/domain"
)

type paymentUsecase struct {
	userRepository 	  domain.UserRepository
	paymentRepository domain.PaymentRepository
	contextTimeout    time.Duration
}

func NewPaymentUsecase(paymentRepository domain.PaymentRepository, userRepository domain.UserRepository, timeout time.Duration) domain.PaymentUsecase {
	return &paymentUsecase{
		userRepository: userRepository,
		paymentRepository: paymentRepository,
		contextTimeout: timeout,
	}
}

func (pu *paymentUsecase) ActivatePro(c context.Context, id string, until *time.Time) error {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.userRepository.ActivatePro(ctx, id, until)
}

func (pu *paymentUsecase) Create(c context.Context, payment *domain.Payment) error {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.paymentRepository.Create(ctx, payment)
}

func (pu *paymentUsecase) GetByID(c context.Context, id string) (domain.Payment, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.paymentRepository.GetByID(ctx, id)
}

func (pu *paymentUsecase) SetPaid(c context.Context, id string) error {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.paymentRepository.SetPaid(ctx, id)
}

func (pu *paymentUsecase) GetUser(c context.Context, id string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.userRepository.GetByID(ctx, id)
}