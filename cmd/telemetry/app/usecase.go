package app

import (
	"context"
	"fmt"
	"net"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/telemetry/internal/parsers"
	"github.com/telemetry/internal/repo"
)

func connect(address string) (net.Conn, error) {
	c, err := net.Dial("tcp", address)
	if err != nil {
		log.WithField("err", err.Error()).Fatal(fmt.Sprintf("Unable to connect to %s", address))
		return c, err
	}

	return c, nil
}

func addTelemetry(ctx context.Context, conn net.Conn, parser parsers.Parser, repository repo.Repo) error {
	p, err := parser.Parse(conn)
	if err != nil {
		log.WithField("err", err.Error()).Error("Unable to add telemetry")
		return errors.Wrap(err, "Unable to parse telemetry")
	}

	if err := repository.Insert(ctx, p); err != nil {
		log.WithField("err", err.Error()).Error("Unable to add telemetry")
		return errors.Wrap(err, "Unable to write telemetry")
	}

	log.WithFields(log.Fields{
		"timestamp":   p.Timestamp,
		"telemetryID": p.TelemetryID,
		"value":       p.Value,
		"sourceType":  parser.Name(),
	}).Info("Added new telemetry point")

	return nil
}
