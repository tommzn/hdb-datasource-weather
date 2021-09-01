package weather

import (
	core "github.com/tommzn/hdb-datasource-core"
	events "github.com/tommzn/hdb-events-go"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func New() core.DataSource {
	return &WeatherDataSource{}
}

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
