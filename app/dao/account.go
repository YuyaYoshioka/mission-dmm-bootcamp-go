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
	// Implementation for repository.Account
	account struct {
		db *sqlx.DB
	}
)

// Create accout repository
func NewAccount(db *sqlx.DB) repository.Account {
	return &account{db: db}
}

func (r *account) CreateAccount(ctx context.Context, username string, passwordHash string, displayName string, note string, avatar string, header string) (*object.Account, error) {
	entity := new(object.Account)
	stmt, err := r.db.Prepare("insert into account (username, password_hash, display_name, avatar, header, note) values(?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	
	result, err := stmt.Exec(username, passwordHash, displayName, note, avatar, header)
	if err != nil {
		return nil, err
	}

	userId, _ := result.LastInsertId()

	r.db.QueryRowxContext(ctx, "select * from account where id = ?", userId).StructScan(entity)
	return entity, nil
}

// FindByUsername : ユーザ名からユーザを取得
func (r *account) FindByUsername(ctx context.Context, username string) (*object.Account, error) {
	entity := new(object.Account)
	err := r.db.QueryRowxContext(ctx, "select * from account where username = ?", username).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("%w", err)
	}

	return entity, nil
}
