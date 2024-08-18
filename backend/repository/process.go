package repository

import (
	"context"

	"binalyze-test/types"

	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
)

type ProcessRepository struct {
	logger *logrus.Logger
	DB     *bun.DB
}

func NewRepository(logger *logrus.Logger, DB *bun.DB) *ProcessRepository {
	return &ProcessRepository{
		logger: logger,
		DB:     DB,
	}
}

func (p ProcessRepository) GetProcesses(ctx context.Context, filter types.ProcessFilter) ([]types.Process, int, error) {
	processes := []types.Process{}

	q := p.DB.NewSelect().Model(&processes)

	if filter.State != "" {
		q = q.Where("state = ?", filter.State)
	}

	if filter.User != "" {
		q = q.Where("user = ?", filter.State)
	}

	if filter.Search != "" {
		// name LIKE
		q = q.WhereGroup("AND", func(sq *bun.SelectQuery) *bun.SelectQuery {
			return sq.WhereOr("user LIKE", "%"+filter.Search+"%").WhereOr("command LIKE", "%"+filter.Search+"%")
		})
	}

	count, err := q.ScanAndCount(ctx)
	if err != nil {
		return nil, 0, err
	}

	return processes, count, nil
}

func (p ProcessRepository) InsertProcesses(ctx context.Context, processes types.Process) ([]types.Process, int, error) {
	processes := []types.Process{}

	q := p.DB.NewSelect().Model(&processes)

	if filter.State != "" {
		q = q.Where("state = ?", filter.State)
	}

	if filter.User != "" {
		q = q.Where("user = ?", filter.State)
	}

	if filter.Search != "" {
		// name LIKE
		q = q.WhereGroup("AND", func(sq *bun.SelectQuery) *bun.SelectQuery {
			return sq.WhereOr("user LIKE", "%"+filter.Search+"%").WhereOr("command LIKE", "%"+filter.Search+"%")
		})
	}

	count, err := q.ScanAndCount(ctx)
	if err != nil {
		return nil, 0, err
	}

	return processes, count, nil
}
