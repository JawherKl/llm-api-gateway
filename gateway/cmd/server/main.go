package main

import (
    "github.com/gin-gonic/gin"
    "github.com/JawherKl/gateway/internal/config"
    "github.com/JawherKl/gateway/internal/handlers"
)

func main() {
    cfg := config.Load()
    r := gin.Default()

    // Register the Auth middleware globally or for specific routes
    r.Use(handlers.AuthMiddleware())

    // Routes
    r.POST("/gateway/query", handlers.QueryHandler(cfg))

    r.Run(":" + cfg.ServerPort)
}
