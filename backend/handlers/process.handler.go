package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"binalyze-test/configs"
	"binalyze-test/processes"
	. "binalyze-test/types"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

func FetchAndInsertProcess() {
	// processLists, err := ps.Processes()
	processLists, err := processes.GetProcesses()
	if err != nil {
		log.Println("ps.Processes() Failed, are you using windows?")
		return
	}

	fmt.Println("processlist", processLists)

	// sqlStr := "INSERT INTO processes (user, pid , cpuUsage , memoryPercentageUsage , virtualMemorySize , residentMemorySize , tty , state , started , totalTime , command , createdAt , updatedAt) VALUES "

	// upsertStatement := " ON CONFLICT (pid) DO UPDATE SET cpuUsage=excluded.cpuUsage , memoryPercentageUsage=excluded.memoryPercentageUsage , virtualMemorySize=excluded.virtualMemorySize , residentMemorySize=excluded.residentMemorySize, state=excluded.state , totalTime=excluded.totalTime, updatedAt=excluded.updatedAt ;"
	// vals := []interface{}{}

	// // map ages
	// for _, processList := range processLists {

	// 	// fmt.Println("Process : ", processList.PID, processList.TotalTime)

	// 	sqlStr += "(?, ? , ? , ? ,?, ? , ? , ? , ? , ? , ? , ? , ?),"
	// 	vals = append(vals, processList.User, processList.PID, processList.CpuUsage, processList.MemoryPercentageUsage, processList.VirtualMemorySize, processList.ResidentMemorySize, processList.Tty, processList.State, processList.TotalTime, processList.Command, time.Now(), time.Now())
	// }

	// // trim the last ,
	// sqlStr = sqlStr[0 : len(sqlStr)-1]

	// sqlStr = sqlStr + upsertStatement
	// // // prepare the statement
	// // // stmt, _ := configs.Db.Prepare(sqlStr)

	// // // format all vals at once
	// // _, err = configs.Db.Exec(sqlStr, vals...)
	// if err != nil {
	// 	fmt.Println("err:", err)
	// }

	fmt.Println("Insertion sucessful")
}

func fetchDbProcess() ([]Process, error) {
	rows, err := configs.Db.Query("SELECT * FROM processes")
	data := []Process{}

	if err != nil {
		return data, err
	}

	for rows.Next() {
		i := Process{}
		err = rows.Scan(&i.ID, &i.User, &i.PID, &i.CpuUsage, &i.MemoryPercentageUsage, &i.VirtualMemorySize, &i.ResidentMemorySize, &i.Tty, &i.State, &i.Application, &i.TotalTime, &i.Command, &i.CreatedAt, &i.UpdatedAt)
		if err != nil {
			return data, err
		}
		data = append(data, i)
	}

	defer rows.Close()

	return data, nil
}

func GetProcess(c echo.Context) error {
	// limit := c.Param("limit")

	data, err := fetchDbProcess()
	if err != nil {
		fmt.Println("err", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"data":    data,
			"success": false,
			"message": "Operation not successful",
		})
	}

	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    data,
			"success": true,
			"message": "Operation successful",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    data,
		"success": true,
		"message": "Operation successful",
	})
}

func GetProcessRealTime(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			// Write

			data, err := fetchDbProcess()
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
