package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.POST("/Todo", AddTodo)
	r.GET("/Todo", ListOfTodo)
	r.PATCH("/Todo/{:id}", UpdateTodo)
	r.DELETE("/Todo/{:id}", DeleteTodo)
}
