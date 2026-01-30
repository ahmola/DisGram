package main

import (
	"log/slog"
	"os"
	"services/follow-service/internal"
	"services/pkg/common"
	"services/pkg/proto/follow"

	"github.com/gin-gonic/gin"

	_ "services/follow-service/main/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title				Follow API
// @version				2.0
// @description			API built with Gin
// @host				localhost:8082
// @BasePath			/api/v2
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
		v2.POST("/", hdl.CreateFollow)
		v2.GET("/followees", hdl.GetFolloweesByFollowerID)
		v2.GET("/followers", hdl.GetFollowersByFolloweeID)
		v2.DELETE("/:id", hdl.DeleteFollow)
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
	slog.Info("gRPC Init success", "addr", listen.Addr().String())

	// Run gRPC by go Routine(Async)
	common.RunGrpcWithGoRoutine(listen, grpcServer)
	slog.Info("Follow gRPC Server is ready")

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

	slog.Info("Follow Service is Ready")
	r.Run(port)
}
