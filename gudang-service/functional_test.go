//go:build integration
// +build integration

package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupTestDB(t *testing.T) {
	t.Helper()

	// inisialisasi koneksi database
	err := InitDB()
	if err != nil {
		t.Fatalf("database connection failed: %v", err)
	}

	// validasi apakah database berhasil terkoneksi
	if GetDB() == nil {
		t.Fatal("database is nil")
	}

	t.Log("Gudang DB connected")
}

func setupServer() *httptest.Server {
	service := NewSortingService()
	handler := NewSortingHandler(service)

	mux := http.NewServeMux()

	// endpoint untuk sorting package
	mux.HandleFunc("/sort", handler.StartSort)

	// endpoint health check service
	mux.HandleFunc("/health", handler.Health)

	return httptest.NewServer(mux)
}

func TestFunctional_StartSort(t *testing.T) {
	setupTestDB(t)

	server := setupServer()
	defer server.Close()

	request := SortRequest{

		// isi dengan nomor resi package
		Resi: "",

		// warehouse_zone digunakan untuk menunjukkan
		// area atau zona gudang tempat package disortir
		WarehouseZone: "",

		// isi dengan status sorting
		Status: "",
	}

	// convert request ke format JSON
	body, _ := json.Marshal(request)

	resp, err := http.Post(
		server.URL+"/sort",
		"application/json",
		bytes.NewBuffer(body),
	)

	// validasi apakah request berhasil dikirim
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	// expected response ketika sorting berhasil
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200 got %d", resp.StatusCode)
	}
}
