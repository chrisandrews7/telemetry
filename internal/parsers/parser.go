package parsers

import (
	"io"

	"github.com/telemetry/internal/point"
)

type Parser interface {
	Parse(io.Reader) (*point.Point, error)
	Name() string
}
