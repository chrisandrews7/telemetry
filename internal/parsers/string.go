package parsers

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/telemetry/internal/point"
)

type StringParser struct{}

func (s *StringParser) generatePoint(timestamp, telemetryID, value string) (*point.Point, error) {
	parsedTimestamp, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return nil, err
	}

	parsedTelemetryID, err := strconv.ParseUint(telemetryID, 10, 16)
	if err != nil {
		return nil, err
	}

	parsedValue, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return nil, err
	}

	return &point.Point{
		Timestamp:   parsedTimestamp,
		TelemetryID: uint16(parsedTelemetryID),
		Value:       float32(parsedValue),
	}, nil
}

func (s *StringParser) Parse(rd io.Reader) (*point.Point, error) {
	message, err := bufio.NewReader(rd).ReadString(']')
	if err != nil {
		return nil, errors.Wrap(err, "Unable to read point")
	}

	parsedMessage := strings.Split(strings.Trim(message, "[]"), ":")

	point, err := s.generatePoint(parsedMessage[0], parsedMessage[1], parsedMessage[2])
	if err != nil {
		return nil, errors.Wrap(err, "Unable to parse point")
	}

	return point, nil
}

func (s *StringParser) Name() string {
	return "string"
}

func NewStringParser() *StringParser {
	return &StringParser{}
}
