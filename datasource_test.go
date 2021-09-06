package weather

import (
	"testing"

	"github.com/stretchr/testify/suite"
	events "github.com/tommzn/hdb-events-go"
)

type DataSourceTestSuite struct {
	suite.Suite
}

func TestDataSourceTestSuite(t *testing.T) {
	suite.Run(t, new(DataSourceTestSuite))
}

func (suite *DataSourceTestSuite) TestFetch() {

	ds := New(loadConfigForTest())
	event, err := ds.Fetch()
	suite.Nil(err)
	weather, ok := event.(*events.Weather)
	suite.True(ok)
	suite.Equal(int64(27), weather.Temperature)
	suite.Equal(int64(3), weather.Wind)
	suite.Equal("Sunny", weather.Description)
	suite.Equal(432.654, weather.Location.Longitude)
	suite.Equal(321.764, weather.Location.Latitude)
}
