package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type TimelinesPublic interface {
	// Fetch all status
	FetchAll(ctx context.Context) ([]object.Status, error)
}
