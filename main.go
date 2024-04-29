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

	var name string
	var available bool
	var price float64

	query := "SELECT name, available, price FROM product WHERE id = $1"
	err = db.QueryRow(query, pk).Scan(&name, &available, &price)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("No row found with id %d", 111)
		}
		log.Fatal(err)
	}

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Name: %t\n", available)
	fmt.Printf("Name: %f\n", price)
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
