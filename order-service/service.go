package main

import "errors"

import (
	"fmt"
	"net/http"
	"time"
)

var orders []Order
var nextID = 1

var UserServiceURL = "http://host.docker.internal:8081"

type Validator interface {
	CheckUser(userID int, token string) bool
}

type RealValidator struct{}

func (v RealValidator) CheckUser(userID int, token string) bool {
	req, _ := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/profile?id=%d", UserServiceURL, userID),
		nil,
	)

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("PROFILE ERROR:", err)
		return false
	}

	fmt.Println("PROFILE STATUS:", resp.StatusCode)

	return resp.StatusCode == 200
}

func GenerateResi() string {
	return ""
}

func CalculateETA() string {
	return ""
}

func CreateOrder(
	req Order,
	token string,
	validator Validator,
	repo OrderRepository,
) (Order, error) {

	valid := validator.CheckUser(req.UserID, token)

	if !valid {
		return Order{}, errors.New("user tidak valid")
	}

	req.Status = "created"

	err := repo.Save(req)

	if err != nil {
		return Order{}, err
	}

	return req, nil
}

func GetOrder(id int) *Order {
	return nil
}

func UpdateOrderStatus(id int, status string) bool {
	return false
}

func GetETA(id int) string {
	return ""
}
