package weather

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"

	config "github.com/tommzn/go-config"
	log "github.com/tommzn/go-log"
	secrets "github.com/tommzn/go-secrets"
	events "github.com/tommzn/hdb-events-go"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// newWeatherApi creates a new OpenWeatherMap client with given config. If expects a secretsmanager which is
// able to obtain api key as OWM_API_KEY.
func newWeatherApi(conf config.Config, logger log.Logger, secretsmanager secrets.SecretsManager) (*OpenWeatherMapClient, error) {

	apiKey, err := secretsmanager.Obtain("OWM_API_KEY")
	if err != nil {
		return nil, err
	}

	owmUrl := conf.Get("weather.owm.url", nil)
	if owmUrl == nil {
		return nil, errors.New("Missing OpenWeatherMap url.")
	}

	latitude := conf.Get("weather.owm.latitude", nil)
	if latitude == nil {
		return nil, errors.New("Missing OpenWeatherMap location, latitude.")
	}

	longitude := conf.Get("weather.owm.longitude", nil)
	if longitude == nil {
		return nil, errors.New("Missing OpenWeatherMap location, longitude.")
	}

	units := conf.Get("weather.owm.units", nil)

	return &OpenWeatherMapClient{
		ownUrl:     *owmUrl,
		latitude:   *latitude,
		longitude:  *longitude,
		units:      units,
		apiKey:     *apiKey,
		httpClient: &http.Client{},
		logger:     logger,
	}, nil
}

// Fetch calls the OpenWeatherMap One Call api to get the current weather and a 7-days forecast.
func (client *OpenWeatherMapClient) Fetch() (proto.Message, error) {

	req := client.newRequestForOneCallApi()
	client.logger.Info("OWN URL: ", req.URL.String())

	res, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode > 399 {
		return nil, fmt.Errorf("Unexpected API response code: %d\n", res.StatusCode)
	}
	defer res.Body.Close()
	b, _ := ioutil.ReadAll(res.Body)

	var oneCallResponse openWeatherMapOneCallApiResponse
	json.Unmarshal(b, &oneCallResponse)
	return toWeatherDataEvent(oneCallResponse, client.units), nil
}

// newRequestForOneCallApi creates a new http GET request to get weather data from OpenWeatherMap One Call API.
// It ommits minutely/hourly weather data and alerts.
func (client *OpenWeatherMapClient) newRequestForOneCallApi() *http.Request {

	req, _ := http.NewRequest("GET", client.ownUrl+"/onecall", nil)

	query := req.URL.Query()
	query.Add("appid", client.apiKey)
	query.Add("lat", client.latitude)
	query.Add("lon", client.longitude)

	if client.units != nil {
		query.Add("units", *client.units)
	}

	// We're not interested in a minutely or hourly forecast in the moment and alerts are skipped as well
	query.Add("exclude", "minutely,hourly,alerts")

	req.URL.RawQuery = query.Encode()
	return req
}

// toWeatherDataEvent converts response from OpenWeatherMap One Call API to a weather data event.
func toWeatherDataEvent(oneCallResponse openWeatherMapOneCallApiResponse, units *string) *events.WeatherData {

	weatherData := events.WeatherData{
		Location: &events.Location{
			Longitude: oneCallResponse.Longitude,
			Latitude:  oneCallResponse.Latitude,
		},
		Current: &events.CurrentWeather{
			Timestamp:   asTimeStamp(oneCallResponse.Current.TimeStamp),
			Temperature: oneCallResponse.Current.Temperature,
			WindSpeed:   oneCallResponse.Current.WindSpeed,
			Weather:     toWeatherDetailsEventData(oneCallResponse.Current.Weather[0]),
		},
		Forecast: []*events.ForecastWeather{},
	}
	if units != nil {
		weatherData.Units = *units
	}
	for _, forecast := range oneCallResponse.DailyForcast {
		weatherData.Forecast = append(weatherData.Forecast, &events.ForecastWeather{
			Timestamp: asTimeStamp(forecast.TimeStamp),
			Temperatures: &events.ForecastTemperatures{
				Morning: forecast.Temperatures.Morning,
				Day:     forecast.Temperatures.Day,
				Evening: forecast.Temperatures.Evening,
				Night:   forecast.Temperatures.Night,
				DayMin:  forecast.Temperatures.DayMin,
				DayMax:  forecast.Temperatures.DayMax,
			},
			WindSpeed: forecast.WindSpeed,
			Weather:   toWeatherDetailsEventData(forecast.Weather[0]),
		})
	}
	return &weatherData
}

// toWeatherDetailsEventData converts given OpenWeatherMap weather information to event data.
func toWeatherDetailsEventData(weatherDetails weatherDetails) *events.WeatherDetails {
	return &events.WeatherDetails{
		ConditionId: weatherDetails.ConditionId,
		Group:       weatherDetails.Group,
		Description: weatherDetails.Description,
		Icon:        weatherDetails.Icon,
	}
}

// asTimeStamp converts a unix epoch timestamp to a Protobuf timestamp.
func asTimeStamp(epoch int64) *timestamppb.Timestamp {
	return timestamppb.New(time.Unix(epoch, 0))
}
