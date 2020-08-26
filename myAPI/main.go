package main

import (
	"log"
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
	// c.JSON(http.StatusOK, "hi")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte("hi"))
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

func setupRouter() *gin.Engine {
	r := gin.Default()
	// r.GET("/hello", helloHandler)
	// r.Run() // listen and serve on 127.0.0.0:8080 <- default port: 8080

	// add LOGGING middleware
	r.Use(func(c *gin.Context) {
		log.Println("Start LOGGING Middleware")
		c.Next()
		log.Println("Exit LOGGING Middleware")
	})

	// alternatively for CHAINED middleware
	// the execution order will be accordingly
	// r.Use([]gin.HandlerFunc{mw1,mw2}) ; mw is of type func(c *gin.Context)

	// add AUTHEN middleware
	r.Use(func(c *gin.Context) {
		log.Println("Start authentication")
		if c.GetHeader("Authorization") != "Bearer 1234" {
			c.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))

			// need to abort gin
			c.Abort()
			return
		}

		// move to next component behind AUTHEN middleware
		c.Next()
		log.Println("Exit authentication")
	})

	// group setting for version control
	v1 := r.Group("/v1")
	v1.GET("/todos", getTodosHandler)
	v1.GET("/todos/:id", getTodosByIDHandler)
	v1.GET("/hello", helloHandler)

	v1.POST("/todos", postTodosHandler)
	v1.PUT("/todos/:id", updateTodosHandler)
	v1.DELETE("/todos/:id", deleteTodosHandler)
	return r
}

func main() {
	r := setupRouter()
	r.Run(":1234")
}
