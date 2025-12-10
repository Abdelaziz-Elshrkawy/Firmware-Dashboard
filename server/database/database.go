package database

import (
	"database/sql"
	"firmware_server/env"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// database instance for the app
var DB *sql.DB

func Connect() error {
	uri := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", env.DBuser, env.DBpassword, env.DBname)

	var err error
	DB, err = sql.Open("mysql", uri)

	if err != nil {
		return err
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("Database ping error:", err)
		return err
	}

	fmt.Println("Database Connected")

	return nil
}
