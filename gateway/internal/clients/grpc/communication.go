package grpc

import (
	comm "github.com/alserov/hrs/communication/pkg/proto/gen"
	"github.com/alserov/hrs/gateway/internal/utils"
	gRPC "google.golang.org/grpc"
	"time"
)

const (
	refreshPeriod            time.Duration = 3 * time.Second
	CommunicationServiceName               = "communication"
)

type CommunicationServiceClient interface {
	comm.CommunicationClient
}

func NewCommunicationServiceClient(balancing utils.Strategy) CommunicationServiceClient {
	cl := &communicationServiceClient{
		b: utils.NewBalancer(balancing),
	}

	go cl.refresh()

	return cl
}

type communicationServiceClient struct {
	comm.CommunicationClient

	b *utils.Balancer
}

func (ccl *communicationServiceClient) refresh() {
	t := time.NewTicker(refreshPeriod)
	defer t.Stop()

	var (
		conn gRPC.ClientConnInterface
		err  error
	)

	for range t.C {
		//addrs, _ := utils.Discover(CommunicationServiceName)
		addrs := []string{"123"}
		ccl.b.Set(addrs)

		conn, err = dial(ccl.b.Get())
		for err != nil {
			conn, err = dial(ccl.b.Get())
		}

		ccl.CommunicationClient = comm.NewCommunicationClient(conn)
	}
}
