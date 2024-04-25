package clients

import "github.com/alserov/hrs/gateway/internal/clients/grpc"

type Clients struct {
	*grpc.Clients
}

func NewClients() *Clients {
	return &Clients{
		grpc.NewClients(),
	}
}
