package main

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// log init
	logHandler := slog.NewJSONHandler(os.Stdout, nil)
	slog.SetDefault(slog.New(logHandler))

	slog.Info("Start DB Connection")
	hdl := db_init()
	slog.Info("DB Connection Success")

	// Gin Init
	slog.Info("Gin Init")
	r := gin.Default()

	// API V2
	v2 := r.Group("/api/v2/follows")
	slog.Info("Define Routes: v2")
	{
		v2.POST("/")
		v2.GET("/")
		v2.DELETE("/")
	}

	// gRPC Init
	grpc_init(hdl)

	r.Run(":8080")
	slog.Info("Follow Service is Ready")

}
