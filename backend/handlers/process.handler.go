package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"binalyze-test/configs"
	processHandler "binalyze-test/process"
	. "binalyze-test/types"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

func insertManyProcessQuery(processes []Process) error {
	sqlStr := "INSERT INTO processes (user, pid , cpuUsage , memoryUsage , residentMemorySize,  virtualMemorySize , state , totalTime , cpuTime , command , priority, createdAt , updatedAt) VALUES "

	upsertStatement := " ON CONFLICT (pid) DO UPDATE SET cpuUsage=excluded.cpuUsage , memoryUsage=excluded.memoryUsage , residentMemorySize=excluded.residentMemorySize, virtualMemorySize=excluded.virtualMemorySize , state=excluded.state , totalTime=excluded.totalTime,  cpuTime=excluded.cpuTime,  priority=excluded.priority, updatedAt=excluded.updatedAt ;"
	vals := []interface{}{}

	for _, process := range processes {

		sqlStr += "(?, ? , ? , ? ,?, ? , ? , ? , ? , ? , ? , ? , ?),"
		vals = append(vals, process.User, process.PID, process.CpuUsage, process.MemoryUsage, process.ResidentMemorySize, process.VirtualMemorySize, process.State, process.TotalTime, process.CpuTime, process.Command, process.Priority, time.Now(), time.Now())
	}

	// trim the last ,
	sqlStr = sqlStr[0 : len(sqlStr)-1]

	// concatenate upsert statement
	sqlStr += upsertStatement

	// format all vals at once
	_, err := configs.Db.Exec(sqlStr, vals...)
	if err != nil {
		fmt.Println("err", err)
		return err
	}
	return nil
}

func FetchAndInsertProcess() {
	processes := processHandler.GetProcesses()

	err := insertManyProcessQuery(processes)
	if err != nil {
		// log errror
		fmt.Println("err", err)
	}

	fmt.Println("got here")
}

func selectProcessesQuery(query string) ([]Process, error) {
	processes := []Process{}

	rows, err := configs.Db.Query(query)
	if err != nil {
		return processes, err
	}

	for rows.Next() {
		fmt.Println("row", rows)
		i := Process{}
		err = rows.Scan(&i.ID, &i.User, &i.PID, &i.CpuUsage, &i.MemoryUsage, &i.VirtualMemorySize, &i.ResidentMemorySize, &i.State, &i.TotalTime, &i.CpuTime, &i.Command, &i.Priority, &i.CreatedAt, &i.UpdatedAt)
		if err != nil {
			return processes, err
		}
		processes = append(processes, i)
	}

	// defer rows.Close()

	return processes, nil
}

func buildQuery(params url.Values) string {
	query := "SELECT * FROM processes"
	whereQuery := ""
	// build query based on params
	page, limit := 1, 10

	if params.Has("state") {
		whereQuery = fmt.Sprintf("%s state = %s AND", whereQuery, params.Get("state"))
	}

	if params.Has("user") {
		whereQuery = fmt.Sprintf("%s user = %s AND", whereQuery, params.Get("user"))
	}

	if len(whereQuery) > 0 {
		// find last AND and remove
		lastAndIndex := strings.LastIndex(query, "AND")
		whereQuery = whereQuery[:lastAndIndex]

		whereQuery = "WHERE" + whereQuery + " "
	}

	if val, err := strconv.Atoi(params.Get("page")); err == nil {
		page = val
	}

	if val, err := strconv.Atoi(params.Get("limit")); err == nil {
		limit = val
	}

	offset := (page - 1) * limit

	query = fmt.Sprintf("%s %sLIMIT %d OFFSET %d", query, whereQuery, limit, offset)

	return query
}

func GetProcess(c echo.Context) error {
	query := buildQuery(c.QueryParams())

	processes, err := selectProcessesQuery(query)
	if err != nil {
		fmt.Println("err", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"data":    processes,
			"success": false,
			"message": "Operation not successful",
		})
	}

	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    processes,
			"success": true,
			"message": "Operation successful",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    processes,
		"success": true,
		"message": "Operation successful",
	})
}

func GetProcessRealTime(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			// Write

			query := buildQuery(c.QueryParams())

			data, err := selectProcessesQuery(query)
			if err != nil {
				err := websocket.Message.Send(ws, "[]")
				if err != nil {
					c.Logger().Error(err)
				}
			}

			var _json []byte

			_json, err = json.Marshal(data)

			if err != nil {
				c.Logger().Error(err)
			}

			msg := string(_json)

			err = websocket.Message.Send(ws, msg)
			if err != nil {
				fmt.Println("err", err)
				c.Logger().Error(err)
			}
		}
	}).ServeHTTP(c.Response(), c.Request())

	return nil
}
