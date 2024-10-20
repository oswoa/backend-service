package router

import (
	"context"

	proto "github.com/oswoa/backend-service/grpc/proto"
)

func (i IRouter) Hello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloResponse, error) {

	var msg string

	switch request.Kind {
	case 1:
		msg = "Hello, World"

	case 2:
		msg = "Hello, gRPC"

	default:
		msg = "none"
	}

	return &proto.HelloResponse{
		Msg: msg,
	}, nil
}
