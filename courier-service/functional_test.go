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

	err := InitDB()
	if err != nil {
		t.Fatalf("database connection failed: %v", err)
	}

	if GetDB() == nil {
		t.Fatal("database is nil")
	}

	t.Log("Courier DB connected")
}

func setupServer() *httptest.Server {
	service := NewCourierService()
	handler := NewCourierHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("/delivery", handler.StartDelivery)
	mux.HandleFunc("/courier/deliveries", handler.GetCourierDeliveries)
	mux.HandleFunc("/health", handler.Health)

	return httptest.NewServer(mux)
}

func TestFunctional_StartDelivery(t *testing.T) {
	setupTestDB(t)

	server := setupServer()
	defer server.Close()

	request := DeliveryRequest{
		Resi:         "RES001",
		CourierID:    1,
		AssignedZone: "Jakarta",
	}

	body, _ := json.Marshal(request)

	resp, err := http.Post(
		server.URL+"/delivery",
		"application/json",
		bytes.NewBuffer(body),
	)

	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200 got %d", resp.StatusCode)
	}
}