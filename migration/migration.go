package migration

import (
	"fmt"
	"os"
	"time"

	"github.com/DavinPr/toserba-go/db"
	"github.com/DavinPr/toserba-go/errors"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var appMigrationFilesPath string
var appMigrate *migrate.Migrate

func InitPostgres(dbConfig db.PostgresConfig, migrationFilesPath string) {
	appMigrationFilesPath = migrationFilesPath

	var err error
	appMigrate, err = migrate.New("file://"+migrationFilesPath, dbConfig.GetConnectionURL())
	if err != nil {
		panic(errors.Wrap(err, "failed to init migration"))
	}
}

func Create(filename string) error {
	if len(filename) == 0 {
		return errors.New("Migration filename is not provided")
	}

	timeStamp := time.Now().Unix()
	upMigrationFilePath := fmt.Sprintf("%s/%d_%s.up.sql", appMigrationFilesPath, timeStamp, filename)
	downMigrationFilePath := fmt.Sprintf("%s/%d_%s.down.sql", appMigrationFilesPath, timeStamp, filename)

	if err := createFile(upMigrationFilePath); err != nil {
		return err
	}
	fmt.Printf("Created %s\n", upMigrationFilePath)

	if err := createFile(downMigrationFilePath); err != nil {
		os.Remove(upMigrationFilePath)
		return err
	}

	fmt.Printf("Created %s\n", downMigrationFilePath)

	return nil
}

func Run() error {
	err := appMigrate.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	fmt.Println("Migrations successful")
	return nil
}

func RollbackLatest() error {
	err := appMigrate.Steps(-1)
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	fmt.Println("Rollback successful")
	return nil
}

func createFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	return nil
}
