package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	var db *sql.DB
	var err error

	// Retry koneksi MySQL
	for i := 0; i < 10; i++ {

		db, err = sql.Open(
			"mysql",
			"root:root@tcp(mysql:3306)/tubesdb",
		)

		if err == nil {

			err = db.Ping()

			if err == nil {
				break
			}
		}

		fmt.Println("Waiting MySQL...")
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		panic(err)
	}

	// ==================================
	// CREATE TABLE OTOMATIS
	// ==================================
	query := `
	CREATE TABLE IF NOT EXISTS tracking_events (
		id INT AUTO_INCREMENT PRIMARY KEY,
		resi VARCHAR(100),
		lokasi VARCHAR(255),
		event VARCHAR(255),
		timestamp DATETIME
	)
	`

	_, err = db.Exec(query)

	if err != nil {
		panic(err)
	}

	fmt.Println("TRACKING TABLE READY")

	trackingRepo = MySQLRepository{
		DB: db,
	}

	// Endpoint untuk Tracking Service
	http.HandleFunc("/tracking", getTrackingHandler)
	http.HandleFunc("/tracking/event", insertTrackingEventHandler)

	// Endpoint untuk Integrasi Peta (Map/Location)
	http.HandleFunc("/distance", calculateDistanceHandler)
	http.HandleFunc("/route", calculateRouteHandler)
	http.HandleFunc("/location", getCourierLocationHandler)

	fmt.Println("Tracking Service running on :8084")
	http.ListenAndServe(":8084", nil)
}
