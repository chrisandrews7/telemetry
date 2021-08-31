package parsers

import (
	"encoding/binary"
	"io"

	"github.com/pkg/errors"
	"github.com/telemetry/internal/point"
)

type BinaryParser struct{}

func (b *BinaryParser) Parse(rd io.Reader) (*point.Point, error) {
	var p struct {
		Header [4]byte
		point.Point
	}

	err := binary.Read(rd, binary.LittleEndian, &p)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to parse point")
	}

	return &point.Point{
		Timestamp:   p.Timestamp,
		TelemetryID: p.TelemetryID,
		Value:       p.Value,
	}, nil
}

func (b *BinaryParser) Name() string {
	return "binary"
}

func NewBinaryParser() *BinaryParser {
	return &BinaryParser{}
}
