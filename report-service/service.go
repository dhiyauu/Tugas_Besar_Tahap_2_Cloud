package main

import "context"

type ReportRepository interface {
	GetDailyReport(ctx context.Context, date string) (*DailyReport, error)
}

type ReportService struct {
	repo ReportRepository
}

func NewReportService(repo ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetDailyReport(ctx context.Context, date string) (*DailyReport, error) {
	return s.repo.GetDailyReport(ctx, date)
}
