package weather

import (
	"os"

	config "github.com/tommzn/go-config"
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
