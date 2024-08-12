package repository

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

var (
	ErrZeroRowsAffected = errors.New("no rows affected")
)

type BaseRepository struct {
	db        *sqlx.DB
	appName   string
	tableName string
	txRunner  TxRunner
}

func NewBaseRepository(db *sqlx.DB, appName string, tableName string) BaseRepository {
	return BaseRepository{
		db:        db,
		appName:   appName,
		tableName: tableName,
		txRunner:  NewTxRunner(db),
	}
}

func (r *BaseRepository) DBGet(dest interface{}, query string, args ...interface{}) error {
	return r.db.Get(dest, query, args...)
}

func (r *BaseRepository) DBGetInTx(tx *sqlx.Tx, dest interface{}, query string, args ...interface{}) error {
	if tx == nil {
		return r.DBGet(dest, query, args...)
	}
	return tx.Get(dest, query, args...)
}

func (r *BaseRepository) DBSelect(dest interface{}, query string, args ...interface{}) error {
	return r.db.Select(dest, query, args...)
}

func (r *BaseRepository) DBSelectInTx(tx *sqlx.Tx, dest interface{}, query string, args ...interface{}) error {
	if tx == nil {
		return r.DBSelect(dest, query, args...)
	}
	return tx.Select(dest, query, args...)
}

func (r *BaseRepository) DBExec(query string, args ...interface{}) error {
	return r.txRunner.RunInTx(func(tx *sqlx.Tx) error {
		return r.DBExecInTx(tx, query, args...)
	})
}

func (r *BaseRepository) DBExecInTx(tx *sqlx.Tx, query string, args ...interface{}) error {
	if tx == nil {
		return r.DBExec(query, args...)
	}
	res, err := tx.Exec(query, args...)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if rowsAffected == int64(0) {
		return ErrZeroRowsAffected
	}
	return err
}

func (r *BaseRepository) DBSoftExec(query string, args ...interface{}) error {
	return r.txRunner.RunInTx(func(tx *sqlx.Tx) error {
		return r.DBSoftExecInTx(tx, query, args...)
	})
}

func (r *BaseRepository) DBSoftExecInTx(tx *sqlx.Tx, query string, args ...interface{}) error {
	if tx == nil {
		return r.DBSoftExec(query, args...)
	}
	res, err := tx.Exec(query, args...)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	return err
}

func (r *BaseRepository) DBExecReturning(dest interface{}, query string, args ...interface{}) error {
	return r.txRunner.RunInTx(func(tx *sqlx.Tx) error {
		return r.DBExecReturningInTx(tx, dest, query, args...)
	})
}

func (r *BaseRepository) DBExecReturningInTx(tx *sqlx.Tx, dest interface{}, query string, args ...interface{}) error {
	if tx == nil {
		return r.DBExecReturning(dest, query, args...)
	}
	row := tx.QueryRowx(query, args...)
	return row.StructScan(dest)
}
