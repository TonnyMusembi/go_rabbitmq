package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type TestInput struct {
	Message string `json:"message"`
}


func main() {
	// Create a new Gin router
	r := gin.Default()

	// Define a route
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, test api!",
		})
		 r.NoRoute(func(c *gin.Context) {
        // In gin this is how you return a JSON response
        c.JSON(404, gin.H{"message": "Not found"})
    })
	})
	r.POST("/test", func(c *gin.Context) {
		var input TestInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	message := input.Message
		c.JSON(http.StatusOK, gin.H{
			"message": message,
		})
	})

	// Start the server on port 8080
	r.Run(":8089")
}
