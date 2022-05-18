package repo

import (
	"context"
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	influxdb2API "github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/pkg/errors"
	"github.com/telemetry/internal/point"
)

const MEASUREMENT_NAME = "telemetry"

//go:generate go run github.com/vektra/mockery/v2@latest --dir ${GOBASE}/vendor/github.com/influxdata/influxdb-client-go/v2 --name Client
//go:generate go run github.com/vektra/mockery/v2@latest --dir ${GOBASE}/vendor/github.com/influxdata/influxdb-client-go/v2/api --name WriteAPIBlocking
type InfluxRepo struct {
	client influxdb2.Client
	api    influxdb2API.WriteAPIBlocking
}

func (r *InfluxRepo) Insert(ctx context.Context, p *point.Point) error {
	record := influxdb2.NewPoint(MEASUREMENT_NAME,
		map[string]string{"telemetryID": fmt.Sprint(p.TelemetryID)},
		map[string]interface{}{"value": p.Value},
		time.Unix(p.Timestamp, 0))

	err := r.api.WritePoint(ctx, record)
	if err != nil {
		return errors.Wrap(err, "Unable to insert point")
	}

	return nil
}

func (r *InfluxRepo) Close() {
	r.client.Close()
}

func NewInfluxRepo(host, credentials, dbName string) *InfluxRepo {
	client := influxdb2.NewClient(host, credentials)

	return NewInfluxRepoWithClient(
		client,
		dbName,
	)
}

func NewInfluxRepoWithClient(client influxdb2.Client, dbName string) *InfluxRepo {
	api := client.WriteAPIBlocking("", dbName)

	return &InfluxRepo{
		client,
		api,
	}
}
