package weather

import (
	config "github.com/tommzn/go-config"
	log "github.com/tommzn/go-log"
	secrets "github.com/tommzn/go-secrets"
	core "github.com/tommzn/hdb-datasource-core"
)

// New returns a new weather datasource
func New(config config.Config, secrestmanager secrets.SecretsManager, logger log.Logger) (core.DataSource, error) {
	return newWeatherApi(config, secrestmanager, logger)

}
