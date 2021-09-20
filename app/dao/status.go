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
	// Implementation for repository.Status
	status struct {
		db *sqlx.DB
	}
)

// Create status repository
func NewStatus(db *sqlx.DB) repository.Status {
	return &status{db: db}
}

func (r *status) CreateStatus(ctx context.Context, content string, accountID int64) (*object.Status, error) {
	entity := new(object.Status)
	stmt, err := r.db.Prepare("insert into status (content, account_id) values(?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	
	result, err := stmt.Exec(content, accountID)
	if err != nil {
		return nil, err
	}

	statusId, _ := result.LastInsertId()

	r.db.QueryRowxContext(ctx, "select * from status where id = ?", statusId).StructScan(entity)
	return entity, nil
}

func (r *status) FindByID(ctx context.Context, ID int64) (*object.Status, error) {
	entity := new(object.Status)
	err := r.db.QueryRowxContext(ctx, "select * from status where id = ?", ID).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("%w", err)
	}

	return entity, nil	
}

func (r *status) DeleteByID(ctx context.Context, ID int64) (error) {
	_, err := r.db.Query("delete from status where id = ?", ID)
	if err != nil {
		fmt.Println("cccc")
		return err
	}

	fmt.Println("eeeeee")
	return nil
}
