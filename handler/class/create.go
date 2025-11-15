package class

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"teaching_evaluate_backend/dal"
	db "teaching_evaluate_backend/db/model"
	"teaching_evaluate_backend/handler"
	"teaching_evaluate_backend/kitex_gen/teaching_evaluate"
	"teaching_evaluate_backend/utils"
	"time"
)

func CreateClass(ctx context.Context, req *teaching_evaluate.CreateClassRequest) (*teaching_evaluate.CreateClassResponse, error) {
	if err := CheckClassParam(req.GetClassNumber()); err != nil {
		klog.CtxErrorf(ctx, "class param validation error: %v", err)
		return nil, err
	}

	userInfo, err := utils.GetUserInfo(ctx)
	if err != nil {
		klog.CtxErrorf(ctx, "get user info error: %v", err)
		return nil, err
	}

	if err := dal.CreateClass(ctx, db.DB, &db.Class{
		ID:          utils.GetId(),
		ClassNumber: req.GetClassNumber(),
		CreatedAt:   time.Now().Unix(),
		CreatedOpID: userInfo.Id,
	}); err != nil {
		klog.CtxErrorf(ctx, "create class error: %v", err)
		return nil, err
	}

	return &teaching_evaluate.CreateClassResponse{
		Id:       utils.GetId(),
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}

func CheckClassParam(classNumber string) error {
	if classNumber == "" {
		return fmt.Errorf("class number is required")
	}

	return nil
}
