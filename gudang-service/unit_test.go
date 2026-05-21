package main

import (
	"testing"
)

func TestStartSortingSuccess(t *testing.T) {
	service := NewSortingService()

	pkg := &Package{
		UserID:        0,  // isi dengan user_id valid
		Resi:          "", // isi dengan nomor resi
		NamaBarang:    "", // isi dengan nama barang
		Berat:         0,  // isi dengan berat barang
		WarehouseZone: "", // isi dengan warehouse zone
		Status:        "", // isi dengan status pending
	}

	err := service.StartSorting(pkg)
	if err != nil {
		t.Errorf("StartSorting failed: %v", err)
	}

	if pkg.Status != "sorting" {
		t.Errorf("Expected status 'sorting', got '%s'", pkg.Status)
	}
}

func TestStartSortingNil(t *testing.T) {
	service := NewSortingService()

	err := service.StartSorting(nil)
	if err == nil {
		t.Error("Expected error for nil package, got nil")
	}
}

func TestStartSortingEmptyResi(t *testing.T) {
	service := NewSortingService()

	pkg := &Package{
		UserID:        0,  // isi dengan user_id valid
		Resi:          "", // kosongkan nomor resi
		WarehouseZone: "", // isi dengan warehouse zone
		Status:        "", // isi dengan status pending
	}

	err := service.StartSorting(pkg)
	if err == nil {
		t.Error("Expected error for empty resi, got nil")
	}
}

func TestStartSortingEmptyWarehouseZone(t *testing.T) {
	service := NewSortingService()

	pkg := &Package{
		UserID:        0,  // isi dengan user_id valid
		Resi:          "", // isi dengan nomor resi
		WarehouseZone: "", // kosongkan warehouse zone
		Status:        "", // isi dengan status pending
	}

	err := service.StartSorting(pkg)
	if err == nil {
		t.Error("Expected error for empty warehouse_zone, got nil")
	}
}

func TestStartSortingInvalidStatus(t *testing.T) {
	service := NewSortingService()

	pkg := &Package{
		UserID:        0,  // isi dengan user_id valid
		Resi:          "", // isi dengan nomor resi
		WarehouseZone: "", // isi dengan warehouse zone
		Status:        "", // isi dengan status selain pending
	}

	err := service.StartSorting(pkg)
	if err == nil {
		t.Error("Expected error for non-pending package, got nil")
	}
}

func TestCompleteSortingSuccess(t *testing.T) {
	service := NewSortingService()

	pkg := &Package{
		Resi:          "", // isi dengan nomor resi
		Status:        "", // isi dengan status sorting
		WarehouseZone: "", // isi dengan warehouse zone
	}

	err := service.CompleteSorting(pkg)
	if err != nil {
		t.Errorf("CompleteSorting failed: %v", err)
	}

	if pkg.Status != "ready" {
		t.Errorf("Expected status 'ready', got '%s'", pkg.Status)
	}

	if pkg.SortedAt == nil {
		t.Error("SortedAt should not be nil")
	}
}

func TestCompleteSortingNil(t *testing.T) {
	service := NewSortingService()

	err := service.CompleteSorting(nil)
	if err == nil {
		t.Error("Expected error for nil package, got nil")
	}
}

func TestCompleteSortingNotSorting(t *testing.T) {
	service := NewSortingService()

	pkg := &Package{
		Resi:          "", // isi dengan nomor resi
		Status:        "", // isi dengan status selain sorting
		WarehouseZone: "", // isi dengan warehouse zone
	}

	err := service.CompleteSorting(pkg)
	if err == nil {
		t.Error("Expected error for non-sorting package, got nil")
	}
}

func TestGetPendingPackagesSuccess(t *testing.T) {
	service := NewSortingService()

	packages := []Package{
		{
			Resi:  "", // isi dengan nomor resi
			Status: "", // isi dengan status pending
		},
		{
			Resi:  "", // isi dengan nomor resi
			Status: "", // isi dengan status sorting
		},
		{
			Resi:  "", // isi dengan nomor resi
			Status: "", // isi dengan status pending
		},
	}

	pending := service.GetPendingPackages(packages)

	// sesuaikan jumlah expected package pending
	if len(pending) != 0 {
		t.Errorf("Expected pending packages count mismatch, got %d", len(pending))
	}
}

func TestGetPendingPackagesEmpty(t *testing.T) {
	service := NewSortingService()

	packages := []Package{
		{
			Resi:  "", // isi dengan nomor resi
			Status: "", // isi dengan status sorting
		},
		{
			Resi:  "", // isi dengan nomor resi
			Status: "", // isi dengan status ready
		},
	}

	pending := service.GetPendingPackages(packages)

	if len(pending) != 0 {
		t.Errorf("Expected 0 pending packages, got %d", len(pending))
	}
}

func TestValidatePackageSuccess(t *testing.T) {
	service := NewSortingService()

	pkg := &Package{
		Resi:          "", // isi dengan nomor resi
		UserID:        0,  // isi dengan user_id valid
		Berat:         0,  // isi dengan berat barang
		WarehouseZone: "", // isi dengan warehouse zone
	}

	err := service.ValidatePackage(pkg)
	if err != nil {
		t.Errorf("ValidatePackage failed: %v", err)
	}
}

func TestValidatePackageNil(t *testing.T) {
	service := NewSortingService()

	err := service.ValidatePackage(nil)
	if err == nil {
		t.Error("Expected error for nil package, got nil")
	}
}

func TestValidatePackageEmptyResi(t *testing.T) {
	service := NewSortingService()

	pkg := &Package{
		Resi:          "", // kosongkan nomor resi
		UserID:        0,  // isi dengan user_id valid
		Berat:         0,  // isi dengan berat barang
		WarehouseZone: "", // isi dengan warehouse zone
	}

	err := service.ValidatePackage(pkg)
	if err == nil {
		t.Error("Expected error for empty resi, got nil")
	}
}

func TestValidatePackageInvalidUserID(t *testing.T) {
	service := NewSortingService()

	pkg := &Package{
		Resi:          "", // isi dengan nomor resi
		UserID:        0,  // isi dengan user_id tidak valid
		Berat:         0,  // isi dengan berat barang
		WarehouseZone: "", // isi dengan warehouse zone
	}

	err := service.ValidatePackage(pkg)
	if err == nil {
		t.Error("Expected error for invalid user_id, got nil")
	}
}

func TestValidatePackageInvalidWeight(t *testing.T) {
	service := NewSortingService()

	pkg := &Package{
		Resi:          "", // isi dengan nomor resi
		UserID:        0,  // isi dengan user_id valid
		Berat:         0,  // isi dengan berat tidak valid
		WarehouseZone: "", // isi dengan warehouse zone
	}

	err := service.ValidatePackage(pkg)
	if err == nil {
		t.Error("Expected error for zero weight, got nil")
	}
}

func TestValidatePackageEmptyWarehouseZone(t *testing.T) {
	service := NewSortingService()

	pkg := &Package{
		Resi:          "", // isi dengan nomor resi
		UserID:        0,  // isi dengan user_id valid
		Berat:         0,  // isi dengan berat barang
		WarehouseZone: "", // kosongkan warehouse zone
	}

	err := service.ValidatePackage(pkg)
	if err == nil {
		t.Error("Expected error for empty warehouse_zone, got nil")
	}
}

// Benchmark StartSorting
func BenchmarkStartSorting(b *testing.B) {
	service := NewSortingService()

	pkg := &Package{
		UserID:        0,  // isi dengan user_id valid
		Resi:          "", // isi dengan nomor resi
		WarehouseZone: "", // isi dengan warehouse zone
		Status:        "", // isi dengan status pending
	}

	/* function dijalankan berulang kali untuk mengukur performa sorting */
	for i := 0; i < b.N; i++ {
		service.StartSorting(pkg)

		// reset status package
		pkg.Status = ""
	}
}

// Benchmark ValidatePackage
func BenchmarkValidatePackage(b *testing.B) {
	service := NewSortingService()

	pkg := &Package{
		Resi:          "", // isi dengan nomor resi
		UserID:        0,  // isi dengan user_id valid
		Berat:         0,  // isi dengan berat barang
		WarehouseZone: "", // isi dengan warehouse zone
	}

	/* function dijalankan berulang kali untuk mengukur performa validasi package */
	for i := 0; i < b.N; i++ {
		service.ValidatePackage(pkg)
	}
}
