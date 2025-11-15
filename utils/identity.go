package utils

import (
	"context"
	"fmt"
	"teaching_evaluate_backend/consts"
	"teaching_evaluate_backend/kitex_gen/teaching_evaluate"
)

func GetUserInfo(ctx context.Context) (*teaching_evaluate.UserInfo, error) {
	value := ContextGetKeyValue(ctx, consts.AuthorizationHeader)
	if value == nil {
		return nil, fmt.Errorf("user not login")
	}

	parseToken, err := ParseToken(value.(string))
	if err != nil || parseToken == nil {
		return nil, fmt.Errorf("parse token error: %v", err)
	}

	return &teaching_evaluate.UserInfo{
		Id:       parseToken.ID,
		Name:     parseToken.Username,
		Role:     parseToken.Role,
		CreateAt: parseToken.CreateAt,
	}, nil
}
