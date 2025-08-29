package handlers

import (
    "net/http"
    "os"
    "github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        apiKey := c.GetHeader("Authorization")
        validKey := os.Getenv("GATEWAY_API_KEY")
        if apiKey == "" || apiKey != validKey {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            return
        }
        c.Next()
    }
}