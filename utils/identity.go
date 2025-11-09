package utils

import (
	"context"
	"fmt"
	"teaching_evaluate_backend/kitex_gen/teaching_evaluate"
)

const (
	UserInfoKey = "UserInfoKey"
)

func SetCurrentUserInfo(ctx context.Context, userInfo interface{}) context.Context {
	return context.WithValue(ctx, UserInfoKey, userInfo)
}

func GetUserInfoFromContext(ctx context.Context) (*teaching_evaluate.UserInfo, error) {
	val := ContextGetKeyValue(ctx, UserInfoKey)
	if user, ok := val.(teaching_evaluate.UserInfo); ok {
		return &user, nil
	}
	return nil, fmt.Errorf("user not found in context")
}
