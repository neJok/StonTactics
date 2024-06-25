package usecase

import (
	"context"
	"time"

	"github.com/neJok/StonTactics/domain"
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

func (fu *folderUsecase) Create(c context.Context, folder *domain.Folder) error {
	ctx, cancel := context.WithTimeout(c, fu.contextTimeout)
	defer cancel()
	return fu.folderRepository.Create(ctx, folder)
}

func (fu *folderUsecase) FetchByUserID(c context.Context, userID string) ([]domain.Folder, error) {
	ctx, cancel := context.WithTimeout(c, fu.contextTimeout)
	defer cancel()
	return fu.folderRepository.FetchByUserID(ctx, userID)
}

func (fu *folderUsecase) FetchOneByID(c context.Context, userID string, folderID string) (domain.Folder, error) {
	ctx, cancel := context.WithTimeout(c, fu.contextTimeout)
	defer cancel()
	return fu.folderRepository.FetchOneByID(ctx, userID, folderID)
}

func (fu *folderUsecase) AddStrategy(c context.Context, userID string, folderID string, strategyID string) error {
	ctx, cancel := context.WithTimeout(c, fu.contextTimeout)
	defer cancel()
	return fu.folderRepository.AddStrategy(ctx, userID, folderID, strategyID)
}

func (fu *folderUsecase) AddSpreading(c context.Context, userID string, folderID string, spreadingID string) error {
	ctx, cancel := context.WithTimeout(c, fu.contextTimeout)
	defer cancel()
	return fu.folderRepository.AddSpreading(ctx, userID, folderID, spreadingID)
}

func (fu *folderUsecase) RemoveStrategy(c context.Context, userID string, folderID string, strategyID string) error {
	ctx, cancel := context.WithTimeout(c, fu.contextTimeout)
	defer cancel()
	return fu.folderRepository.RemoveStrategy(ctx, userID, folderID, strategyID)
}

func (fu *folderUsecase) RemoveSpreading(c context.Context, userID string, folderID string, spreadingID string) error {
	ctx, cancel := context.WithTimeout(c, fu.contextTimeout)
	defer cancel()
	return fu.folderRepository.RemoveSpreading(ctx, userID, folderID, spreadingID)
}
