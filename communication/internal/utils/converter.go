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
	return models.Message{}
}

func (*Converter) ToMessageEdit(msg *comm.MessageEdit) models.MessageEdit {
	return models.MessageEdit{}
}
