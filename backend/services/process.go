package services

import (
	"context"
	"log"

	"binalyze-test/process"
	"binalyze-test/types"

	"github.com/sirupsen/logrus"
)

type ProcessService struct {
	logger *logrus.Logger
	repo   types.ProcessRepository
}

func NewProcessService(logger *logrus.Logger, repo types.ProcessRepository) *ProcessService {
	return &ProcessService{
		logger: logger,
		repo:   repo,
	}
}

func (p ProcessService) FetchAndInsertProcess() {
	processes := process.GetProcesses()

	err := insertManyProcessQuery(processes)
	if err != nil {
		// log errror
		log.Println("err", err)
	}

	log.Println("Running processes fetched and inserted into db")
}

func (c ProcessService) GetProcesses(ctx context.Context, filter types.ProcessFilter) (*types.ProcessList, error) {
	useLogger := c.logger.WithContext(ctx).WithField("function", "GetProcesses")
	processes, total, err := c.repo.GetProcesses(ctx, filter)
	if err != nil {
		useLogger.WithError(err).Error("Error fetching processes")
		return nil, err
	}
	return &types.ProcessList{
		Processes: processes,
		Totoal:    total,
	}, nil
}
