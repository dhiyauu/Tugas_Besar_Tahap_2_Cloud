package main

import (
	"context"
)

type MockReportRepo struct{}

func (m MockReportRepo) GetDailyReport(ctx context.Context, date string) (*DailyReport, error) {
	return &DailyReport{
		TotalPaket:  10,
		Delivered:   8,
		Pending:     1,
		Terlambat:   1,
		RataRataETA: 2.5,
	}, nil
}
