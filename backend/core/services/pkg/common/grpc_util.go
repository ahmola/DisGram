package common

import (
	"log/slog"
	"net"

	"google.golang.org/grpc"
)

func StartGrpcServer(port string, serviceName string) (net.Listener, *grpc.Server) {
	// gRPC Server Init
	slog.Info("Start Listening ", serviceName, " gRPC Server")
	listen, err := net.Listen("tcp", port)
	if err != nil {
		slog.Error("failed to open tcp ", port, "Error", err)
	}
	slog.Info("Listening: ", listen.Addr().String())

	slog.Info("gRPC Server Init")
	grpcServer := grpc.NewServer()

	return listen, grpcServer
}
