package db

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/DavinPr/toserba-go/config"
)

type PostgresConfig struct {
	host                  string
	port                  int
	username              string
	password              string
	name                  string
	maxPoolSize           int
	maxIdleConnections    int
	connMaxIdleTime       time.Duration
	connMaxLifeTime       time.Duration
	connMaxLifeTimeJitter time.Duration
}

func NewPostgresConfig() PostgresConfig {
	return PostgresConfig{
		host:               config.MustGetString("DB_HOST"),
		port:               config.MustGetInt("DB_PORT"),
		name:               config.MustGetString("DB_NAME"),
		username:           config.MustGetString("DB_USER"),
		password:           config.MustGetString("DB_PASSWORD"),
		maxPoolSize:        config.MustGetInt("DB_POOL"),
		maxIdleConnections: config.GetInt("DB_MAX_IDLE_CONNECTIONS"),
		// This config will only be used if go version >= 1.15
		connMaxIdleTime:       config.GetDuration("DB_CONNECTION_MAX_IDLE_TIME"),
		connMaxLifeTime:       config.GetDuration("DB_CONNECTION_MAX_LIFE_TIME"),
		connMaxLifeTimeJitter: config.GetDuration("DB_CONNECTION_MAX_LIFE_TIME_JITTER"),
	}
}

func (c PostgresConfig) GetMaxIdleConnections() int {
	return c.maxIdleConnections
}

func (c PostgresConfig) GetMaxPoolSize() int {
	return c.maxPoolSize
}

func (c PostgresConfig) GetConnectionString() string {
	return fmt.Sprintf("dbname=%s user=%s password='%s' host=%s port=%d sslmode=disable", c.name, c.username, c.password, c.host, c.port)
}

func (c PostgresConfig) GetConnectionMaxIdleTime() time.Duration {
	return c.connMaxIdleTime
}

func (c PostgresConfig) GetConnectionMaxLifeTime() time.Duration {
	var jitter time.Duration
	if c.connMaxLifeTimeJitter > 0 {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		jitter = time.Duration(r.Int63n(int64(c.connMaxLifeTimeJitter)))
	}

	return c.connMaxLifeTime + jitter
}

func (c PostgresConfig) GetConnectionURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", c.username, c.password, c.host, c.port, c.name)
}

func (c PostgresConfig) GetName() string {
	return c.name
}
