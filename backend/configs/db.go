package configs

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var Db sql.DB

const PATH = "../db/processes.db"

func runMigrations() {
}

func ConnectDb() {
	var err error

	// Create new sql file if file deos not exist
	if _, err := os.Stat(PATH); err != nil {
		// Make directory
		err := os.Mkdir("../db", 0o777)
		if err != nil {
			fmt.Println("direcorty err", err)
			panic(err)
		}
		file, err := os.Create(PATH)
		if err != nil {
			fmt.Println("err", err)
			panic(err)
		}

		file.Close()
	}

	fmt.Println("golang", os.Getenv("SQLITE_DB"))
	db, err := sql.Open("sqlite3", PATH)
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

	// Run migrations

	Db = *db

	Db.Exec("DROP table processes")

	_, derr := Db.Exec(`CREATE TABLE IF NOT EXISTS processes (
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

	if derr != nil {
		fmt.Println("derr", derr)
	}

	fmt.Println("Creating completed")
}
