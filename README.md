[![Go Reference](https://pkg.go.dev/badge/github.com/tommzn/hdb-datasource-weather.svg)](https://pkg.go.dev/github.com/tommzn/hdb-datasource-weather)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/tommzn/hdb-datasource-weather)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/tommzn/hdb-datasource-weather)
[![Go Report Card](https://goreportcard.com/badge/github.com/tommzn/hdb-datasource-weather)](https://goreportcard.com/report/github.com/tommzn/hdb-datasource-weather)
[![Actions Status](https://github.com/tommzn/hdb-datasource-weather/actions/workflows/go.pkg.auto-ci.yml/badge.svg)](https://github.com/tommzn/hdb-datasource-weather/actions)

# HomeDashboard Weather DataSource
Fetches weather data from Open Weather Map API using it's One Call API to get current weather and a 7-days forcast.

## Config
You have to pass the OpenWeatherMap API url, together with a geographical location (logitude, latitude) as config. Optional config param are units.
More details about loading config at https://github.com/tommzn/go-config

### Config example with all required values
```yaml
weather:
  owm:
    url: https://api.openweathermap.org/data/2.5
    longitude: 37.33195305634116 
    latitude: -122.0309010022451
```

### Full Config
```yaml
weather:
  owm:
    url: https://api.openweathermap.org/data/2.5
    longitude: 37.33195305634116 
    latitude: -122.0309010022451
    units: metric
```

## Secrets
OpenWeatherMap requires an API key to call the One Call API. You have to pass a [SecretsManager](https://github.com/tommzn/go-secrets) which is able to obtain an API key as OWM_API_KEY

## Get Weather Data
After creating a new datasource, you can fetch current weather data and a forecast. If anything works well Fetch will return a [weather data struct](https://github.com/tommzn/hdb-events-go/blob/main/weather.pb.go) or otherwise an error.
```golang

    datasource, err := New(config, secretsmanager)
    if err != nil {
        panic(err)
    }

    weatherData, err := datasource.Fetch()
    if err != nil {
        panic(err)
    }

    fmt.Printf("Current Temperature: %.2f\n", weatherData.(events.WeatherData).Current.Temperature)
```
