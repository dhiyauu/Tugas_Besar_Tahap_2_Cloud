package main

import (
	"encoding/json"
	"net/http"
)

func createPaymentHandler(w http.ResponseWriter, r *http.Request) {

	var req Payment

	json.NewDecoder(r.Body).Decode(&req)

	repo := MySQLRepository{}

	payment, err := CreatePayment(
		req,
		repo,
	)

	if err != nil {
		w.WriteHeader(400)

		json.NewEncoder(w).Encode(
			map[string]string{
				"error": err.Error(),
			},
		)

		return
	}

	json.NewEncoder(w).Encode(payment)
}

func getPaymentHandler(w http.ResponseWriter, r *http.Request) {}

func updatePaymentHandler(w http.ResponseWriter, r *http.Request) {}