package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var todos = map[int]*Todo{
	1: &Todo{ID: 1, Title: "pay phone bills", Status: "active"},
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}

func getTodosHandler(c *gin.Context) {
	var items []*Todo
	stat := c.Query("status")

	for _, i := range todos {
		if stat == "" || stat == i.Status {
			items = append(items, i)
		}
	}
	c.JSON(http.StatusOK, items)
}

func getTodosByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, ok := todos[id]
	if !ok {
		// DO NOT return NIL in restAPI --> not recommended
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// func getTodosByFilterHandler(c *gin.Context) {
// 	var items []*Todo
// 	stat := c.Query("status")
// 	for _, v := range todos {
// 		if stat != "" && v.Status == stat {
// 			items = append(items, v)
// 		}
// 	}
// 	c.JSON(http.StatusOK, items)
// }

func postTodosHandler(c *gin.Context) {
	t := Todo{}
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := len(todos)
	id++

	t.ID = id
	todos[id] = &t

	c.JSON(http.StatusCreated, "created todo.")

}

func updateTodosHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := Todo{}
	if err = c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	t.ID = id
	todos[id] = &t
	c.JSON(http.StatusOK, t)
}

func deleteTodosHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	delete(todos, id)
	c.JSON(http.StatusOK, "deleted todo")
}

func main() {
	r := gin.Default()
	// r.GET("/hello", helloHandler)
	// r.Run() // listen and serve on 127.0.0.0:8080 <- default port: 8080

	r.GET("/todos", getTodosHandler)
	r.GET("/todos/:id", getTodosByIDHandler)
	r.POST("/todos", postTodosHandler)
	r.PUT("/todos/:id", updateTodosHandler)
	r.DELETE("/todos/:id", deleteTodosHandler)
	r.Run(":1234")
}
