package common

import (
	"log/slog"
	"net"
	"os"

	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
)

/*
* Deprecated. No longer be used
 */

func StartGrpcServer(port string, serviceName string) (net.Listener, *grpc.Server) {
	// gRPC Server Init
	slog.Info("Start Listening ", serviceName, " gRPC Server")
	listen, err := net.Listen("tcp", port)
	if err != nil {
		slog.Error("failed to open tcp ", port, "Error", err)
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

	return listen, grpcServer
}

func RunGrpcWithGoRoutine(listen net.Listener, grpcServer *grpc.Server) {
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
}
