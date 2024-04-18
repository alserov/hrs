package grpc

import (
	"context"
	"github.com/alserov/hrs/communication/internal/usecase"
	comm "github.com/alserov/hrs/communication/pkg/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ comm.CommunicationServer = &Adapter{}

func NewAdapter(uc *usecase.Usecase) *grpc.Server {
	s := grpc.NewServer()
	comm.RegisterCommunicationServer(s, &Adapter{uc: uc})
	return s
}

type Adapter struct {
	comm.UnimplementedCommunicationServer

	uc *usecase.Usecase
}

func (a Adapter) CreateMessage(ctx context.Context, message *comm.Message) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (a Adapter) DeleteMessage(ctx context.Context, id *comm.MessageID) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (a Adapter) EditMessage(ctx context.Context, edit *comm.MessageEdit) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}
