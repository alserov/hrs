package grpc

import (
	"context"
	"github.com/alserov/hrs/communication/internal/log"
	"github.com/alserov/hrs/communication/internal/usecase"
	"github.com/alserov/hrs/communication/internal/utils"
	comm "github.com/alserov/hrs/communication/pkg/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ comm.CommunicationServer = &Adapter{}

func NewAdapter(uc *usecase.Usecase, log log.Logger) *grpc.Server {
	s := grpc.NewServer()
	comm.RegisterCommunicationServer(s, &Adapter{uc: uc, log: log, conv: utils.NewConverter()})
	return s
}

type Adapter struct {
	comm.UnimplementedCommunicationServer

	log  log.Logger
	uc   *usecase.Usecase
	conv *utils.Converter
}

func (a *Adapter) CreateMessage(ctx context.Context, message *comm.Message) (*emptypb.Empty, error) {
	a.log.Debug("received request")

	ctx = log.WithLogger(ctx, a.log)

	err := a.uc.CreateMessage(ctx, a.conv.ToMessage(message))
	if err != nil {
		return nil, utils.FromError(err)
	}

	return &emptypb.Empty{}, nil
}

func (a *Adapter) DeleteMessage(ctx context.Context, id *comm.MessageID) (*emptypb.Empty, error) {
	a.log.Debug("received request")

	ctx = log.WithLogger(ctx, a.log)

	err := a.uc.DeleteMessage(ctx, id.ID)
	if err != nil {
		return nil, utils.FromError(err)
	}

	return &emptypb.Empty{}, nil
}

func (a *Adapter) EditMessage(ctx context.Context, edit *comm.MessageEdit) (*emptypb.Empty, error) {
	a.log.Debug("received request")

	ctx = log.WithLogger(ctx, a.log)

	err := a.uc.EditMessage(ctx, a.conv.ToMessageEdit(edit))
	if err != nil {
		return nil, utils.FromError(err)
	}

	return &emptypb.Empty{}, nil
}
