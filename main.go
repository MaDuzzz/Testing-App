package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func main() {
	router := gin.Default()

	router.GET("/api/hello", func(c *gin.Context) {
		shouldDelay := rand.Intn(2) // 0 or 1

		if shouldDelay == 1 {
			delay := 50 + rand.Intn(51) // 50-100 milliseconds
			time.Sleep(time.Duration(delay) * time.Millisecond)
		}

		response := Message{
			Message: "Hello from Go Backend with Gin!",
			Status:  "success",
		}
		c.JSON(http.StatusOK, response)
	})

	router.Run(":8080")
}
