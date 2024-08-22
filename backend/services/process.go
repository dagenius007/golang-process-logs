package services

import (
	"context"

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

func (p ProcessService) FetchAndInsertProcess(ctx context.Context) {
	useLogger := p.logger.WithContext(ctx).WithField("function", "FetchAndInsertProcess")

	processes := process.GetProcesses()

	err := p.repo.InsertProcesses(ctx, processes)
	if err != nil {
		useLogger.Error(err)
	}

	useLogger.Info("Running processes fetched and inserted into db")
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

func (c ProcessService) GetProcessReport(ctx context.Context) ([]types.ProcessUserReport, error) {
	useLogger := c.logger.WithContext(ctx).WithField("function", "GetProcessReport")
	report, err := c.repo.GetProcessReport(ctx)
	if err != nil {
		useLogger.WithError(err).Error("Error fetching processes report")
		return nil, err
	}

	return report, nil
}

func (c ProcessService) GetProcessUsers(ctx context.Context) ([]string, error) {
	useLogger := c.logger.WithContext(ctx).WithField("function", "GetProcessUsers")
	users, err := c.repo.GetUsers(ctx)
	if err != nil {
		useLogger.WithError(err).Error("Error fetching processes users")
		return nil, err
	}

	return users, nil
}

func (c ProcessService) GetDashboardCounts(ctx context.Context) (types.DashboardCounts, error) {
	useLogger := c.logger.WithContext(ctx).WithField("function", "GetDashboardCounts")
	counts, err := c.repo.GetCounts(ctx)
	if err != nil {
		useLogger.WithError(err).Error("Error fetching user and processes count")
		return counts, err
	}

	return counts, nil
}
