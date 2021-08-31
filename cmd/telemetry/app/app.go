package app

import (
	"context"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/telemetry/internal/parsers"
	"github.com/telemetry/internal/repo"
)

func worker(ctx context.Context, wg *sync.WaitGroup, fn func() error) {
	for {
		select {
		default:
			wg.Add(1)
			defer wg.Done()

			err := fn()
			if err != nil {
				continue
			}

		case <-ctx.Done():
			return
		}
	}
}

type Config struct {
	InfluxAddress     string
	InfluxCredentials string
	InfluxDBName      string

	GroundStationStringAddress string
	GroundStationBinaryAddress string
}

func Run(config Config) (func(), error) {
	repository := repo.NewInfluxRepo(config.InfluxAddress, config.InfluxCredentials, config.InfluxDBName)
	wg := sync.WaitGroup{}
	ctx := context.Background()

	// String format ground station
	sConn, err := connect(config.GroundStationStringAddress)
	if err != nil {
		return func() {}, err
	}
	sCtx, sCancel := context.WithCancel(ctx)
	addStringTelemetry := func() error {
		return addTelemetry(ctx, sConn, parsers.NewStringParser(), repository)
	}
	go worker(sCtx, &wg, addStringTelemetry)

	// Binary format ground station
	bConn, err := connect(config.GroundStationBinaryAddress)
	if err != nil {
		return func() {}, err
	}
	bCtx, bCancel := context.WithCancel(ctx)
	addBinaryTelemetry := func() error {
		return addTelemetry(ctx, bConn, parsers.NewBinaryParser(), repository)
	}
	go worker(bCtx, &wg, addBinaryTelemetry)

	log.Info("Starting processing telemetry")

	return func() {
		log.Info("Shutting down")

		log.Info("Finishing processing telemetry")
		sCancel()
		bCancel()
		wg.Wait()

		log.Info("Closing connections")
		sConn.Close()
		bConn.Close()
		repository.Close()
	}, nil
}
