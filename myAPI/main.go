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
	c.JSON(http.StatusOK, gin.H{
		"message": "hello Gin GET todos",
	})
}

func main() {
	r := gin.Default()
	r.GET("/todos", getTodosHandler)
	r.Run(":1234") // listen and serve on 127.0.0.0:8080
}
