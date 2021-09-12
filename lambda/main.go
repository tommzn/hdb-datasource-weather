package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	config "github.com/tommzn/go-config"
	log "github.com/tommzn/go-log"
	secrets "github.com/tommzn/go-secrets"

	core "github.com/tommzn/hdb-datasource-core"
	weather "github.com/tommzn/hdb-datasource-weather"
)

func main() {

	collector, err := bootstrap()
	if err != nil {
		panic(err)
	}
	lambda.Start(collector.Run)
}

// bootstrap loads config and creates a new scheduled collector with a weather datasource.
func bootstrap() (core.Collector, error) {

	conf := loadConfig()
	secretsManager := newSecretsManager()
	logger := newLogger(conf, secretsManager)
	datasource, err := weather.New(conf, logger, secretsManager)
	if err != nil {
		return nil, err
	}

	queue := conf.Get("hdb.queue", config.AsStringPtr("de.tsl.hdb.weather"))
	return core.NewScheduledCollector(*queue, datasource, conf, logger), nil
}

// loadConfig from config file.
func loadConfig() config.Config {

	configSource, err := config.NewS3ConfigSourceFromEnv()
	if err != nil {
		panic(err)
	}

	conf, err := configSource.Load()
	if err != nil {
		panic(err)
	}
	return conf
}

// newSecretsManager retruns a new secrets manager from passed config.
func newSecretsManager() secrets.SecretsManager {
	return secrets.NewSecretsManager()
}

// newLogger creates a new logger from  passed config.
func newLogger(conf config.Config, secretsMenager secrets.SecretsManager) log.Logger {
	return log.NewLoggerFromConfig(conf, secretsMenager)
}
