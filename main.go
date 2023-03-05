package main

import (
    "fmt"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Register a health check endpoint
    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status": "ok",
        })
    })

    // Start a background task to check the health check endpoint every 30 seconds
    ticker := time.NewTicker(30 * time.Second)
    go func() {
        for range ticker.C {
            checkHealth()
        }
    }()

    // Start the server
    r.Run(":8080")
}

// checkHealth checks the health check endpoint and logs the result
func checkHealth() {
    resp, err := http.Get("http://localhost:8080/health")
    if err != nil {
        fmt.Println("Health check failed:", err)
    } else {
        fmt.Println("Health check successful:", resp.StatusCode)
        resp.Body.Close()
    }
}
