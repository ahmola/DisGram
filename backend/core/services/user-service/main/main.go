package main

import (
	"log/slog"
	"os"
	"services/pkg/common"
	"services/pkg/proto/user"
	"services/user-service/internal"

	"github.com/gin-gonic/gin"

	_ "services/user-service/main/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title				User API
// @version				2.0
// @description			API built with Gin
// @host				localhost:8081
// @BasePath			/api/v2
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
	slog.Info("gRPC Init success", "addr", listen.Addr().String(), "info", grpcServer.GetServiceInfo())

	// Run gRPC by go Routine(Async)
	common.RunGrpcWithGoRoutine(listen, grpcServer)
	slog.Info("User gRPC Server is ready")

	// Server Init
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = ":8080"
	}
	slog.Info("Check Environment Variable and port : ", "SERVER_PORT", port)

	// Swagger UI
	slog.Info("Register Swagger Handler")
	r.GET("/swagger/*any", func(c *gin.Context) {
		slog.Debug("Swagger page requested", "url", c.Request.URL.String())
		c.Next()
	}, ginSwagger.WrapHandler(swaggerFiles.Handler))
	slog.Info("OpenAPI Docs Opened!")

	slog.Info("User Service is ready")
	r.Run(port)
}
