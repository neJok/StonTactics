package usecase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"stontactics/domain"
	"time"
)

type strategyUsecase struct {
	strategyRepository domain.StrategyRepository
	contextTimeout     time.Duration
}

func NewStrategyUsecase(strategyRepository domain.StrategyRepository, timeout time.Duration) domain.StrategyUsecase {
	return &strategyUsecase{
		strategyRepository: strategyRepository,
		contextTimeout:     timeout,
	}
}

func (su *strategyUsecase) Create(c context.Context, strategy *domain.Strategy) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.strategyRepository.Create(ctx, strategy)
}

func (su *strategyUsecase) FetchMany(c context.Context, userID string, ids []primitive.ObjectID) ([]domain.StrategyResponse, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.strategyRepository.FetchMany(ctx, userID, ids)
}

func (su *strategyUsecase) FetchByID(c context.Context, id string) (domain.Strategy, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.strategyRepository.FetchByID(ctx, id)
}

func (su *strategyUsecase) Update(c context.Context, id string, parts map[string]interface{}, mapName string) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.strategyRepository.Update(ctx, id, parts, mapName)
}

func (su *strategyUsecase) GetCount(c context.Context, userID string) int64 {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.strategyRepository.GetCount(ctx, userID)
}
