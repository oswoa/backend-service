package router

import (
	"context"

	"github.com/oswoa/backend-service/config"
	proto "github.com/oswoa/backend-service/infrastructure/grpc/proto"
	"github.com/oswoa/backend-service/infrastructure/router/validator"
	"github.com/oswoa/backend-service/infrastructure/router/validator/model"
	"github.com/oswoa/backend-service/internal_error/business_error"
)

// ユーザ一覧照会
func (i Router) UserList(ctx context.Context, request *proto.UserListRequest) (*proto.UserListResponse, error) {
	// API名の設定
	ctx = context.WithValue(ctx, config.API_NAME, config.API_USER_LIST)

	// リクエストのバリデーション
	reqModel := model.ValidatorUserListRequest{
		Email:     request.Email,
		IsDeleted: request.IsDeleted,
	}
	if validateErr := validator.Validate(reqModel); validateErr != nil {
		err := business_error.RequestValidateError{
			ErrCode:   business_error.ERR_CODE_VALIDATE_REQUEST,
			ErrTag:    validateErr.ErrTag,
			FieldName: validateErr.FieldName,
		}
		business_error.PrintError(ctx, err)
		return &proto.UserListResponse{}, nil
	}

	// usecase呼び出し
	userList := i.service.UserList(reqModel.Email, reqModel.IsDeleted)

	// レスポンスのバリデーション
	for _, v := range userList {
		resModel := model.ValidatorUserListResponse{
			UserId:      v.UserId,
			Email:       v.Email,
			IsAvailable: v.IsAvailable,
		}
		if validateErr := validator.Validate(resModel); validateErr != nil {
			err := business_error.ResponseValidateError{
				ErrCode:   business_error.ERR_CODE_VALIDATE_RESPONSE,
				ErrTag:    validateErr.ErrTag,
				FieldName: validateErr.FieldName,
			}
			business_error.PrintError(ctx, err)
			return &proto.UserListResponse{}, nil
		}
	}

	response := make([]*proto.UserDetail, 0)
	for _, v := range userList {
		val := &proto.UserDetail{
			UserId:      v.UserId,
			Email:       v.Email,
			IsAvailable: v.IsAvailable,
		}
		response = append(response, val)
	}

	return &proto.UserListResponse{
		UserDetail: response,
	}, nil
}
