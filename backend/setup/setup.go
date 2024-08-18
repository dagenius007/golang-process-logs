package setup

import (
	"database/sql"

	"binalyze-test/repository"
	"binalyze-test/services"

	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
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
		return nil, err
	}

	processRepo := repository.NewRepository(logger, service.DB)

	service.ProcessService = services.NewProcessService(logger, processRepo)

	return service, nil
}

func connectDb() (*bun.DB, error) {
	sqldb, err := sql.Open(sqliteshim.ShimName, "../db/processes.db")
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return bun.NewDB(sqldb, sqlitedialect.New()), nil
}
