package router

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
)

// 新規gRPCサーバの作成
func NewGrpcServer(port string) (*grpc.Server, net.Listener) {
	// 指定されたポートでLisnterを作成
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}

	// gRPCサーバーを作成
	server := grpc.NewServer()

	return server, listener
}
