package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mock "binalyze-test/mocks"
	"binalyze-test/types"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	processes = []types.Process{}

	countData = map[string]int{"processCount": 3, "usersCount": 2}

	m mock.Mock
)

func init() {
	process1 := types.Process{
		User:               "joshuaoluikpe",
		PID:                1,
		CpuUsage:           0.50,
		MemoryUsage:        0.60,
		ResidentMemorySize: 53434,
		VirtualMemorySize:  56,
		State:              "running",
		TotalTime:          "20.6",
		CpuTime:            "670.9",
		Command:            "go",
		Priority:           "medium",
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	process2 := types.Process{
		User:               "root",
		PID:                2,
		CpuUsage:           3.5,
		MemoryUsage:        0.9,
		ResidentMemorySize: 58454,
		VirtualMemorySize:  56,
		State:              "sleeping",
		TotalTime:          "20.9",
		CpuTime:            "67.9",
		Command:            "/System/Library",
		Priority:           "medium",
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	process3 := types.Process{
		User:               "joshuaoluikpe",
		PID:                3,
		CpuUsage:           0.7,
		MemoryUsage:        0.8,
		ResidentMemorySize: 67544,
		VirtualMemorySize:  5,
		State:              "sleeping",
		TotalTime:          "309.9",
		CpuTime:            "67.9",
		Command:            "Music",
		Priority:           "high",
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	processes = append(processes, process1, process2, process3)

	m.InsertProccess(processes)
}

func TestGetProcesses(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/data", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, m.GetProcess(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetProcessCount(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/count", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, m.GetProcessCount(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response map[string]int
		json.Unmarshal(rec.Body.Bytes(), &response)
		processCount := response["processCount"]

		assert.Equal(t, countData["processCount"], processCount)

		userCount := response["userCount"]

		assert.Equal(t, countData["userCount"], userCount)
	}
}
