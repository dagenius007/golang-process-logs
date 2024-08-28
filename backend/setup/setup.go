package setup

import (
	"context"
	"database/sql"
	"os"

	"binalyze-test/repository"
	"binalyze-test/services"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/extra/bundebug"

	_ "github.com/mattn/go-sqlite3"
)

const (
	Path     = "../db/"
	FullPath = "../db/processes.db"
)

type ServiceDependencies struct {
	Logger         *logrus.Logger
	DB             *bun.DB
	ProcessService *services.ProcessService
}

func ConfigureServiceDependencies(logger *logrus.Logger) (*ServiceDependencies, error) {
	service := &ServiceDependencies{
		Logger: logger,
	}

	var err error

	if service.DB, err = connectDb(); err != nil {
		logger.WithError(err).Info("Error connecting to DB/ migration")
		return nil, err
	}

	// service.DB.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	service.DB.AddQueryHook(bundebug.NewQueryHook())

	processRepo := repository.NewRepository(logger, service.DB)

	service.ProcessService = services.NewProcessService(logger, processRepo)

	// Prepopulate DB on load

	ctx := context.Background()

	service.ProcessService.FetchAndInsertProcess(ctx)
	return service, nil
}

func connectDb() (*bun.DB, error) {
	err := createDBFile()
	if err != nil {
		return nil, err
	}
	sqldb, err := sql.Open("sqlite3", "../db/processes.db")
	if err != nil {
		return nil, err
	}

	err = runMigrations(sqldb)
	if err != nil {
		return nil, err
	}

	return bun.NewDB(sqldb, sqlitedialect.New()), nil
}

func createDBFile() error {
	// Create new sql file

	_, err := os.Stat(Path)

	if err == nil {
		err := os.RemoveAll(Path)
		if err != nil {
			logrus.Error(err)
			return err
		}
	}

	err = os.Mkdir(Path, 0o777)
	if err != nil {
		return err
	}
	file, err := os.Create(FullPath)
	if err != nil {
		return err
	}

	defer file.Close()

	return nil
}

func runMigrations(db *sql.DB) error {
	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return err
	}

	fSrc, err := (&file.File{}).Open("./migrations")
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance("file", fSrc, "sqlite3", instance)
	if err != nil {
		return err
	}

	// modify for Down
	if err := m.Up(); err != nil {
		return err
	}

	return nil
}
