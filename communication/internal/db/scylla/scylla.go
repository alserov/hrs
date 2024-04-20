package scylla

import (
	"context"
	"github.com/alserov/hrs/communication/internal/db"
	"github.com/alserov/hrs/communication/internal/usecase/models"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

var _ db.Repository = &Scylla{}

func NewRepository(s gocqlx.Session) *Scylla {
	return &Scylla{s}
}

func MustConnect(dsn string) gocqlx.Session {
	cluster := gocql.NewCluster(dsn)

	ws, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		panic("failed to create session: " + err.Error())
	}

	return ws
}

type Scylla struct {
	gocqlx.Session
}

func (s Scylla) CreateMessage(ctx context.Context, msg models.Message) error {

	//if err != nil {
	//	return utils.NewError(utils.ErrInternal, err.Error())
	//}

	return nil
}

func (s Scylla) DeleteMessage(ctx context.Context, msgID string) error {

	//if err != nil {
	//	return utils.NewError(utils.ErrInternal, err.Error())
	//}

	return nil
}

func (s Scylla) EditMessage(ctx context.Context, edit models.MessageEdit) error {

	//if err != nil {
	//	return utils.NewError(utils.ErrInternal, err.Error())
	//}

	return nil
}
