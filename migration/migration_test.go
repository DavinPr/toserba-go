package migration_test

import (
	"fmt"
	"os/exec"
	"testing"

	"github.com/DavinPr/toserba-go/config"
	"github.com/DavinPr/toserba-go/db"
	"github.com/DavinPr/toserba-go/migration"
	"github.com/stretchr/testify/suite"
)

type MigrationTestSuite struct {
	suite.Suite
	migrationPath string
	bashPath      string
}

func (s *MigrationTestSuite) SetupSuite() {
	config.Init("application", "yml")

	s.bashPath = config.GetString("BASH_PATH")
	s.migrationPath = "sample"

	migration.InitPostgres(db.NewPostgresConfig(), s.migrationPath)
}

func TestMigrationTestSuite(t *testing.T) {
	suite.Run(t, new(MigrationTestSuite))
}

func (s *MigrationTestSuite) TearDownSuite() {
}

func (s *MigrationTestSuite) TestCreate() {
	migrationFileName := "create_test_table"

	err := migration.Create(migrationFileName)
	s.NoError(err)

	removeFileCmd := fmt.Sprintf("find %s -type f -name '*%s*' -exec rm {} +", s.migrationPath, migrationFileName)
	err = exec.Command(s.bashPath, "-c", removeFileCmd).Run()
	s.NoError(err, "Failed to remove sample directory")
}

func (s *MigrationTestSuite) TestRunAndRollback() {
	err := migration.Run()
	s.NoError(err)

	err = migration.RollbackLatest()
	s.NoError(err)
}
