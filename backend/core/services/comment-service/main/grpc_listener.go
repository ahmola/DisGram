package main

import (
	"log"
	"log/slog"
	"net"
	"services/comment-service/internal"
	"services/pkg/proto/comment"

	"google.golang.org/grpc"
)

func grpc_init(hdl *internal.CommentHandler) {
	// gRPC Server Init
	slog.Info("gRPC Server Init")
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	comment.RegisterCommetnServiceServer(grpcServer, &internal.CommentGrpcHandler{
		Svc: hdl.Service,
	})

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRPC : %v", err)
	}
	slog.Info("gRPC Server is ready")
}
