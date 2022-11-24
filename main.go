package main

import (
	"github.com/Javin-Ambridge/go.base/app"
	"github.com/Javin-Ambridge/go.base/server"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		app.Module,
		server.Server,
	).Run()
}
