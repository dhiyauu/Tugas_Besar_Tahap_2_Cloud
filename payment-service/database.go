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
			createTable()
			return
		}

		time.Sleep(5 * time.Second)
	}

	panic(err)
}

func createTable() {

	query := `
	CREATE TABLE IF NOT EXISTS payments (
		payment_id INT AUTO_INCREMENT PRIMARY KEY,
		order_id INT,
		method VARCHAR(50),
		payment_option VARCHAR(50),
		status VARCHAR(50)
	)
	`

	_, err := DB.Exec(query)

	if err != nil {
		panic(err)
	}
}
