package app

import (
	"github.com/Javin-Ambridge/go.base/go.base/config"
	"github.com/Javin-Ambridge/go.base/go.base/controller"
	"github.com/Javin-Ambridge/go.base/go.base/db"
	"github.com/Javin-Ambridge/go.base/go.base/gateway"
	"github.com/Javin-Ambridge/go.base/go.base/handler"
	"github.com/Javin-Ambridge/go.base/go.base/observability"
	"github.com/Javin-Ambridge/go.base/go.base/secrets"
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
