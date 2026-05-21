package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open(
		"mysql",
		"root:root@tcp(mysql:3306)/tubesdb",
	)

	if err != nil {
		panic(err)
	}

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