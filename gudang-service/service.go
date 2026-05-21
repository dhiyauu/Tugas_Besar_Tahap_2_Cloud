package main

import (
	"errors"
	"time"
)

type SortingService struct {
}

func NewSortingService() *SortingService {
	return &SortingService{}
}

// StartSorting
func (s *SortingService) StartSorting(pkg *Package) error {

	if pkg == nil {
		return errors.New("package nil")
	}

	if pkg.Resi == "" {
		return errors.New("resi kosong")
	}

	if pkg.WarehouseZone == "" {
		return errors.New("warehouse zone kosong")
	}

	if pkg.Status != "pending" {
		return errors.New("status tidak pending")
	}

	pkg.Status = "sorting"

	return nil
}

// CompleteSorting
func (s *SortingService) CompleteSorting(pkg *Package) error {

	if pkg == nil {
		return errors.New("package nil")
	}

	if pkg.Status != "sorting" {
		return errors.New("status tidak sorting")
	}

	now := time.Now()

	pkg.Status = "ready"
	pkg.SortedAt = &now

	return nil
}

// GetPendingPackages
func (s *SortingService) GetPendingPackages(packages []Package) []Package {

	var result []Package

	for _, p := range packages {

		if p.Status == "pending" {
			result = append(result, p)
		}
	}

	return result
}

// ValidatePackage
func (s *SortingService) ValidatePackage(pkg *Package) error {

	if pkg == nil {
		return errors.New("package nil")
	}

	if pkg.Resi == "" {
		return errors.New("resi kosong")
	}

	if pkg.UserID <= 0 {
		return errors.New("user id invalid")
	}

	if pkg.Berat <= 0 {
		return errors.New("berat invalid")
	}

	if pkg.WarehouseZone == "" {
		return errors.New("warehouse zone kosong")
	}

	return nil
}
