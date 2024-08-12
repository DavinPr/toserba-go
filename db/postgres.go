package db

import (
	"time"

	"github.com/DavinPr/toserba-go/errors"
	"github.com/jmoiron/sqlx"
)

func NewPostgres(cfg PostgresConfig) *sqlx.DB {
	db, err := sqlx.Open("postgres", cfg.GetConnectionString())
	if err != nil {
		panic(errors.Wrap(err, "failed to load the database"))
	}

	if err = db.Ping(); err != nil {
		panic(errors.Wrap(err, "failed to ping to the database"))
	}

	db.SetMaxOpenConns(cfg.GetMaxPoolSize())
	if cfg.GetMaxIdleConnections() != 0 {
		db.SetMaxIdleConns(cfg.GetMaxIdleConnections())
	}

	if maxLifeTime := cfg.GetConnectionMaxLifeTime(); maxLifeTime != 0 {
		db.SetConnMaxLifetime(maxLifeTime)
	}

	// backward compatibility for go under 1.15
	var v interface{} = db
	if d, ok := v.(interface {
		SetConnMaxIdleTime(time.Duration)
	}); ok {
		d.SetConnMaxIdleTime(cfg.GetConnectionMaxIdleTime())
	} else if cfg.GetConnectionMaxIdleTime() != 0 {
		panic(errors.New("SetConnMaxIdleTime feature requires Go version 1.15 or higher.\n" +
			"To resolve this issue, please update your Go version to 1.15 or higher, or set the DB_CONNECTION_MAX_IDLE_TIME configuration to 0."))
	}

	return db
}
