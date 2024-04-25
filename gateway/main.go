package main

import (
	"github.com/alserov/hrs/gateway/internal/app"
	"github.com/alserov/hrs/gateway/internal/config"
)

func main() {
	app.MustStart(config.MustLoad())
}
