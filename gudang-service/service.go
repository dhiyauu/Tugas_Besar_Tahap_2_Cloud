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

// StartSorting - Mulai proses sorting package
func (s *SortingService) StartSorting(pkg *Package) error {

	if pkg == nil {
		return errors.New("package nil")
	}

	if pkg.Resi == "" {
		return errors.New("resi kosong")
	}

	if pkg.WarehouseZone == "" {
		return errors.New("warehouse kosong")
	}

	if pkg.Status != "pending" {
		return errors.New("status tidak valid")
	}

	pkg.Status = "sorting"

	return nil
}

// CompleteSorting - Selesaikan proses sorting
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

// GetPendingPackages - Ambil semua package dengan status pending
func (s *SortingService) GetPendingPackages(packages []Package) []Package {

	var result []Package

	for _, pkg := range packages {

		if pkg.Status == "pending" {
			result = append(result, pkg)
		}
	}

	return result
}

// ValidatePackage - Validasi package data
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
		return errors.New("weight invalid")
	}

	if pkg.WarehouseZone == "" {
		return errors.New("warehouse kosong")
	}

	return nil
}
