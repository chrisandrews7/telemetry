package main

import (
	"os"
	"os/signal"

	"github.com/telemetry/cmd/telemetry/app"
)

func main() {
	close, err := app.Run(app.Config{
		InfluxAddress:              os.Getenv("APP_INFLUX_ADDRESS"),
		InfluxCredentials:          os.Getenv("APP_INFLUX_CREDENTIALS"),
		InfluxDBName:               os.Getenv("APP_INFLUX_DB"),
		GroundStationStringAddress: os.Getenv("APP_GROUNDSTATION_STRING"),
		GroundStationBinaryAddress: os.Getenv("APP_GROUNDSTATION_BINARY"),
	})
	if err != nil {
		return
	}

	isTerminated := make(chan os.Signal)
	signal.Notify(isTerminated, os.Interrupt)
	<-isTerminated
	close()
}
