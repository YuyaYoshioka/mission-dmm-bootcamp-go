package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Account interface {
	// Fetch account which has specified username
	FindByUsername(ctx context.Context, username string) (*object.Account, error)
	// Create account
	CreateAccount(ctx context.Context, username string, passwordHash string, displayName string, note string, avatar string, header string) (*object.Account, error)
	// Fetch account by ID
	FindByID(ctx context.Context, id int64) (*object.Account, error)
}
