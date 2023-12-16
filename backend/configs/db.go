package configs

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var Db sql.DB

func ConnectDb() {
	file := "/Users/joshuaoluikpe/processes.db"

	var err error
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		// Log to error text file
		fmt.Println("err", err)
		panic(err)
	}

	if err = db.Ping(); err != nil {
		// log panic
		fmt.Println("verr", err)
		panic(err)
	}

	fmt.Println("Db connected sucessfully")

	Db = *db

	Db.Exec("DROP table processes")

	Db.Exec(`CREATE TABLE IF NOT EXISTS processes (
		id INTEGER NOT NULL PRIMARY KEY,
		user text,
		pid integer NOT NULL UNIQUE,
		cpuUsage integer,
		memoryPercentageUsage integer,
		virtualMemorySize integer,
		residentMemorySize integer,
		tty text,
		state text,
		started text,
		totalTime text,
		command text,
		createdAt timestamp,
		updatedAt timestamp
	);`)
}
