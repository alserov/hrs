package utils

import (
	comm "github.com/alserov/hrs/communication/pkg/proto/gen"
	"github.com/alserov/hrs/gateway/internal/models"
)

type Converter struct{}

func (c *Converter) ToMessage(in models.Message) *comm.Message {
	return &comm.Message{}
}
