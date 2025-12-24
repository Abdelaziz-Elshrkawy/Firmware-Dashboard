package database

import (
	"database/sql"
	"firmware_server/env"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// database instance for the app
var DB *sql.DB

func Connect() error {
	uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", env.DBuser, env.DBpassword, env.DBHost, env.DBPort, env.DBname)

	var err error
	DB, err = sql.Open("mysql", uri)

	if err != nil {
		return err
	}

	if err := DB.Ping(); err != nil {
		return err
	}

	fmt.Println("Database Connected")

	return nil
}
