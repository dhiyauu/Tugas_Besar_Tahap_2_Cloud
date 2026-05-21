package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	// default untuk local / test
	if host == "" {
		host = "localhost"
	}

	if port == "" {
		port = "3306"
	}

	dsn := fmt.Sprintf(
		"root:root@tcp(%s:%s)/tubesdb",
		host,
		port,
	)

	var err error

	for i := 0; i < 10; i++ {

		DB, err = sql.Open("mysql", dsn)

		if err == nil {
			err = DB.Ping()
		}

		if err == nil {
			fmt.Println("REPORT DB CONNECTED")
			return
		}

		fmt.Println("Waiting MySQL...", err)
		time.Sleep(5 * time.Second)
	}
}