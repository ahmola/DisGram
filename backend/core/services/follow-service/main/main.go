package main

import (
	"log/slog"
	"os"
	"services/follow-service/internal"
	"services/pkg/common"
	"services/pkg/proto/follow"

	"github.com/gin-gonic/gin"
)

func main() {
	// log init
	logHandler := slog.NewJSONHandler(os.Stdout, nil)
	slog.SetDefault(slog.New(logHandler))

	slog.Info("Start DB Connection")
	hdl := dbInit()
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
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = ":9090"
	}
	listen, grpcServer := common.StartGrpcServer(grpcPort, "follow")
	follow.RegisterFollowServiceServer(grpcServer, &internal.FollowGrpcHandler{
		Svc: hdl.Svc,
	})
	slog.Info("Follow gRPC Server is ready")
	if err := grpcServer.Serve(listen); err != nil {
		slog.Error("faild to serve gRPC : ", "Error", err)
	}
	// Server Init
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = ":8080"
	}
	r.Run(port)
	slog.Info("Follow Service is Ready")

}
