package weather

import (
	"os"

	config "github.com/tommzn/go-config"
	log "github.com/tommzn/go-log"
	secrets "github.com/tommzn/go-secrets"
)

// loadConfigForTest loads test config.
func loadConfigForTest() config.Config {

	configFile, ok := os.LookupEnv("CONFIG_FILE")
	if !ok {
		configFile = "testconfig.yml"
	}
	configLoader := config.NewFileConfigSource(&configFile)
	config, _ := configLoader.Load()
	return config
}

// secretsManagerForTest returns a default secrets manager for testing
func secretsManagerForTest() secrets.SecretsManager {
	return secrets.NewSecretsManager()
}

// loggerForTest creates a new stdout logger for testing.
func loggerForTest() log.Logger {
	return log.NewLogger(log.Debug, nil, nil)
}
