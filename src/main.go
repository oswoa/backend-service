package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/oswoa/backend-service/config"
	proto "github.com/oswoa/backend-service/grpc/proto"
	"github.com/oswoa/backend-service/infrastructure/router"
	"github.com/oswoa/backend-service/util"
)

func main() {

	port, err := util.GetEnv(config.PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	// gRPCサーバーを指定したポートで稼働させる
	server, listener := router.NewGrpcServer(port)
	go func() {
		log.Printf("gRPCサーバを起動します port: %s", port)
		server.Serve(listener)
	}()

	// Routerの登録
	proto.RegisterBackendServiceServer(server, router.NewRouter())

	// Ctrl+Cが入力されたらGraceful shutdownされるようにする
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("gRPCサーバを停止します...")
	server.GracefulStop()
}
