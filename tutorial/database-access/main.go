package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/dotenv-org/godotenvvault"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// load .env
	if err := godotenvvault.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}
	// Get a
	// database
	// handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}
