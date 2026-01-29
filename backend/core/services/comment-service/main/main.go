package main

import (
	"log/slog"
	"os"
	"services/comment-service/internal"
	"services/pkg/common"
	"services/pkg/proto/comment"

	"github.com/gin-gonic/gin"
)

func main() {
	// log init
	logHandler := slog.NewJSONHandler(os.Stdout, nil)
	slog.SetDefault(slog.New(logHandler))

	// DB Connection, return handler
	slog.Info("Start DB Connection")
	hdl := dbInit()
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
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = ":9090"
	}
	listen, grpcServer := common.StartGrpcServer(grpcPort, "comment")
	comment.RegisterCommentServiceServer(grpcServer, &internal.CommentGrpcHandler{
		Svc: hdl.Service,
	})
	slog.Info("Comment gRPC Server is ready")
	if err := grpcServer.Serve(listen); err != nil {
		slog.Error("faild to serve gRPC : ", "Error", err)
	}

	// Server Init
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = ":8080"
	}
	r.Run(port)
	slog.Info("Comment Service is ready")
}
