package configs

import (
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

const (
	PATH      = "../db/"
	FULL_PATH = "../db/processes.db"
)

func setupDb() {
	var err error

	// Create new sql file if file deos not exist

	_, err = os.Stat(PATH)

	if err == nil {
		err := os.RemoveAll(PATH)
		if err != nil {
			log.Fatal("Error removing:", err)
		}
	}

	err = os.Mkdir(PATH, 0o777)

	if err != nil {
		log.Fatal("Creating direcorty err", err)
		panic(err)
	}
	file, err := os.Create(FULL_PATH)
	if err != nil {
		log.Fatal("Error in creating DB path", err)
		panic(err)
	}

	file.Close()
}

func runMigrations() {
	instance, err := sqlite3.WithInstance(Db, &sqlite3.Config{})
	if err != nil {
		log.Fatal(err)
	}

	fSrc, err := (&file.File{}).Open("./migrations")
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithInstance("file", fSrc, "sqlite3", instance)
	if err != nil {
		log.Fatal(err)
	}

	// modify for Down
	if err := m.Up(); err != nil {
		log.Println("Error on migration", err)
	}

	log.Println("Db created successfully")
}

func ConnectDb() {
	setupDb()
	db, err := sql.Open("sqlite3", FULL_PATH)
	if err != nil {
		log.Fatal("err", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Error on connecting db", err)
	}

	log.Println("Db connected sucessfully")

	Db = db

	// Run migrations
	runMigrations()

	log.Println("Db migration was successful")
}
