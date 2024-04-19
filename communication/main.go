package main

import (
	"github.com/alserov/hrs/communication/internal/app"
	"github.com/alserov/hrs/communication/internal/config"
)

func main() {
	app.MustStart(config.MustLoad())
}
