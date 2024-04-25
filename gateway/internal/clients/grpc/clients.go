package grpc

import (
	comm "github.com/alserov/hrs/communication/pkg/proto/gen"
	"github.com/alserov/hrs/gateway/internal/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Clients struct {
	Comm comm.CommunicationClient
}

func NewClients() *Clients {
	return &Clients{
		Comm: NewCommunicationServiceClient(utils.Random),
	}
}

func dial(addr string) (grpc.ClientConnInterface, error) {
	return grpc.Dial(addr, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
}
