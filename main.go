package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	cache = make(map[string][]byte)
	mu    sync.RWMutex
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

	router.GET("/api", func(c *gin.Context) {
		id := c.Query("id") // Mỗi ID khác nhau sẽ tốn thêm RAM

		// 1. Kiểm tra xem ID này đã được "đổ" dữ liệu vào RAM chưa
		mu.RLock()
		data, exists := cache[id]
		mu.RUnlock()

		if !exists {
			mu.Lock()
			// Kiểm tra lại lần nữa (double-check) để tránh ghi đè khi nhiều request cùng vào
			if _, stillNotExists := cache[id]; stillNotExists {
				// Tạo một mảng byte khoảng 1MB
				// Nếu muốn RAM tăng nhanh hơn, hãy tăng con số 1024 * 1024 lên
				newData := make([]byte, 5*1024*1024)

				// "Ghi" cái gì đó vào để đảm bảo hệ điều hành thực sự cấp phát RAM
				for i := 0; i < len(newData); i += 4096 {
					newData[i] = 1
				}

				cache[id] = newData
				data = newData
				fmt.Printf("Đã nạp thêm 1MB vào RAM cho ID: %s. Tổng số keys: %d\n", id, len(cache))
			}
			mu.Unlock()
		}

		c.JSON(200, gin.H{
			"message": "Data loaded in memory",
			"id":      id,
			"size_mb": len(data) / 1024 / 1024,
		})
	})

	router.Run(":8080")
}
