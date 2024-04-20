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
	"golang.org/x/sync/errgroup"
	"strconv"
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

func (s Scylla) GetHistory(ctx context.Context, param models.HistoryParam) ([]models.Message, error) {
	cql := fmt.Sprintf(`
					SELECT * FROM %s 
         			WHERE 
         			   ( sender_id = ? AND recipient_id = ?)
         			    OR
         			    (recipient_id = ? AND sender_id = ?)
         			    ORDER BY created_at desc LIMIT ? OFFSET ?`, s.message.Name())

	query := s.Query(cql, []string{param.UserID, param.RecipientID, param.RecipientID, param.UserID, strconv.Itoa(int(param.Limit)), strconv.Itoa(int(param.Offset))})

	if err := query.Err(); err != nil {
		return nil, utils.NewError(utils.ErrInternal, err.Error())
	}

	var msgs []models.Message

	iter := query.Iter()
	for {
		var msg models.Message
		if !iter.Scan(
			&msg.ID,
			&msg.Value,
			&msg.Files,
			&msg.SenderID,
			&msg.RecipientID,
			&msg.CreatedAt,
			&msg.UpdatedAt) {
			break
		}
		msgs = append(msgs, msg)
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}

	log.GetLogger(ctx).Debug("repo passed")

	return msgs, nil
}

func (s Scylla) GetChats(ctx context.Context, param models.ChatsParam) ([]models.Chat, error) {
	var (
		chChats = make(chan models.Chat, 3)
		eg      = errgroup.Group{}
	)
	defer close(chChats)

	eg.Go(func() error {
		cql := fmt.Sprintf(`SELECT DISTINCT recipient_id, id, value, files, sender_id, created_at, updated_at 
			FROM %s WHERE sender_id = ? ORDER BY created_at desc LIMIT ? OFFSET ?`, s.message.Name())

		query := s.Query(cql, []string{param.UserID, string(param.Limit), string(param.Offset)})

		if err := query.Err(); err != nil {
			return utils.NewError(utils.ErrInternal, err.Error())
		}

		iter := query.Iter()
		for {
			var chat models.Chat
			if !iter.Scan(
				&chat.LastMessage.ID,
				&chat.LastMessage.Value,
				&chat.LastMessage.Files,
				&chat.LastMessage.SenderID,
				&chat.LastMessage.RecipientID,
				&chat.LastMessage.CreatedAt,
				&chat.LastMessage.UpdatedAt) {
				break
			}
			chChats <- chat
		}
		if err := iter.Close(); err != nil {
			return err
		}

		return nil
	})

	eg.Go(func() error {
		cql := fmt.Sprintf(`SELECT DISTINCT sender_id, id, value, files, recipient_id, created_at, updated_at 
			FROM %s WHERE recipient_id = ? ORDER BY created_at desc LIMIT ? OFFSET ?`, s.message.Name())

		query := s.Query(cql, []string{param.UserID, string(param.Limit), string(param.Offset)})

		if err := query.Err(); err != nil {
			return utils.NewError(utils.ErrInternal, err.Error())
		}

		iter := query.Iter()
		for {
			var chat models.Chat
			if !iter.Scan(
				&chat.LastMessage.ID,
				&chat.LastMessage.Value,
				&chat.LastMessage.Files,
				&chat.LastMessage.SenderID,
				&chat.LastMessage.RecipientID,
				&chat.LastMessage.CreatedAt,
				&chat.LastMessage.UpdatedAt) {
				break
			}
			chChats <- chat
		}
		if err := iter.Close(); err != nil {
			return err
		}

		return nil
	})

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	var chats []models.Chat
	for ch := range chChats {
		chats = append(chats, ch)
	}

	log.GetLogger(ctx).Debug("repo passed")

	return chats, nil
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
