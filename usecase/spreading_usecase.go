package usecase

import (
	"context"
	"time"

	"github.com/neJok/StonTactics/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (su *spreadingUsecase) Update(c context.Context, id string, elements []map[string]interface{}, mapName string) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.spreadingRepository.Update(ctx, id, elements, mapName)
}

func (su *spreadingUsecase) GetCount(c context.Context, userID string) int64 {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.spreadingRepository.GetCount(ctx, userID)
}

func (su *spreadingUsecase) DeleteByIDS(c context.Context, userID string, spreadoutsIDS []string) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.spreadingRepository.DeleteByIDS(ctx, userID, spreadoutsIDS)
}
