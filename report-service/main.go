package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
)

type MySQLReportRepository struct {
	db *sql.DB
}

func (r MySQLReportRepository) GetDailyReport(ctx context.Context, date string) (*DailyReport, error) {
	return &DailyReport{
		TotalPaket:  10,
		Delivered:   8,
		Pending:     1,
		Terlambat:   1,
		RataRataETA: 2.5,
	}, nil
}

func main() {

	ConnectDB()

	repo := MySQLReportRepository{db: DB}
	svc := NewReportService(repo)
	h := NewReportHandler(svc)

	http.HandleFunc("/report/daily", h.DailyReport)

	log.Println("Running on :8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
}
