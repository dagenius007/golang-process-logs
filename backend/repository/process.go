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

	q := p.DB.NewSelect().Column("id", "user", "pid").ColumnExpr("CAST(cpu_usage AS REAL) AS cpu_usage").ColumnExpr("CAST(memory_usage AS REAL) AS memory_usage").Column("resident_memory_size", "virtual_memory_size", "state", "total_time", "cpu_time", "command", "priority", "created_at", "updated_at").Model(&processes)
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
	count, err := q.Limit(filter.Limit).Offset(filter.Offset).Order("cpu_usage DESC").Order("memory_usage DESC").ScanAndCount(ctx)
	if err != nil {
		return nil, 0, err
	}

	return processes, count, nil
}

func (p ProcessRepository) InsertProcesses(ctx context.Context, processes []*types.Process) error {
	if _, err := p.DB.NewInsert().
		Model(&processes).
		On("CONFLICT (pid) DO UPDATE").
		Set("cpu_usage = EXCLUDED.cpu_usage").
		Set("memory_usage = EXCLUDED.memory_usage").
		Set("resident_memory_size = EXCLUDED.resident_memory_size").
		Set("virtual_memory_size = EXCLUDED.virtual_memory_size").
		Set("state = EXCLUDED.state").
		Set("total_time = EXCLUDED.total_time").
		Set("cpu_time = EXCLUDED.cpu_time").
		Set("command = EXCLUDED.command").
		Set("priority = EXCLUDED.priority").
		Set("updated_at = EXCLUDED.updated_at").
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (p ProcessRepository) GetProcessReport(ctx context.Context) ([]types.ProcessUserReport, error) {
	processesReport := []types.ProcessUserReport{}

	if err := p.DB.NewRaw("SELECT user, ROUND(SUM(cpu_usage),2) AS total_cpu_usage, ROUND(SUM(memory_usage),2) AS total_memory_usage , COUNT(pid) as total_processes FROM ? GROUP BY user ORDER BY COUNT(pid) DESC", bun.Ident("processes")).Scan(ctx, &processesReport); err != nil {
		return processesReport, err
	}

	return processesReport, nil
}

func (p ProcessRepository) GetUsers(ctx context.Context) ([]string, error) {
	users := make([]string, 0)

	// "SELECT user FROM processes GROUP BY user"
	if err := p.DB.NewSelect().Model(((*types.Process)(nil))).Column("user").Group("user").Scan(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (p ProcessRepository) GetCounts(ctx context.Context) (types.DashboardCounts, error) {
	counts := types.DashboardCounts{}

	// "SELECT user FROM processes GROUP BY user"
	if err := p.DB.NewRaw("SELECT COUNT(DISTINCT user) as total_users , COUNT(DISTINCT pid) as total_processes FROM processes").Scan(ctx, &counts); err != nil {
		return counts, err
	}

	return counts, nil
}
