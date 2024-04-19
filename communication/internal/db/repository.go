package db

import (
	"context"
	"github.com/alserov/hrs/communication/internal/usecase/models"
)

type Repository interface {
	CreateMessage(ctx context.Context, msg models.Message) error
	DeleteMessage(ctx context.Context, msgID string) error
	EditMessage(ctx context.Context, edit models.MessageEdit) error
}
