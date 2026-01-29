package main

import (
	"log/slog"
	"os"
	"services/pkg/common"
	"services/pkg/proto/post"
	"services/post-service/internal"

	"github.com/gin-gonic/gin"
)

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

		v2.GET("/likes/:postId", hdl.GetAllLikesByPostID)
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
	slog.Info("Post gRPC Server is ready")
	if err := grpcServer.Serve(listen); err != nil {
		slog.Error("faild to serve gRPC : ", "Error", err)
	}

	// Server Init
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = ":8080"
	}
	r.Run(port)
	slog.Info("Post Service is ready")
}
