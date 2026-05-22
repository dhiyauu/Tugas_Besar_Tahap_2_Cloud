package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() error {
	connStr := "root:root@tcp(host.docker.internal:3306)/tubesdb?parseTime=true"

	var err error

	
	for i := 0; i < 10; i++ {
		db, err = sql.Open("mysql", connStr)
		if err == nil {
			err = db.Ping()
			if err == nil {
				break
			}
		}

		log.Printf("DB not ready, retrying... (%d/10)\n", i+1)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		return fmt.Errorf("database unreachable: %v", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	log.Println("Gudang database connected")
	return nil
}

func GetDB() *sql.DB {
	return db
}

func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

func main() {
	err := InitDB()
	if err != nil {
		log.Fatalf("database failed: %v", err)
	}
	defer CloseDB()

	service := NewSortingService()
	handler := NewSortingHandler(service)

	http.HandleFunc("/sort", handler.StartSort)
	http.HandleFunc("/health", handler.Health)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Gudang Service running on %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
