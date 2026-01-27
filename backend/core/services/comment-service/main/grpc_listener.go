package main

import (
	"log/slog"
	"net"
	"services/comment-service/internal"
	"services/pkg/proto/comment"

	"google.golang.org/grpc"
)

func grpc_init(hdl *internal.CommentHandler) {
	// gRPC Server Init
	slog.Info("Start Listening gRPC Server")
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		slog.Error("failed to listen: ", "Error", err)
	}
	slog.Info("Listening : ", listen.Addr().String())

	slog.Info("gRPC Server Init")
	grpcServer := grpc.NewServer()
	comment.RegisterCommetnServiceServer(grpcServer, &internal.CommentGrpcHandler{
		Svc: hdl.Service,
	})
	slog.Info("Comment gRPC Server is ready")

	if err := grpcServer.Serve(listen); err != nil {
		slog.Error("failed to serve gRPC : ", "Error", err)
	}
}
