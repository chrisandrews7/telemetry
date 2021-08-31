package repo_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/stretchr/testify/assert"
	"github.com/telemetry/internal/point"
	"github.com/telemetry/internal/repo"
	"github.com/telemetry/internal/repo/mocks"
)

func TestRepoInsert(t *testing.T) {
	ctx := context.Background()
	mockClient := &mocks.Client{}
	mockAPI := &mocks.WriteAPIBlocking{}
	mockPoint := point.Point{
		Timestamp:   1630317006,
		TelemetryID: 2,
		Value:       12.04319953918457,
	}
	expectedRecord := influxdb2.NewPoint("telemetry",
		map[string]string{"telemetryID": fmt.Sprint(2)},
		map[string]interface{}{"value": 12.04319953918457},
		time.Unix(1630317006, 0))

	mockClient.On("WriteAPIBlocking", "", "testDB").Return(mockAPI, nil)
	mockAPI.On("WritePoint", ctx, expectedRecord).Return(nil)

	r := repo.NewInfluxRepoWithClient(mockClient, "testDB")
	err := r.Insert(ctx, &mockPoint)

	assert.NoError(t, err)
	mockClient.AssertExpectations(t)
	mockAPI.AssertExpectations(t)
}

func TestRepoInsertError(t *testing.T) {
	ctx := context.Background()
	mockClient := &mocks.Client{}
	mockAPI := &mocks.WriteAPIBlocking{}
	mockPoint := point.Point{
		Timestamp:   1630317006,
		TelemetryID: 2,
		Value:       12.04319953918457,
	}
	expectedRecord := influxdb2.NewPoint("telemetry",
		map[string]string{"telemetryID": fmt.Sprint(2)},
		map[string]interface{}{"value": 12.04319953918457},
		time.Unix(1630317006, 0))

	mockClient.On("WriteAPIBlocking", "", "testDB").Return(mockAPI, nil)
	mockAPI.On("WritePoint", ctx, expectedRecord).Return(errors.New("Problem"))

	r := repo.NewInfluxRepoWithClient(mockClient, "testDB")
	err := r.Insert(ctx, &mockPoint)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Unable to insert point")
}
