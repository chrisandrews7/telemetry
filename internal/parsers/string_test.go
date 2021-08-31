package parsers_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/telemetry/internal/parsers"
)

func TestStringParser(t *testing.T) {
	input := []byte("[1604614491:1:2.000000]")
	rd := bytes.NewReader(input)

	sp := parsers.NewStringParser()
	point, err := sp.Parse(rd)

	assert.NoError(t, err)
	assert.Equal(t, int64(1604614491), point.Timestamp)
	assert.Equal(t, uint16(1), point.TelemetryID)
	assert.Equal(t, float32(2.000000), point.Value)
}

func TestInvalidStringParser(t *testing.T) {
	input := []byte("[1604614491:CAT:2.000000]")
	rd := bytes.NewReader(input)

	sp := parsers.NewStringParser()
	_, err := sp.Parse(rd)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Unable to parse point")
}
