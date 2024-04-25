package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// r := gin.Default()
	// r.POST("/Todo", AddTodo)
	// r.GET("/Todo", ListOfTodo)
	// r.PATCH("/Todo/{:id}", UpdateTodo)
	// r.DELETE("/Todo/{:id}", DeleteTodo)

	connStr := "postgres://postgres:asdfg@localhost:5432/gopgtest?sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
