package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestStartSort_AllValid(t *testing.T) {
	handler := NewSortingHandler(nil)

	body := `{
		"resi":"", // isi dengan nomor resi
		"warehouse_zone":"", // isi dengan warehouse zone (untuk menunjukkan zona atau area gudang tempat paket disortir )
		"status":"" // isi dengan status sorting
	}`

	req := httptest.NewRequest("POST", "/sort", strings.NewReader(body))
	w := httptest.NewRecorder()

	handler.StartSort(w, req)

	// expected response ketika request valid
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestStartSort_AllErrorPaths(t *testing.T) {
	handler := NewSortingHandler(nil)

	tests := []string{
		// isi dengan format JSON tidak valid
		``,

		// kosongkan nomor resi
		`{
			"resi":"",
			"warehouse_zone":"",
			"status":""
		}`,

		// kosongkan warehouse zone
		`{
			"resi":"",
			"warehouse_zone":"",
			"status":""
		}`,

		// kosongkan status
		`{
			"resi":"",
			"warehouse_zone":"",
			"status":""
		}`,

		// isi dengan status selain sorting
		`{
			"resi":"",
			"warehouse_zone":"",
			"status":""
		}`,
	}

	for _, body := range tests {
		req := httptest.NewRequest("POST", "/sort", strings.NewReader(body))
		w := httptest.NewRecorder()

		handler.StartSort(w, req)

		// expected response ketika request tidak valid
		if w.Code != http.StatusBadRequest {
			t.Errorf("expected 400, got %d", w.Code)
		}
	}
}

func TestHealth_OK(t *testing.T) {
	handler := NewSortingHandler(nil)

	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	handler.Health(w, req)

	// expected response ketika service berjalan normal
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}
