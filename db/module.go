package db

import (
	"github.com/Javin-Ambridge/go.base/go.base/db/postgres"
	"go.uber.org/fx"
)

// Module exports the DB resources to Fx at startup
var Module = fx.Provide(
	postgres.New,
)
