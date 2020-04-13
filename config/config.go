package config

import (
	"os"

	"go.uber.org/fx"

	"github.com/Javin-Ambridge/go.base/go.base/constants"

	"github.com/Javin-Ambridge/go.base/go.base/entity"
	"github.com/Javin-Ambridge/go.base/go.base/utils/goutils"
	"go.uber.org/config"
)

// Module exports all config resources to Fx at startup
var Module = fx.Provide(
	New,
)

// New creates a new Config for the server
func New() (entity.Config, error) {
	provider, err := config.NewYAML(config.File("./config/base.yaml"))
	if err != nil {
		return entity.Config{}, goutils.ErrWrap(err)
	}

	var conf entity.Config
	if err := provider.Get("").Populate(&conf); err != nil {
		return entity.Config{}, goutils.ErrWrap(err)
	}

	// If we have the SERVER_ENV variable set, and its production, we are in production, else development
	conf.Env = constants.EnvDevelopment
	if os.Getenv("SERVER_ENV") == constants.EnvProduction {
		conf.Env = constants.EnvProduction
	}

	return conf, nil
}
