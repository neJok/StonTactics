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

func (fu *folderUsecase) DeleteOneByID(c context.Context, userID string, folderID string) error {
	ctx, cancel := context.WithTimeout(c, fu.contextTimeout)
	defer cancel()
	return fu.folderRepository.DeleteOneByID(ctx, userID, folderID)
}

func (fu *folderUsecase) AddStrategies(c context.Context, userID string, folderID string, strategiesIDS []string) error {
	ctx, cancel := context.WithTimeout(c, fu.contextTimeout)
	defer cancel()
	return fu.folderRepository.AddStrategies(ctx, userID, folderID, strategiesIDS)
}

func (fu *folderUsecase) AddSpreadouts(c context.Context, userID string, folderID string, spreadoutsIDS []string) error {
	ctx, cancel := context.WithTimeout(c, fu.contextTimeout)
	defer cancel()
	return fu.folderRepository.AddSpreadouts(ctx, userID, folderID, spreadoutsIDS)
}

func (fu *folderUsecase) RemoveStrategies(c context.Context, userID string, folderID string, strategiesIDS []string) error {
	ctx, cancel := context.WithTimeout(c, fu.contextTimeout)
	defer cancel()
	return fu.folderRepository.RemoveStrategies(ctx, userID, folderID, strategiesIDS)
}

func (fu *folderUsecase) RemoveSpreadouts(c context.Context, userID string, folderID string, spreadoutsIDS []string) error {
	ctx, cancel := context.WithTimeout(c, fu.contextTimeout)
	defer cancel()
	return fu.folderRepository.RemoveSpreadouts(ctx, userID, folderID, spreadoutsIDS)
}
