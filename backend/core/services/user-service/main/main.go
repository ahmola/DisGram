package main

import (
	"log/slog"
	"os"
	"services/pkg/common"
	"services/pkg/proto/user"
	"services/user-service/internal"

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
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = ":9090"
	}
	listen, grpcServer := common.StartGrpcServer(grpcPort, "user")
	user.RegisterUserServiceServer(grpcServer, &internal.UserGrpcHandler{
		Svc: hdl.Svc,
	})
	slog.Info("User gRPC Server is ready")
	if err := grpcServer.Serve(listen); err != nil {
		slog.Error("faild to serve gRPC : ", "Error", err)
	}

	// Server Init
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = ":8080"
	}
	r.Run(port)
	slog.Info("User Server is ready")
}
