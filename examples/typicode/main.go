package main

import (
	"net"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	userapi "github.com/Kuwerin/protoc-gen-go-httpclient/examples/typicode/gen/go/user"
	"github.com/Kuwerin/protoc-gen-go-httpclient/pkg/transport"
)

func main() {
	// Create the logger.
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "svc", "erp-client")
		logger.Log("app", os.Args[0], "event", "starting")
	}

	baseURL, _ := url.Parse("https://jsonplaceholder.typicode.com/")

	// Create user client.
	var userClient userapi.UserServiceServer
	{
		userClient = userapi.NewHTTPClient(&userapi.HTTPClientParams{
			URL:                   baseURL,
			AllowUndeclaredFields: true,
		})
		userClient = userapi.LoggingMiddleware(logger)(userClient)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	ln, err := net.Listen("tcp", ":5000")
	if err != nil {
		logger.Log("error", err)
		os.Exit(1)
	}

	// Register user server.
	userServer := &transport.Server[userapi.UserServiceServer]{
		Server:             grpcServer,
		ContextAPI:         &userClient,
		RegisterServerFunc: userapi.RegisterUserServiceServer,
	}
	if err := userServer.Register(); err != nil {
		logger.Log("entity", "transport.grpc.user", "error", err)
		os.Exit(1)
	}
	logger.Log("entity", "transport.grpc.user", "event", "registred")

	// Start listening gRPC server.
	go func() {
		if err := grpcServer.Serve(ln); err != nil {
			logger.Log("error", err)
			os.Exit(1)
		}
	}()
	logger.Log("event", "started listening")

	// Wait for a signal for graceful shutdown.
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	s := <-quit

	logger.Log("event", "received os signal to shutdown application", "signal", s)

	grpcServer.GracefulStop()

	logger.Log("event", "gRPC server stopped gracefully")

	logger.Log("event", "application stopped gracefully")

}
