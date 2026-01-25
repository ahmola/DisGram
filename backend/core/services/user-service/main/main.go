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

	// DB Connection, return handler
	slog.Info("Start DB Connection")
	hdl := db_init()
	slog.Info("DB Connection Success")

	// Gin init
	slog.Info("Gin Init")
	r := gin.Default()

	// v2 Group
	v2 := r.Group("/api/v2/users")
	slog.Info("Define Routes : v2")
	{
		v2.GET("/:id", hdl.GetUserByID)
		v2.GET("/", hdl.GetUserByUsername)
		v2.POST("/", hdl.CreateUser)
		v2.PUT("/:id", hdl.UpdateUser)
		v2.DELETE("/:id", hdl.DeleteUser)
	}

	// gRPC Init
	grpc_init(hdl)

	// Server Init
	r.Run(":8080")
	slog.Info("Server is ready")
}
