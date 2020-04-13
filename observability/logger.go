package observability

import (
	"github.com/Javin-Ambridge/go.base/go.base/constants"
	"github.com/Javin-Ambridge/go.base/go.base/entity"
	"github.com/Javin-Ambridge/go.base/go.base/utils/goutils"
	"go.uber.org/zap"
)

// NewLogger creates a new SugaredLogger for the server
func NewLogger(
	config entity.Config,
) (*zap.SugaredLogger, error) {
	var logger *zap.Logger
	if config.Env == constants.EnvProduction {
		var err error
		logger, err = zap.NewProduction()
		if err != nil {
			return nil, goutils.ErrWrap(err)
		}
	} else {
		var err error
		logger, err = zap.NewDevelopment()
		if err != nil {
			return nil, goutils.ErrWrap(err)
		}
	}

	defer logger.Sync()
	return logger.With(
		zap.String(constants.AppIDField, config.AppID),
	).Sugar(), nil
}
