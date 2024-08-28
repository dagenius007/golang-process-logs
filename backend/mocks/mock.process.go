package mock

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"process-logs/types"
)

type (
	Mock struct {
		db []types.Process
	}
)

func (m *Mock) InsertProccess(processes []types.Process) {
	m.db = processes
}

func (m *Mock) GetProcess(c echo.Context) error {
	processes := m.db

	return c.JSON(http.StatusOK, processes)
}

func (m *Mock) GetProcessCount(c echo.Context) error {
	data := map[string]int{
		"processCount": 0,
		"usersCount":   0,
	}

	if len(m.db) == 0 {
		return c.JSON(http.StatusOK, data)
	}

	data["processCount"] = 3
	data["usersCount"] = 2

	return c.JSON(http.StatusOK, data)
}
