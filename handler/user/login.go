package user

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/metadata"
	"teaching_evaluate_backend/handler"
	"teaching_evaluate_backend/kitex_gen/teaching_evaluate"
	"teaching_evaluate_backend/service/user"
	"teaching_evaluate_backend/utils"
	"time"
)

func UserLogin(ctx context.Context, req *teaching_evaluate.UserLoginRequest) (*teaching_evaluate.UserLoginResponse, error) {
	if err := user.ValidateLoginParam(req.GetUserAccount(), req.GetUserPassword()); err != nil {
		return nil, err
	}
	value := utils.ContextGetKeyValue(ctx, "headers")
	klog.CtxInfof(ctx, "ctx: %s", utils.Obj2JsonStr(ctx, value))
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if tokens, exists := md["Authorization"]; exists {
			klog.CtxInfof(ctx, "Tokens: %v", tokens)
		}
	}
	klog.CtxInfof(ctx, "value: %v", value)

	userInfo, expireTime, err := user.FindAndBuildUserInfo(ctx, req)
	if err != nil {
		return nil, err
	}

	expireAt := time.Now().Add(time.Duration(expireTime) * time.Second)
	token, err := utils.GenerateToken(expireAt, userInfo)
	if err != nil {
		return nil, fmt.Errorf("generate token failed: %w", err)
	}

	return &teaching_evaluate.UserLoginResponse{
		BaseResp: handler.ConstructSuccessResp(),
		Token:    token,
		ExpireAt: expireAt.Unix(),
		UserInfo: userInfo,
	}, nil
}
