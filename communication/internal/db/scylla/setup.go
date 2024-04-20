package scylla

import (
	"github.com/alserov/hrs/communication/internal/config"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

func MustConnect(cfg config.DB) gocqlx.Session {
	cluster := gocql.NewCluster(cfg.Host)
	cluster.Port = cfg.Port

	ws, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		panic("failed to create session: " + err.Error())
	}

	return ws
}

type M map[string]any
