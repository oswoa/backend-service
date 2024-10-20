package router

import (
	"context"

	"github.com/oswoa/backend-service/config"
	proto "github.com/oswoa/backend-service/infrastructure/grpc/proto"
	"github.com/oswoa/backend-service/infrastructure/router/validator"
	"github.com/oswoa/backend-service/infrastructure/router/validator/model"
	"github.com/oswoa/backend-service/internal_error/business_error"
)

func (i IRouter) Hello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloResponse, error) {

	// API名の設定
	ctx = context.WithValue(ctx, config.API_NAME, config.API_HELLO)

	// リクエストのバリデーション
	reqModel := model.ValidatorHello{
		Kind: request.Kind,
	}
	if err := validator.RequestValidate(reqModel); err != nil {
		business_error.PrintError(ctx.Value(config.API_NAME), err)
		return &proto.HelloResponse{}, nil
	}

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
