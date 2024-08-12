package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type BaseContextRepository struct {
	db        *sqlx.DB
	appName   string
	tableName string
	txRunner  TxRunner
}

func NewBaseContextRepository(db *sqlx.DB, appName string, tableName string) BaseContextRepository {
	return BaseContextRepository{
		db:        db,
		appName:   appName,
		tableName: tableName,
		txRunner:  NewTxRunner(db),
	}
}

func (r *BaseContextRepository) DBGet(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return r.db.GetContext(ctx, dest, query, args...)
}

func (r *BaseContextRepository) DBGetInTx(ctx context.Context, tx *sqlx.Tx, dest interface{}, query string, args ...interface{}) error {
	if tx == nil {
		return r.DBGet(ctx, dest, query, args...)
	}
	return tx.GetContext(ctx, dest, query, args...)
}

func (r *BaseContextRepository) DBSelect(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return r.db.SelectContext(ctx, dest, query, args...)
}

func (r *BaseContextRepository) DBSelectInTx(ctx context.Context, tx *sqlx.Tx, dest interface{}, query string, args ...interface{}) error {
	if tx == nil {
		return r.DBSelect(ctx, dest, query, args...)
	}
	return tx.SelectContext(ctx, dest, query, args...)
}

func (r *BaseContextRepository) DBExec(ctx context.Context, query string, args ...interface{}) error {
	return r.txRunner.RunInTxContext(ctx, func(tx *sqlx.Tx) error {
		return r.DBExecInTx(ctx, tx, query, args...)
	})
}

func (r *BaseContextRepository) DBExecInTx(ctx context.Context, tx *sqlx.Tx, query string, args ...interface{}) error {
	if tx == nil {
		return r.DBExec(ctx, query, args...)
	}
	res, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if rowsAffected == int64(0) {
		return ErrZeroRowsAffected
	}
	return err
}

func (r *BaseContextRepository) DBSoftExec(ctx context.Context, query string, args ...interface{}) error {
	return r.txRunner.RunInTxContext(ctx, func(tx *sqlx.Tx) error {
		return r.DBSoftExecInTx(ctx, tx, query, args...)
	})
}

func (r *BaseContextRepository) DBSoftExecInTx(ctx context.Context, tx *sqlx.Tx, query string, args ...interface{}) error {
	if tx == nil {
		return r.DBSoftExec(ctx, query, args...)
	}
	res, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	return err
}

func (r *BaseContextRepository) DBNamedExec(ctx context.Context, query string, arg interface{}) error {
	return r.txRunner.RunInTxContext(ctx, func(tx *sqlx.Tx) error {
		return r.DBNamedExecInTx(ctx, tx, query, arg)
	})
}

func (r *BaseContextRepository) DBNamedExecInTx(ctx context.Context, tx *sqlx.Tx, query string, arg interface{}) error {
	if tx == nil {
		return r.DBNamedExec(ctx, query, arg)
	}

	res, err := tx.NamedExecContext(ctx, query, arg)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if rowsAffected == int64(0) {
		return ErrZeroRowsAffected
	}

	return err
}

func (r *BaseContextRepository) DBExecReturning(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return r.txRunner.RunInTxContext(ctx, func(tx *sqlx.Tx) error {
		return r.DBExecReturningInTx(ctx, tx, dest, query, args...)
	})
}

func (r *BaseContextRepository) DBExecReturningInTx(ctx context.Context, tx *sqlx.Tx, dest interface{}, query string, args ...interface{}) error {
	if tx == nil {
		return r.DBExecReturning(ctx, dest, query, args...)
	}
	row := tx.QueryRowxContext(ctx, query, args...)
	return row.StructScan(dest)
}

func (r *BaseContextRepository) DBNamedExecReturning(ctx context.Context, dest interface{}, query string, arg interface{}) error {
	return r.txRunner.RunInTxContext(ctx, func(tx *sqlx.Tx) error {
		return r.DBNamedExecReturningInTx(ctx, tx, dest, query, arg)
	})
}

func (r *BaseContextRepository) DBNamedExecReturningInTx(ctx context.Context, tx *sqlx.Tx, dest interface{}, query string, arg interface{}) error {
	if tx == nil {
		return r.DBNamedExecReturning(ctx, dest, query, arg)
	}

	query, args, err := r.db.BindNamed(query, arg)
	if err != nil {
		return err
	}

	row := tx.QueryRowxContext(ctx, query, args...)
	return row.StructScan(dest)
}

func (r *BaseContextRepository) BindNamed(query string, arg interface{}) (string, []interface{}, error) {
	return r.db.BindNamed(query, arg)
}
