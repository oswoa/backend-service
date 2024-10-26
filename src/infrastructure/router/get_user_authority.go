package router

import (
	"context"

	"github.com/oswoa/backend-service/config"
	proto "github.com/oswoa/backend-service/infrastructure/grpc/proto"
	"github.com/oswoa/backend-service/infrastructure/router/validator"
	"github.com/oswoa/backend-service/infrastructure/router/validator/model"
	"github.com/oswoa/backend-service/internal_error/business_error"
)

// ユーザ権限取得
func (i Router) GetUserAuthority(ctx context.Context, request *proto.GetUserAuthorityRequest) (*proto.GetUserAuthorityResponse, error) {
	// API名の設定
	ctx = context.WithValue(ctx, config.API_NAME, config.API_GET_USER_AUTHORITY)

	// リクエストのバリデーション
	reqModel := model.ValidatorGetUserAuthorityRequest{
		UserId: request.UserId,
	}
	if err := validator.Validate(reqModel, validator.VALIDATE_REQUEST); err != nil {
		business_error.PrintError(ctx, err)
		return &proto.GetUserAuthorityResponse{}, nil
	}

	// usecase呼び出し
	getUserAuthority := i.service.GetUserAuthority(reqModel.UserId)

	// レスポンスのバリデーション
	resModel := model.ValidatorGetUserAuthorityResponse{
		AuthorityName: getUserAuthority.AuthorityName,
		CanSearch:     getUserAuthority.CanSearch,
		CanCreate:     getUserAuthority.CanCreate,
		CanUpdate:     getUserAuthority.CanUpdate,
		CanDelete:     getUserAuthority.CanDelete,
		CanApprove:    getUserAuthority.CanApprove,
		CanPullBack:   getUserAuthority.CanPullBack,
	}
	if err := validator.Validate(resModel, validator.VALIDATE_RESPONSE); err != nil {
		business_error.PrintError(ctx, err)
		return &proto.GetUserAuthorityResponse{}, nil
	}

	return &proto.GetUserAuthorityResponse{
		AuthorityName: resModel.AuthorityName,
		CanSearch:     resModel.CanSearch,
		CanCreate:     resModel.CanCreate,
		CanUpdate:     resModel.CanUpdate,
		CanDelete:     resModel.CanDelete,
		CanApprove:    resModel.CanApprove,
		CanPullBack:   resModel.CanPullBack,
	}, nil
}
