package router

import (
	proto "github.com/oswoa/backend-service/infrastructure/grpc/proto"
	"github.com/oswoa/backend-service/usercase/service"
)

// 下位層、及びrouter(proto)層のinterfaceを実装する構造体
type Router struct {
	// protoファイルで定義したAPIを埋め込んだinterface
	proto.BackendServiceServer

	// service層にアクセスするためのinterface
	service service.IService
}

// gRPCサーバに登録するrouter層の実体を作成
func NewRouter(service *service.Service) *Router {
	router := new(Router)
	router.service = service
	return router
}
