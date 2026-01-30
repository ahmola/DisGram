package main

import (
	"log/slog"
	"os"
	"services/pkg/common"
	"services/pkg/proto/post"
	"services/post-service/internal"

	"github.com/gin-gonic/gin"

	_ "services/post-service/main/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title				Post API
// @version				2.0
// @description			API built with Gin
// @host				localhost:8083
// @BasePath			/api/v2
func main() {
	// log init
	logHandler := slog.NewJSONHandler(os.Stdout, nil)
	slog.SetDefault(slog.New(logHandler))

	// DB Connection, return handler
	slog.Info("Start DB Connection")
	hdl := dbMigrate()
	slog.Info("DB Connection Success")

	// Gin init
	slog.Info("Gin Init")
	r := gin.Default()

	// v2 Group
	v2 := r.Group("/api/v2/posts")
	slog.Info("Define Routes : v2")
	{
		v2.GET("/:id", hdl.GetPostById)
		v2.POST("/", hdl.CreatePost)
		v2.PUT("/:id", hdl.UpdatePost)
		v2.DELETE("/:id", hdl.DeletePost)

		v2.GET("/:postId/likes", hdl.GetAllLikesByPostID)
		v2.POST("/likes", hdl.CreateLike)
		v2.DELETE("/likes/:id", hdl.DeleteLike)
	}

	// gRPC Init
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = ":9090"
	}
	listen, grpcServer := common.StartGrpcServer(grpcPort, "post")
	post.RegisterPostServiceServer(grpcServer, &internal.PostGrpcHandler{
		Svc: hdl.Svc,
	})
	slog.Info("gRPC Init success", "addr", listen.Addr().String(), "info", grpcServer.GetServiceInfo())

	// Run gRPC by go Routine(Async)
	common.RunGrpcWithGoRoutine(listen, grpcServer)
	slog.Info("Post gRPC Server is ready")

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

	slog.Info("Post Service is ready")
	r.Run(port)
}
