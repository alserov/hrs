package scylla

import (
	"context"
	"github.com/alserov/hrs/communication/internal/db"
	"github.com/alserov/hrs/communication/internal/usecase/models"
)

var _ db.Repository = &Scylla{}

func NewRepository() *Scylla {
	return &Scylla{}
}

func Connect() {

}

type Scylla struct {
}

func (s Scylla) CreateMessage(ctx context.Context, msg models.Message) error {
	//TODO implement me
	panic("implement me")
}

func (s Scylla) DeleteMessage(ctx context.Context, msgID string) error {
	//TODO implement me
	panic("implement me")
}

func (s Scylla) EditMessage(ctx context.Context, edit models.EditMessage) error {
	//TODO implement me
	panic("implement me")
}
