package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"

	int_grpc "github.com/KentaKudo/go-todo-service/internal/grpc"
	"github.com/KentaKudo/go-todo-service/internal/pb/service"
	"github.com/KentaKudo/go-todo-service/internal/storage"
	cli "github.com/jawher/mow.cli"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var gitHash = "overridden at compile time"

const (
	appName = "go-todo-service"
	appDesc = "The gRPC Todo service example in go"
)

func main() {
	app := cli.App(appName, appDesc)

	grpcPort := app.Int(cli.IntOpt{
		Name:   "grpc-port",
		Desc:   "gRPC server port",
		EnvVar: "GRPC_PORT",
		Value:  8090,
	})

	logger := log.WithField("git_hash", gitHash)

	app.Action = func() {
		logger.Println("app started")

		inMemoryStorage := storage.NewInMemory()

		lis, err := net.Listen("tcp", net.JoinHostPort("", strconv.Itoa(*grpcPort)))
		if err != nil {
			log.Fatalln("init gRPC server:", err)
		}
		defer lis.Close()

		gSrv := initialiseGRPCServer(int_grpc.NewServer(inMemoryStorage))

		sigCh, errCh := make(chan os.Signal, 1), make(chan error, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			defer wg.Done()

			if err := gSrv.Serve(lis); err != nil {
				errCh <- fmt.Errorf("gRPC server: %w", err)
			}
		}()

		select {
		case err := <-errCh:
			logger.WithError(err).Errorln("error received. attempt graceful shutdown")
		case <-sigCh:
			logger.Println("termination signal received. attempt graceful shutdown")
		}
		gSrv.GracefulStop()
		wg.Wait()

		logger.Println("bye:)")
	}

	if err := app.Run(os.Args); err != nil {
		logger.WithError(err).Fatalln("app run")
	}
}

func initialiseGRPCServer(srv service.TodoAPIServer) *grpc.Server {
	gSrv := grpc.NewServer()

	service.RegisterTodoAPIServer(gSrv, srv)
	return gSrv
}
