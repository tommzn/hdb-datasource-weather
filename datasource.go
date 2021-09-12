package weather

import (
	config "github.com/tommzn/go-config"
	log "github.com/tommzn/go-log"
	secrets "github.com/tommzn/go-secrets"
	core "github.com/tommzn/hdb-datasource-core"
)

// New returns a new weather datasource
func New(config config.Config, logger log.Logger, secrestmanager secrets.SecretsManager) (core.DataSource, error) {
	return newWeatherApi(config, logger, secrestmanager)

}
