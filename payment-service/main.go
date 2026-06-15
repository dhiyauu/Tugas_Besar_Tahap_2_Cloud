package main

import (
	"fmt"
	"net/http"
)

func main() {

	ConnectDB()

	http.HandleFunc("/payment", createPaymentHandler)
	http.HandleFunc("/payment/get", getPaymentHandler)
	http.HandleFunc("/payment/status", updatePaymentHandler)

	fmt.Println("Payment Service running on :8088")

	http.ListenAndServe(":8088", nil)
}
