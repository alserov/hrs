package db

import (
	"context"
	"github.com/alserov/hrs/communication/internal/usecase/models"
)

type Repository interface {
	Chat
	Message
}

type Message interface {
	CreateMessage(ctx context.Context, msg models.Message) error
	DeleteMessage(ctx context.Context, msgID string) error
	EditMessage(ctx context.Context, edit models.MessageEdit) error
}

type Chat interface {
	GetHistory(ctx context.Context, param models.HistoryParam) ([]models.Message, error)
	GetChats(ctx context.Context, param models.ChatsParam) ([]models.Chat, error)
}
