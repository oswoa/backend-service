package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/oswoa/backend-service/config"
	server "github.com/oswoa/backend-service/infrastructure"
	proto "github.com/oswoa/backend-service/infrastructure/grpc/proto"
	"github.com/oswoa/backend-service/infrastructure/router"
	"github.com/oswoa/backend-service/interface/database"
	"github.com/oswoa/backend-service/usercase/service"
	"github.com/oswoa/backend-service/util"
)

func main() {

	port, err := util.GetEnv(config.PORT)
	if err != nil {
		panic(err)
	}

	// gRCPサーバの作成
	server, listener := server.NewGrpcServer(port)

	// サービスの登録
	database := database.NewDatabase()
	service := service.NewService(database)
	router := router.NewRouter(service)
	proto.RegisterBackendServiceServer(server, router)

	// gRPCサーバーを指定したポートで稼働
	go func() {
		log.Printf("gRPCサーバを起動します port: %s", port)
		server.Serve(listener)
	}()

	// Ctrl+Cが入力されたらGraceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("gRPCサーバを停止します...")
	server.GracefulStop()
}
