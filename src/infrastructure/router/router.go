package router

import (
	proto "github.com/oswoa/backend-service/grpc/proto"
)

// バックエンドで扱うAPIを埋め込んだ構造体
type IRouter struct {
	proto.BackendServiceServer
}

// gRPCサーバに登録するrouter情報を返す
func NewRouter() *IRouter {
	return &IRouter{}
}
