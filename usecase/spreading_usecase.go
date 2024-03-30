package usecase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"stontactics/domain"
	"time"
)

type spreadingUsecase struct {
	spreadingRepository domain.SpreadingRepository
	contextTimeout      time.Duration
}

func NewSpreadingUsecase(spreadingRepository domain.SpreadingRepository, timeout time.Duration) domain.SpreadingUsecase {
	return &spreadingUsecase{
		spreadingRepository: spreadingRepository,
		contextTimeout:      timeout,
	}
}

func (su *spreadingUsecase) Create(c context.Context, spreading *domain.Spreading) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.spreadingRepository.Create(ctx, spreading)
}

func (su *spreadingUsecase) FetchMany(c context.Context, userID string, ids []primitive.ObjectID) ([]domain.SpreadingResponse, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.spreadingRepository.FetchMany(ctx, userID, ids)
}

func (su *spreadingUsecase) FetchByID(c context.Context, id string) (domain.Spreading, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.spreadingRepository.FetchByID(ctx, id)
}

func (su *spreadingUsecase) Update(c context.Context, id string, elements []map[string]interface{}) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.spreadingRepository.Update(ctx, id, elements)
}
