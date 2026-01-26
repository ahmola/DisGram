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
	v2 := r.Group("/api/v2/comments")
	slog.Info("Define Routes : v2")
	{
		v2.GET("/:id", hdl.GetCommentByID)
		v2.GET("/", hdl.GetComments)
		v2.POST("/", hdl.CreateComment)
		v2.PUT("/:id", hdl.UpdateComment)
		v2.DELETE("/:id", hdl.DeleteComment)
	}

	// gRPC Init
	grpc_init(hdl)

	// Server Init
	r.Run(":8080")
	slog.Info("Comment Service is ready")
}
