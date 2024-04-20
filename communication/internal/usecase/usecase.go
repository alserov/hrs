package usecase

import (
	"context"
	"fmt"
	"github.com/alserov/hrs/communication/internal/db"
	"github.com/alserov/hrs/communication/internal/log"
	"github.com/alserov/hrs/communication/internal/usecase/models"
	"github.com/google/uuid"
	"time"
)

func NewUsecase(repo db.Repository) *Usecase {
	return &Usecase{repo: repo}
}

type Usecase struct {
	repo db.Repository
}

func (uc *Usecase) GetHistory(ctx context.Context, history models.HistoryParam) ([]models.Message, error) {
	log.GetLogger(ctx).Debug("uc passed")

	msgs, err := uc.repo.GetHistory(ctx, history)
	if err != nil {
		return nil, fmt.Errorf("repo error: %w", err)
	}

	return msgs, nil
}

func (uc *Usecase) GetChats(ctx context.Context, list models.ChatsParam) ([]models.Chat, error) {
	log.GetLogger(ctx).Debug("uc passed")

	chats, err := uc.repo.GetChats(ctx, list)
	if err != nil {
		return nil, fmt.Errorf("repo error: %w", err)
	}

	return chats, nil
}

func (uc *Usecase) CreateMessage(ctx context.Context, msg models.Message) error {
	msg.ID = uuid.NewString()
	msg.CreatedAt = time.Now()
	msg.UpdatedAt = time.Now()

	log.GetLogger(ctx).Debug("uc passed")

	err := uc.repo.CreateMessage(ctx, msg)
	if err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	return nil
}

func (uc *Usecase) DeleteMessage(ctx context.Context, msgID string) error {
	log.GetLogger(ctx).Debug("uc passed")

	err := uc.repo.DeleteMessage(ctx, msgID)
	if err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	return nil
}

func (uc *Usecase) EditMessage(ctx context.Context, edit models.MessageEdit) error {
	edit.UpdatedAt = time.Now()

	log.GetLogger(ctx).Debug("uc passed")

	err := uc.repo.EditMessage(ctx, edit)
	if err != nil {
		return fmt.Errorf("repo error: %w", err)
	}

	return nil
}
