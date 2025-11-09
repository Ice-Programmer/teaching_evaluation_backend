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
)

func CreateClass(ctx context.Context, req *teaching_evaluate.CreateClassRequest) (*teaching_evaluate.CreateClassResponse, error) {
	if err := CheckClassParam(req.GetClassNumber()); err != nil {
		klog.CtxErrorf(ctx, "class param validation error: %v", err)
		return nil, err
	}

	dal.CreateClass(ctx, db.DB, &db.Class{
		ID:          utils.GetId(),
		ClassNumber: req.GetClassNumber(),
	})

	return &teaching_evaluate.CreateClassResponse{
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}

func CheckClassParam(classNumber string) error {
	if classNumber == "" {
		return fmt.Errorf("class number is required")
	}

	return nil
}
