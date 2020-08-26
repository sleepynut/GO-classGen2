package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}

func getTodosHandler(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{
	// 	"id":     1,
	// 	"title":  "pay phone bills",
	// 	"status": "active",
	// })
	var items []*Todo
	for _, i := range todos {
		items = append(items, i)
	}
	c.JSON(http.StatusOK, items)
}

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var todos = map[int]*Todo{
	1: &Todo{ID: 1, Title: "pay phone bills", Status: "Active"},
}

func main() {
	r := gin.Default()
	// r.GET("/hello", helloHandler)
	// r.Run() // listen and serve on 127.0.0.0:8080 <- default port: 8080

	r.GET("/todos", getTodosHandler)
	r.Run(":1234")
}
