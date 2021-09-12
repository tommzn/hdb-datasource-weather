package weather

import (
	"net/http"

	log "github.com/tommzn/go-log"
)

//
// Open Weather Map - One Call API
// Response types, unused attributes are skipped
// Details at: https://openweathermap.org/api/one-call-api
//

// openWeatherMapOneCallApiResponse contains full response from Open Weather Map One Call API
type openWeatherMapOneCallApiResponse struct {
	Longitude      float64               `json:"lat"`             // Geographical coordinates of the location (latitude)
	Latitude       float64               `json:"lon"`             // Geographical coordinates of the location (longitude)
	TimeZone       string                `json:"timezone"`        // Timezone name for the requested location
	TimeZoneOffset int                   `json:"timezone_offset"` // Shift in seconds from UTC
	Current        currentWeatherData    `json:"current"`         // Current weather data API response
	DailyForcast   []forecastWeatherData `json:"daily"`           // 7-days forecast weather data
}

// currentWeatherData contains current weather information
type currentWeatherData struct {
	TimeStamp   int64            `json:"dt"`         // Current time, Unix (Epoch), UTC
	Temperature float64          `json:"temp"`       // Temperature. Units - default: kelvin, metric: Celsius, imperial: Fahrenheit
	WindSpeed   float64          `json:"wind_speed"` // Wind speed. Wind speed. Units – default: metre/sec, metric: metre/sec, imperial: miles/hour
	Weather     []weatherDetails `json:"weather"`    // Weather details
}

// weatherDetails contains information of current weather or forecast
// Full list of condition id, group and description and Icons is available at: https://openweathermap.org/weather-conditions#Weather-Condition-Codes-2
type weatherDetails struct {
	ConditionId int64  `json:"id"`          // Weather condition id
	Group       string `json:"main"`        // Group of weather parameters (Rain, Snow, Extreme etc.)
	Description string `json:"description"` // Weather condition within the group
	Icon        string `json:"icon"`        // Weather icon id. How to get icons
}

// forecastWeatherData contains forecast weather data for a single day
type forecastWeatherData struct {
	TimeStamp    int64                `json:"dt"`         // Current time, Unix (Epoch), UTC
	Temperatures forecastTemperatures `json:"temp"`       // Forecast temperatures for the whole day
	WindSpeed    float64              `json:"wind_speed"` // Wind speed. Wind speed. Units – default: metre/sec, metric: metre/sec, imperial: miles/hour
	Weather      []weatherDetails     `json:"weather"`    // Weather details
}

// forecastTemperatures contains forecast temperature data for a single day
// Units – default: kelvin, metric: Celsius, imperial: Fahrenheit
type forecastTemperatures struct {
	Morning float64 `json:"morn"`  // Morning temperature
	Day     float64 `json:"day"`   // Day temperature
	Evening float64 `json:"eve"`   // Evening temperature
	Night   float64 `json:"night"` // Night temperature
	DayMin  float64 `json:"min"`   // Min daily temperature.
	DayMax  float64 `json:"max"`   // Max daily temperature.
}

// OpenWeatherMapClient handles requests for current weather and a 7-days forecast to API provides by Open Weather Map.
// It can be used as a datasource directly.
type OpenWeatherMapClient struct {
	ownUrl     string
	longitude  string
	latitude   string
	apiKey     string
	units      *string
	httpClient *http.Client
	logger     log.Logger
}
