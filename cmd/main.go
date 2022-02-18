package main

import (
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/toshkentov01/alif-tech-task/user-service/config"
	"github.com/toshkentov01/alif-tech-task/user-service/pkg/logger"
	"github.com/toshkentov01/alif-tech-task/user-service/service"
	"google.golang.org/grpc"

	pb "github.com/toshkentov01/alif-tech-task/user-service/genproto/user-service"
)

func main() {
	if info, err := os.Stat(".env"); !os.IsNotExist(err) {
		if !info.IsDir() {
			godotenv.Load(".env")
		}
	}

	var cfg = config.Get()

	serviceLog := logger.New(cfg.LogLevel, "user-service")
	defer logger.Cleanup(serviceLog)

	listen, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		serviceLog.Fatal("error while listening tcp port: ", logger.Error(err))
	}

	userService := service.NewUserService(serviceLog)
	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, userService)

	serviceLog.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := server.Serve(listen); err != nil {
		serviceLog.Fatal("error listening: %v", logger.Error(err))
	}

}
