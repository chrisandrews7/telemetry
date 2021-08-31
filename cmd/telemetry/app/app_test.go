package app_test

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"os"
	"testing"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/stretchr/testify/assert"
	"github.com/telemetry/cmd/telemetry/app"
)

func TestInsertTelemetry(t *testing.T) {
	// Mock string telemetry
	sTel, err := net.Listen("tcp", ":3333")
	if err != nil {
		t.Fatal(err)
	}
	defer sTel.Close()

	// Mock binary telemetry
	bTel, err := net.Listen("tcp", ":4444")
	if err != nil {
		t.Fatal(err)
	}
	defer bTel.Close()

	// Start App
	config := app.Config{
		InfluxAddress:              os.Getenv("APP_INFLUX_ADDRESS"),
		InfluxCredentials:          os.Getenv("APP_INFLUX_CREDENTIALS"),
		InfluxDBName:               os.Getenv("APP_INFLUX_DB"),
		GroundStationStringAddress: sTel.Addr().String(),
		GroundStationBinaryAddress: bTel.Addr().String(),
	}
	close, err := app.Run(config)
	assert.NoError(t, err)
	defer close()

	client := influxdb2.NewClient(config.InfluxAddress, config.InfluxCredentials)
	defer client.Close()
	queryAPI := client.QueryAPI("")

	// Send string telemetry data
	sConn, err := sTel.Accept()
	if err != nil {
		return
	}
	rand.Seed(time.Now().UnixNano())
	sTelemetryID := rand.Intn(1000)
	sConn.Write([]byte(fmt.Sprintf("[%d:%d:12345.000000]", time.Now().Unix(), sTelemetryID)))
	sConn.Close()

	// Send binary telemetry data
	bConn, err := bTel.Accept()
	if err != nil {
		return
	}
	bTelemetryID := sTelemetryID + 1
	bConn.Write([]byte(fmt.Sprintf("%d%d", time.Now().Unix(), bTelemetryID)))
	bConn.Close()

	// Query the parsed telemetry
	result, err := queryAPI.Query(
		context.Background(),
		fmt.Sprintf(`from(bucket:"%s") |> range(start: -1m) |> sort(columns: ["_time"], desc: true) |> limit(n:2)`, config.InfluxDBName),
	)
	assert.NoError(t, err)

	result.Next()
	assert.NoError(t, result.Err())
	assert.NotNil(t, result.Record())
	assert.Equal(t, 12345.000000, result.Record().Value())
}
