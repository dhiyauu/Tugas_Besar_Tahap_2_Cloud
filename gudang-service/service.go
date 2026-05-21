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

	pkg.Status = "sorting"

	return nil
}

func (s *SortingService) CompleteSorting(pkg *Package) error {

	if pkg == nil {
		return errors.New("package nil")
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

	return nil
}
