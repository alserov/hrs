package utils

import (
	"github.com/alserov/hrs/communication/internal/usecase/models"
	comm "github.com/alserov/hrs/communication/pkg/proto/gen"
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
