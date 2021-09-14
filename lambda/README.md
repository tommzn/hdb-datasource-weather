![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/tommzn/hdb-datasource-weather/lambda/go.mod)
[![Actions Status](https://github.com/tommzn/hdb-datasource-weather/actions/workflows/go.image.build.yml/badge.svg)](https://github.com/tommzn/hdb-datasource-weather/actions)

# Weather Data Collector
This package composes a [data collector](https://github.com/tommzn/hdb-datasource-core/collector.go) and [weather data source](https://github.com/tommzn/hdb-datasource-weather) to fetch weather data and publish it to a SQS queue.