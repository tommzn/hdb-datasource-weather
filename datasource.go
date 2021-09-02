package weather

import (
	config "github.com/tommzn/go-config"
	core "github.com/tommzn/hdb-datasource-core"
	events "github.com/tommzn/hdb-events-go"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// New returns a new weather datasource
func New(config config.Config) core.DataSource {
	return &WeatherDataSource{}
}

// Fetch returns current weather for coordinates defined by given config
func (datasource *WeatherDataSource) Fetch() (interface{}, error) {
	return &events.Weather{
		Timestamp:   timestamppb.Now(),
		Description: "Sunny",
		Temperature: 27,
		Wind:        3,
		Location: &events.Location{
			Longitude: 432.654,
			Latitude:  321.764,
		},
	}, nil
}
