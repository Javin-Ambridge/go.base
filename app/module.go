package app

import (
	"github.com/Javin-Ambridge/go.base/config"
	"github.com/Javin-Ambridge/go.base/controller"
	"github.com/Javin-Ambridge/go.base/db"
	"github.com/Javin-Ambridge/go.base/gateway"
	"github.com/Javin-Ambridge/go.base/handler"
	"github.com/Javin-Ambridge/go.base/observability"
	"github.com/Javin-Ambridge/go.base/secrets"
	"go.uber.org/fx"
)

// Module gets exported to Fx
var Module = fx.Options(
	observability.Module,
	handler.Module,
	config.Module,
	controller.Module,
	db.Module,
	gateway.Module,
	secrets.Module,
)
