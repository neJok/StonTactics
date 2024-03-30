package usecase

import (
	"context"
	"stontactics/domain"
	"time"
)

type folderUsecase struct {
	folderRepository domain.FolderRepository
	contextTimeout   time.Duration
}

func NewFolderUsecase(folderRepository domain.FolderRepository, timeout time.Duration) domain.FolderUsecase {
	return &folderUsecase{
		folderRepository: folderRepository,
		contextTimeout:   timeout,
	}
}

func (su *folderUsecase) Create(c context.Context, folder *domain.Folder) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.folderRepository.Create(ctx, folder)
}

func (su *folderUsecase) FetchByUserID(c context.Context, userID string) ([]domain.Folder, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.folderRepository.FetchByUserID(ctx, userID)
}

func (su *folderUsecase) AddStrategy(c context.Context, userID string, folderID string, strategyID string) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.folderRepository.AddStrategy(ctx, userID, folderID, strategyID)
}

func (su *folderUsecase) AddSpreading(c context.Context, userID string, folderID string, spreadingID string) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.folderRepository.AddSpreading(ctx, userID, folderID, spreadingID)
}
