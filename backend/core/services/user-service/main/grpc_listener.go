package main

import (
	"log/slog"
	"net"
	"services/pkg/proto/user"
	"services/user-service/internal"

	"google.golang.org/grpc"
)

func grpc_init(hdl *internal.UserHandler) {
	// gRPC Server Init
	slog.Info("Start Listening gRPC Server")
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		slog.Error("Failed to listen: ", "error", err)
	}
	slog.Info("Listening : ", listen.Addr().String())

	slog.Info("gRPC Server Init")
	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, &internal.UserGrpcHandler{
		Svc: hdl.Svc,
	})
	slog.Info("gRPC Server is ready")

	if err := grpcServer.Serve(listen); err != nil {
		slog.Error("failed to serve gRPC: ", "error", err)
	}
}
