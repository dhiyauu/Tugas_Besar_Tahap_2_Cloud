

package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// Interface untuk mocking di test
type SortingServiceInterface interface {
	StartSorting(pkg *Package) error
	CompleteSorting(pkg *Package) error
	GetPendingPackages(packages []Package) []Package
	ValidatePackage(pkg *Package) error
}

type SortingHandler struct {
	service SortingServiceInterface
}

func NewSortingHandler(service SortingServiceInterface) *SortingHandler {
	return &SortingHandler{service: service}
}

// POST /sort
func (h *SortingHandler) StartSort(w http.ResponseWriter, r *http.Request) {

	var req SortRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Resi == "" {
		http.Error(w, "Resi is required", http.StatusBadRequest)
		return
	}

	if req.WarehouseZone == "" {
		http.Error(w, "Warehouse zone is required", http.StatusBadRequest)
		return
	}

	now := time.Now()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":         "sorting started",
		"resi":           req.Resi,
		"warehouse_zone": req.WarehouseZone,
		"sorted_at":      now,
	})
}

// GET /health
func (h *SortingHandler) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"status": "healthy",
	})
}
