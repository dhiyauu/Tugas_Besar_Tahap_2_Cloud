package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbUser = "root"
	dbPass = "root"
	dbHost = "host.docker.internal"
	dbPort = "3306"
	dbName = "tubesdb"
)

func TestCreatePayment_Functional(t *testing.T) {

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		t.Fatal(err)
	}

	body := []byte(`{
		"order_id":1,
		"method":"E-Wallet",
		"option":"DANA"
	}`)

	resp, err := http.Post(
		"http://host.docker.internal:8088/payment",
		"application/json",
		bytes.NewBuffer(body),
	)

	if err != nil {
		t.Fatal(err)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	if resp.StatusCode != 200 {
		t.Fatalf("failed: %+v", result)
	}

	var count int

	err = db.QueryRow(
		"SELECT COUNT(*) FROM payments",
	).Scan(&count)

	if err != nil {
		t.Fatal(err)
	}

	if count < 1 {
		t.Fatal("payment tidak masuk database")
	}
}
