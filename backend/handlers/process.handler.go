package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"syscall"
	"time"

	"binalyze-test/utils"

	"binalyze-test/setup"
	"binalyze-test/types"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

type ProcessHandler struct {
	services *setup.ServiceDependencies
}

func UseProcessRoutes(routes *echo.Group, services *setup.ServiceDependencies) {
	p := ProcessHandler{services: services}
	routes.GET("/processes", p.getProcesses)
	routes.GET("/users", p.getProcessUsers)
	routes.GET("/counts", p.getDashboardCounts)
	routes.GET("/reports", p.getProcessReport)
	// routes.GET("/ws", p.getProcessRealTime)
}

func (p ProcessHandler) getProcesses(c echo.Context) error {
	ctx := c.Request().Context()

	page, limit := 1, 10

	if val, err := strconv.Atoi(c.QueryParam("page")); err == nil {
		page = val
	}

	if val, err := strconv.Atoi(c.QueryParam("limit")); err == nil {
		limit = val
	}

	offset := (page - 1) * limit

	processList, err := p.services.ProcessService.GetProcesses(ctx, types.ProcessFilter{
		State:  c.QueryParam("state"),
		User:   c.QueryParam("user"),
		Search: c.QueryParam("search"),
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return utils.SendResponse(c, http.StatusInternalServerError, utils.ErrorResponse("Operation not successful"))
	}

	processList.Limit = limit
	processList.Page = page

	return utils.SendResponse(c, http.StatusOK, utils.SuccessResponse(processList))
}

func (p ProcessHandler) getProcessUsers(c echo.Context) error {
	ctx := c.Request().Context()
	data, err := p.services.ProcessService.GetProcessUsers(ctx)
	if err != nil {
		return utils.SendResponse(c, http.StatusInternalServerError, utils.ErrorResponse("Operation not successful"))
	}

	return utils.SendResponse(c, http.StatusOK, utils.SuccessResponse(data))
}

func (p ProcessHandler) getDashboardCounts(c echo.Context) error {
	ctx := c.Request().Context()

	data, err := p.services.ProcessService.GetDashboardCounts(ctx)
	if err != nil {
		return utils.SendResponse(c, http.StatusInternalServerError, utils.ErrorResponse("Operation not successful"))
	}

	return utils.SendResponse(c, http.StatusOK, utils.SuccessResponse(data))
}

func (p ProcessHandler) getProcessReport(c echo.Context) error {
	ctx := c.Request().Context()

	data, err := p.services.ProcessService.GetProcessReport(ctx)
	if err != nil {
		return utils.SendResponse(c, http.StatusInternalServerError, utils.ErrorResponse("Operation not successful"))
	}

	return utils.SendResponse(c, http.StatusOK, utils.SuccessResponse(data))
}

func (p ProcessHandler) getProcessRealTime(c echo.Context) error {
	ctx := c.Request().Context()
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			// Write

			processList, err := p.services.ProcessService.GetProcesses(ctx, types.ProcessFilter{})
			if err != nil {
			}

			report, err := p.services.ProcessService.GetProcessReport(ctx)
			if err != nil {
			}

			data := types.RealTimeData{
				Processes: *processList,
				Report:    report,
			}

			response, err := json.Marshal(data)
			if err != nil {
				c.Logger().Error(err)
			}

			msg := string(response)

			time.Sleep(1 * time.Minute)

			err = websocket.Message.Send(ws, msg)
			if err != nil {
				if errors.Is(err, syscall.EPIPE) {
					break
				}
			}
		}
	}).ServeHTTP(c.Response(), c.Request())

	return nil
}
