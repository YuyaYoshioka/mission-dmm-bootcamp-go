package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Status interface {
	// Create status
	CreateStatus(ctx context.Context, content string, accountID int64) (*object.Status, error)
	// Fetch status by ID
	FindByID(ctx context.Context, ID int64) (*object.Status, error)
}
