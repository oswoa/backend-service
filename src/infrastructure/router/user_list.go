package router

import (
	"context"

	"github.com/oswoa/backend-service/config"
	proto "github.com/oswoa/backend-service/infrastructure/grpc/proto"
	"github.com/oswoa/backend-service/infrastructure/router/validator/model"
)

func (i Router) UserList(ctx context.Context, request *proto.UserListRequest) (*proto.UserListResponse, error) {
	// API名の設定
	ctx = context.WithValue(ctx, config.API_NAME, config.API_USER_LIST)

	// リクエストのバリデーション
	reqModel := model.ValidatorUserList{
		Email:     request.Email,
		IsDeleted: request.IsDeleted,
	}
	/*
		if err := validator.RequestValidate(reqModel); err != nil {
			business_error.PrintError(ctx.Value(config.API_NAME), err)
			return &proto.UserListResponse{}, nil
		}
	*/

	// usecase呼び出し
	userList := i.service.UserList(reqModel.Email, reqModel.IsDeleted)

	response := make([]*proto.UserDetail, 0)
	for _, v := range userList {
		val := &proto.UserDetail{
			UserId:      v.UserId,
			Email:       v.Email,
			IsAvailable: v.ConvertIsAvailable(),
		}
		response = append(response, val)
	}

	return &proto.UserListResponse{
		UserDetail: response,
	}, nil
}
