package usecase

import (
	"context"
	"github.com/alserov/hrs/communication/internal/usecase/models"
)

type Usecase struct {
}

func (uc *Usecase) CreateMessage(ctx context.Context, msg models.Message) error {
	return nil
}

func (uc *Usecase) DeleteMessage(ctx context.Context, msgID string) error {
	return nil
}

func (uc *Usecase) EditMessage(ctx context.Context, edit models.EditMessage) error {
	return nil
}
