package db_test

import (
	"os"
	"testing"
	"time"

	"github.com/DavinPr/toserba-go/config"
	"github.com/DavinPr/toserba-go/db"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/stretchr/testify/suite"
)

type PostgresTestSuite struct {
	suite.Suite
	cfg db.PostgresConfig
}

func (s *PostgresTestSuite) SetupTest() {
	config.Init("application", "yml")

	s.cfg = db.NewPostgresConfig()
}

func (s *PostgresTestSuite) TearDownTest() {
}

func TestNewPostgresTestSuite(t *testing.T) {
	suite.Run(t, new(PostgresTestSuite))
}

func (s *PostgresTestSuite) TestNewPostgresConfig() {
	s.Equal("toserba", s.cfg.GetName())
	s.Equal("postgres://postgres:postgres@localhost:5432/toserba?sslmode=disable", s.cfg.GetConnectionURL())
	s.Equal("dbname=toserba user=postgres password='postgres' host=localhost port=5432 sslmode=disable", s.cfg.GetConnectionString())
	s.Equal(time.Second*30, s.cfg.GetConnectionMaxIdleTime())
	s.Equal(3, s.cfg.GetMaxPoolSize())
	s.Equal(10, s.cfg.GetMaxIdleConnections())
	s.Equal(time.Minute*5, s.cfg.GetConnectionMaxLifeTime())

	os.Setenv("DB_CONNECTION_MAX_LIFE_TIME_JITTER", "2m")
	defer os.Unsetenv("DB_CONNECTION_MAX_LIFE_TIME_JITTER")

	cfg := db.NewPostgresConfig()
	maxLifeTime := cfg.GetConnectionMaxLifeTime()
	s.Truef(maxLifeTime >= time.Minute*5, "%v should be greater than equal to 5 mins", maxLifeTime)
	s.Truef(maxLifeTime < time.Minute*7, "%v should be less than 7 mins", maxLifeTime)
}

func (s *PostgresTestSuite) TestNewPostgres() {
	pg := db.NewPostgres(s.cfg)
	s.NotNil(pg)
	s.Nil(pg.Close())
}
