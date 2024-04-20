package grpc

import (
	"context"
	"fmt"
	"github.com/alserov/hrs/communication/internal/config"
	"github.com/alserov/hrs/communication/internal/log"
	"github.com/alserov/hrs/communication/internal/mocks"
	"github.com/alserov/hrs/communication/internal/usecase"
	comm "github.com/alserov/hrs/communication/pkg/proto/gen"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"net"
	"testing"
)

var (
	port = 10000
	l    net.Listener
	err  error
	conn grpc.ClientConnInterface
)

func TestMySuite(t *testing.T) {
	suite.Run(t, new(gRPCSuite))
}

type gRPCSuite struct {
	suite.Suite
}

func (s *gRPCSuite) SetupTest() {
	l, err = net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
}

func (s *gRPCSuite) TearDownTest() {
	l.Close()
}

func (s *gRPCSuite) TestAdapter_CreateMessage() {
	msg := comm.Message{
		Value:       "value",
		Files:       []string{"file1", "file2"},
		SenderID:    "senderID",
		RecipientID: "recipientID",
	}

	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()

	repo := mocks.NewMockRepository(ctrl)
	repo.EXPECT().
		CreateMessage(gomock.Any(), gomock.Any()).
		Times(1).
		Return(nil)

	uc := usecase.NewUsecase(repo)

	srvr := NewAdapter(uc, log.NewLogger(config.Local))

	go srvr.Serve(l)

	// ===============

	res, err := initGRPCClient(s.T()).CreateMessage(context.Background(), &msg)
	require.NotEmpty(s.T(), res)
	require.NoError(s.T(), err)
}

func (s *gRPCSuite) TestAdapter_EditMessage() {
	edit := comm.MessageEdit{
		Value: "value",
		ID:    "id",
	}

	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()

	repo := mocks.NewMockRepository(ctrl)
	repo.EXPECT().
		EditMessage(gomock.Any(), gomock.Any()).
		Times(1).
		Return(nil)

	uc := usecase.NewUsecase(repo)

	srvr := NewAdapter(uc, log.NewLogger(config.Local))

	go srvr.Serve(l)

	// ===============

	res, err := initGRPCClient(s.T()).EditMessage(context.Background(), &edit)
	require.NotEmpty(s.T(), res)
	require.NoError(s.T(), err)
}

func (s *gRPCSuite) TestAdapter_DeleteMessage() {
	id := comm.MessageID{
		ID: "id",
	}

	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()

	repo := mocks.NewMockRepository(ctrl)
	repo.EXPECT().
		DeleteMessage(gomock.Any(), gomock.Eq(id.ID)).
		Times(1).
		Return(nil)

	uc := usecase.NewUsecase(repo)

	srvr := NewAdapter(uc, log.NewLogger(config.Local))

	go srvr.Serve(l)

	// ===============

	res, err := initGRPCClient(s.T()).DeleteMessage(context.Background(), &id)
	require.NotEmpty(s.T(), res)
	require.NoError(s.T(), err)
}

func initGRPCClient(t *testing.T) comm.CommunicationClient {
	conn, err = grpc.Dial(fmt.Sprintf("127.0.0.1:%d", port), grpc.WithInsecure())
	require.NoError(t, err)

	return comm.NewCommunicationClient(conn)
}
