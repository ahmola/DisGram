package main

import (
	"net"
	"os"

	"log/slog"
	"services/pkg/proto/user"
	"services/user-service/internal"

	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"

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

	// gRPC Server Init
	slog.Info("Start Listening ", "user gRPC Server")
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		slog.Error("failed to open tcp ", grpcPort, "Error", err)
		os.Exit(1)
	}
	slog.Info("Listening: ", listen.Addr().String())

	slog.Info("gRPC Server Init")
	grpcServer := grpc.NewServer(
		// protect server from shutting down grpc server by panic
		grpc.ChainUnaryInterceptor(
			grpc_recovery.UnaryServerInterceptor(),
		),
	)

	user.RegisterUserServiceServer(grpcServer, &internal.UserGrpcHandler{
		Svc: hdl.Svc,
	})
	slog.Info("gRPC Init success", "addr", listen.Addr().String())

	// Run gRPC by go Routine(Async)
	go func() {
		// executed when go routine is over
		defer func() {
			if r := recover(); r != nil {
				slog.Error("gRPC go routine panicked and recovered", "error", r)
			}
		}()

		if err := grpcServer.Serve(listen); err != nil {
			slog.Error("faild to serve gRPC : ", "Error", err)
			os.Exit(1)
		}
	}()

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

	if err := r.Run(port); err != nil {
		slog.Error("Gin Server failed to start", "error", err)
	}
}
