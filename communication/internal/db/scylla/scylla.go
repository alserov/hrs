package scylla

import (
	"context"
	"fmt"
	"github.com/alserov/hrs/communication/internal/db"
	"github.com/alserov/hrs/communication/internal/log"
	"github.com/alserov/hrs/communication/internal/usecase/models"
	"github.com/alserov/hrs/communication/internal/utils"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/table"
)

var _ db.Repository = &Scylla{}

func NewRepository(s gocqlx.Session) *Scylla {
	return &Scylla{
		s,
		table.New(messageMetaData),
	}
}

// tables
var (
	messageMetaData = table.Metadata{
		Name:    "message",
		Columns: []string{"id", "value", "files", "sender_id", "recipient_id", "created_at", "updated_at"},
		PartKey: []string{"id"},
		SortKey: []string{"created_at"},
	}
)

type Scylla struct {
	gocqlx.Session

	message *table.Table
}

func (s Scylla) CreateMessage(ctx context.Context, msg models.Message) error {
	cql := fmt.Sprintf(`INSERT INTO %s (id, value, created_at) VALUES (?, ?, ?)`, s.message.Name())

	err := s.Query(cql, []string{msg.ID, msg.Value, msg.CreatedAt.String()}).WithContext(ctx).ExecRelease()
	if err != nil {
		return utils.NewError(utils.ErrInternal, err.Error())
	}

	log.GetLogger(ctx).Debug("repo passed")

	return nil
}

func (s Scylla) DeleteMessage(ctx context.Context, msgID string) error {
	cql := fmt.Sprintf(`DELETE FROM %s WHERE id = ?`, s.message.Name())

	err := s.Query(cql, []string{msgID}).WithContext(ctx).ExecRelease()
	if err != nil {
		return utils.NewError(utils.ErrInternal, err.Error())
	}

	log.GetLogger(ctx).Debug("repo passed")

	return nil
}

func (s Scylla) EditMessage(ctx context.Context, edit models.MessageEdit) error {
	cql := fmt.Sprintf(`UPDATE %s SET value = ?, updated_at = ? WHERE id = ?`, s.message.Name())

	err := s.Query(cql, []string{edit.Value, edit.UpdatedAt.String(), edit.MessageID}).WithContext(ctx).ExecRelease()
	if err != nil {
		return utils.NewError(utils.ErrInternal, err.Error())
	}

	log.GetLogger(ctx).Debug("repo passed")

	return nil
}
