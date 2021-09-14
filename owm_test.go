package weather

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	config "github.com/tommzn/go-config"
	log "github.com/tommzn/go-log"
)

type OpenWeatherMapTestSuite struct {
	suite.Suite
	conf   config.Config
	logger log.Logger
}

func TestOpenWeatherMapTestSuite(t *testing.T) {
	suite.Run(t, new(OpenWeatherMapTestSuite))
}

func (suite *OpenWeatherMapTestSuite) SetupTest() {
	suite.conf = loadConfigForTest()
	suite.logger = loggerForTest()
}

func (suite *OpenWeatherMapTestSuite) TestCreateWeatherApiClient() {

	apiKeyEnv := "OWM_API_KEY"
	apiKey := os.Getenv(apiKeyEnv)
	os.Unsetenv(apiKeyEnv)

	ds1, err1 := New(loadConfigForTest(), secretsManagerForTest())
	suite.NotNil(err1)
	suite.Equal("Secret not found: OWM_API_KEY", err1.Error())
	suite.Nil(ds1)

	os.Setenv(apiKeyEnv, apiKey)

	yamlConfig := ""
	config2, err := config.NewStaticConfigSource(yamlConfig).Load()
	suite.Nil(err)
	ds2, err2 := New(config2, secretsManagerForTest())
	suite.NotNil(err2)
	suite.Equal("Missing OpenWeatherMap url.", err2.Error())
	suite.Nil(ds2)

	yamlConfig = yamlConfig + "weather.owm.url: https://test.example.com" + "\n"
	config3, err := config.NewStaticConfigSource(yamlConfig).Load()
	suite.Nil(err)
	ds3, err3 := New(config3, secretsManagerForTest())
	suite.NotNil(err3)
	suite.Equal("Missing OpenWeatherMap location, latitude.", err3.Error())
	suite.Nil(ds3)

	yamlConfig = yamlConfig + "weather.owm.latitude: 123.456" + "\n"
	config4, err := config.NewStaticConfigSource(yamlConfig).Load()
	suite.Nil(err)
	ds4, err4 := New(config4, secretsManagerForTest())
	suite.NotNil(err4)
	suite.Equal("Missing OpenWeatherMap location, longitude.", err4.Error())
	suite.Nil(ds4)

	yamlConfig = yamlConfig + "weather.owm.longitude: 123.456" + "\n"
	config5, err := config.NewStaticConfigSource(yamlConfig).Load()
	suite.Nil(err)
	ds5, err5 := New(config5, secretsManagerForTest())
	suite.Nil(err5)
	suite.NotNil(ds5)
}

func (suite *OpenWeatherMapTestSuite) TestTimeStampConversion() {

	now := time.Now()
	timeStamp := asTimeStamp(now.Unix())
	suite.NotNil(timeStamp)

	suite.Equal(now.UTC().Format(time.RFC3339), timeStamp.AsTime().Format(time.RFC3339))
}

func (suite *OpenWeatherMapTestSuite) TestConvertApiResponse() {

	apiResponse := openWeatherMapOneCallApiResponseForTest()
	units := "metric"

	weatherData := toWeatherDataEvent(apiResponse, &units)
	suite.NotNil(weatherData)
	suite.Equal(21.4, weatherData.Current.Temperature)
	suite.Len(weatherData.Forecast, 3)
}

func (suite *OpenWeatherMapTestSuite) TestWithResponseError() {

	apiKeyEnv := "OWM_API_KEY"
	apiKey := os.Getenv(apiKeyEnv)
	os.Setenv(apiKeyEnv, "xxx")

	ds, err := New(loadConfigForTest(), secretsManagerForTest())
	suite.Nil(err)
	suite.NotNil(ds)

	event, err := ds.Fetch()
	suite.NotNil(err)
	suite.Nil(event)

	os.Setenv(apiKeyEnv, apiKey)

	yamlConfig := "weather.owm.url: https://test.example.com" + "\n"
	yamlConfig = yamlConfig + "weather.owm.latitude: 123.456" + "\n"
	yamlConfig = yamlConfig + "weather.owm.longitude: 123.456" + "\n"
	config1, err1 := config.NewStaticConfigSource(yamlConfig).Load()
	suite.Nil(err1)
	ds1, err1 := New(config1, secretsManagerForTest())
	event1, err1 := ds1.Fetch()
	suite.NotNil(err1)
	suite.Nil(event1)
}

func openWeatherMapOneCallApiResponseForTest() openWeatherMapOneCallApiResponse {
	apiResponse := openWeatherMapOneCallApiResponse{
		Longitude:      23.545,
		Latitude:       23.667,
		TimeZone:       "Europe/Berlin",
		TimeZoneOffset: 3600,
		Current: currentWeatherData{
			TimeStamp:   time.Now().Unix(),
			Temperature: 21.4,
			WindSpeed:   3.45,
			Weather: []weatherDetails{
				weatherDetails{
					ConditionId: 100,
					Group:       "Rain",
					Description: "Light Rain",
					Icon:        "10d",
				}},
		},
		DailyForcast: []forecastWeatherData{},
	}
	for i := 1; i <= 3; i++ {
		apiResponse.DailyForcast = append(apiResponse.DailyForcast, forecastWeatherData{
			TimeStamp: time.Now().Add(time.Duration(i) * 24 * time.Hour).Unix(),
			Temperatures: forecastTemperatures{
				Morning: 11.11,
				Day:     22.22,
				Evening: 15.15,
				Night:   10.10,
				DayMin:  21.21,
				DayMax:  23.23,
			},
			WindSpeed: 34.54,
			Weather: []weatherDetails{
				weatherDetails{
					ConditionId: 100,
					Group:       "Rain",
					Description: "Light Rain",
					Icon:        "10d",
				}},
		})
	}
	return apiResponse
}
