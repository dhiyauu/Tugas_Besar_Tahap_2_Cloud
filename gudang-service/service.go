package main

import (
	"errors"
	"time"
)

type SortingService struct{}

func NewSortingService() *SortingService {
	return &SortingService{}
}

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
		return errors.New("status tidak pending")
	}

	pkg.Status = "sorting"

	return nil
}

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

func (s *SortingService) GetPendingPackages(packages []Package) []Package {

	var result []Package

	for _, pkg := range packages {
		if pkg.Status == "pending" {
			result = append(result, pkg)
		}
	}

	return result
}

func (s *SortingService) ValidatePackage(pkg *Package) error {

	if pkg == nil {
		return errors.New("package nil")
	}

	if pkg.Resi == "" {
		return errors.New("resi kosong")
	}

	if pkg.UserID <= 0 {
		return errors.New("user invalid")
	}

	if pkg.Berat <= 0 {
		return errors.New("berat invalid")
	}

	if pkg.WarehouseZone == "" {
		return errors.New("warehouse kosong")
	}

	return nil
}
