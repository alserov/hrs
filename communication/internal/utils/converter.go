package utils

import (
	"github.com/alserov/hrs/communication/internal/usecase/models"
	comm "github.com/alserov/hrs/communication/pkg/proto/gen"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func NewConverter() *Converter {
	return &Converter{}
}

type Converter struct {
}

func (*Converter) ToMessage(msg *comm.Message) models.Message {
	return models.Message{
		Value:       msg.Value,
		Files:       msg.Files,
		SenderID:    msg.SenderID,
		RecipientID: msg.RecipientID,
	}
}

func (*Converter) ToMessageEdit(msg *comm.MessageEdit) models.MessageEdit {
	return models.MessageEdit{
		Value:     msg.Value,
		MessageID: msg.ID,
	}
}

func (*Converter) ToHistory(msg *comm.HistoryParam) models.HistoryParam {
	return models.HistoryParam{
		UserID:      msg.UserID,
		RecipientID: msg.RecipientID,
		Limit:       msg.Limit,
		Offset:      msg.Offset,
	}
}

func (*Converter) ToChatList(msg *comm.ChatsParam) models.ChatsParam {
	return models.ChatsParam{
		UserID: msg.UserID,
		Limit:  msg.Limit,
		Offset: msg.Offset,
	}
}

func (*Converter) FromMessages(msgs []models.Message) *comm.Messages {
	m := make([]*comm.Message, 0, len(msgs))
	for _, msg := range msgs {
		m = append(m, &comm.Message{
			ID:          msg.ID,
			Value:       msg.Value,
			Files:       msg.Files,
			SenderID:    msg.SenderID,
			RecipientID: msg.RecipientID,
			CreatedAt:   toProtoTime(msg.CreatedAt),
			UpdatedAt:   toProtoTime(msg.UpdatedAt),
		})
	}

	return &comm.Messages{Messages: m}
}

func (*Converter) FromChats(chats []models.Chat) *comm.Chats {
	c := make([]*comm.Chat, 0, len(chats))
	for _, chat := range chats {
		c = append(c, &comm.Chat{
			LastMessage: &comm.Message{
				ID:          chat.LastMessage.ID,
				Value:       chat.LastMessage.Value,
				Files:       chat.LastMessage.Files,
				SenderID:    chat.LastMessage.SenderID,
				RecipientID: chat.LastMessage.RecipientID,
				CreatedAt:   toProtoTime(chat.LastMessage.CreatedAt),
				UpdatedAt:   toProtoTime(chat.LastMessage.UpdatedAt),
			},
		})
	}

	return &comm.Chats{Chats: c}
}

func toProtoTime(t time.Time) *timestamppb.Timestamp {
	return &timestamppb.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
	}
}
