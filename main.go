package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Product struct {
	Name      string
	Price     int
	Available bool
}

func main() {

	connStr := "postgres://postgres:asdfg@localhost:5432/gopgtest?sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	createProductTable(db)

	product := Product{"Book", 15, true}

	pk := insertProduct(db, product)

	fmt.Printf("ID: %d\n", pk)
}

func createProductTable(db *sql.DB) {
	query := ` CREATE TABLE IF NOT EXISTS Product (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		price NUMERIC(6,2) NOT NULL,
		available BOOLEAN,
		created timestamp DEFAULT NOW()
	)`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func insertProduct(db *sql.DB, product Product) int {
	query := `INSERT INTO Product (name, price, available)
		VALUES ($1, $2, $3) RETURNING id`

	var pk int
	err := db.QueryRow(query, product.Name, product.Price, product.Available).Scan(&pk)
	if err != nil {
		log.Fatal(err)
	}
	return pk
}
