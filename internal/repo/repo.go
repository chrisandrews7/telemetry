package repo

import (
	"context"

	"github.com/telemetry/internal/point"
)

type Repo interface {
	Insert(context.Context, *point.Point) error
	Close()
}
