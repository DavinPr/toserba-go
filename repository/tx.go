package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type TxRunner interface {
	RunInTx(fn func(*sqlx.Tx) error) error
	RunInTxContext(ctx context.Context, fn func(*sqlx.Tx) error) error
}

type txRunner struct {
	db *sqlx.DB
}

func NewTxRunner(db *sqlx.DB) *txRunner {
	return &txRunner{db: db}
}

// Based on https://github.com/go-pg/pg/blob/master/tx.go#L51
func (r *txRunner) RunInTx(fn func(*sqlx.Tx) error) (err error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback()
			err = errors.New(fmt.Sprint(r))
		}

	}()
	if err = fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (r *txRunner) RunInTxContext(ctx context.Context, fn func(*sqlx.Tx) error) (err error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback()
			err = errors.New(fmt.Sprint(r))
		}
	}()

	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}
