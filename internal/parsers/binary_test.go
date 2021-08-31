package parsers_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/telemetry/internal/parsers"
)

func TestBinaryParser(t *testing.T) {
	input := []byte{0, 1, 2, 3, 206, 169, 44, 97, 0, 0, 0, 0, 2, 0, 218, 249, 87, 65}
	rd := bytes.NewReader(input)

	bp := parsers.NewBinaryParser()
	point, err := bp.Parse(rd)

	assert.NoError(t, err)
	assert.Equal(t, int64(1630317006), point.Timestamp)
	assert.Equal(t, uint16(2), point.TelemetryID)
	assert.Equal(t, float32(13.498499), point.Value)
}

func TestInvalidBinaryParser(t *testing.T) {
	input := []byte{0, 1, 2, 3}
	rd := bytes.NewReader(input)

	sp := parsers.NewBinaryParser()
	_, err := sp.Parse(rd)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Unable to parse point")
}
