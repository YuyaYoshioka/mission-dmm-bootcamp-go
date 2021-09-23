package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	timelinesPublic struct {
		db *sqlx.DB
	}
)

func NewTimelinesPublic(db *sqlx.DB) repository.TimelinesPublic {
	return &timelinesPublic{db: db}
}

func (r *timelinesPublic) FetchAll(ctx context.Context) ([]object.Status, error) {
	var entity []object.Status

	sqlStatement := `
	SELECT
		s.id as "id",
		s.content as "content",
		s.create_at as "create_at",
		a.username as "account.username",
		a.display_name as "account.display_name",
		a.create_at as "account.create_at",
		a.note as "account.note",
		a.avatar as "account.avatar",
		a.header as "account.header"
	FROM
		status as s
	JOIN
		account as a
	ON
		s.account_id = a.id
	`

	err := r.db.Select(&entity, sqlStatement)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("%w", err)
	}

	return entity, nil
}
