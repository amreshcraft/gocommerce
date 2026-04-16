package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/amreshcraft/gocommerce/internal/config"
	"github.com/amreshcraft/gocommerce/pkg/logger"
)

func main() {
	// init logger
	logger.Init()

	// load config
	cfg := config.LoadConfig()

	// set gin mode
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// create router
	r := gin.Default()

	// health check
	r.GET("/", func(c *gin.Context) {
		logger.Info("Health check hit")
		c.JSON(200, gin.H{
			"message": "gocommerce running 🚀",
			"env":     cfg.Env,
		})
	})

	// start server
	addr := fmt.Sprintf(":%s", cfg.Port)
	logger.Info("Server running on " + addr)

	if err := r.Run(addr); err != nil {
		logger.Error("Failed to start server: " + err.Error())
	}
}