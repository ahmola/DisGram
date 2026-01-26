package main

import (
	"log/slog"
	"net"
	"services/follow-service/internal"
	"services/pkg/proto/follow"

	"google.golang.org/grpc"
)

func grpc_init(hdl *internal.FollowHandler) {
	// gRPC Server Init
	slog.Info("Start Listening gRPC Server")
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		slog.Error("failed to open tcp port 9090 ", "Error", err)
	}
	slog.Info("Listening: ", listen.Addr().String())

	slog.Info("gRPC Server Init")
	grpcServer := grpc.NewServer()
	follow.RegisterFollowServiceServer(grpcServer, &internal.FollowGrpcHandler{
		Svc: hdl.Svc,
	})
	slog.Info("Follow gRPC Server is ready")

	if err := grpcServer.Serve(listen); err != nil {
		slog.Error("faild to server gRPC: ", "Error", err)
	}
}
