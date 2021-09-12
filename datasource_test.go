package weather

import (
	"testing"
	"time"

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

	dateFormat := "20060102"
	ds, err := New(loadConfigForTest(), loggerForTest(), secretsManagerForTest())
	suite.Nil(err)
	suite.NotNil(ds)

	ownClient, ok := ds.(*OpenWeatherMapClient)
	suite.True(ok)
	suite.NotEqual("", ownClient.ownUrl)

	event, err := ds.Fetch()
	suite.Nil(err)
	suite.NotNil(event)

	weatherData, ok := event.(*events.WeatherData)
	suite.True(ok)
	suite.Equal(time.Now().UTC().Format(dateFormat), weatherData.Current.Timestamp.AsTime().Format(dateFormat))
	suite.True(len(weatherData.Forecast) >= 7)
}
