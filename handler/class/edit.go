package class

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"teaching_evaluate_backend/dal"
	db "teaching_evaluate_backend/db/model"
	"teaching_evaluate_backend/handler"
	"teaching_evaluate_backend/kitex_gen/teaching_evaluate"
	"teaching_evaluate_backend/utils"
	"time"
)

func EditClass(ctx context.Context, req *teaching_evaluate.EditClassRequest) (*teaching_evaluate.EditClassResponse, error) {
	if len(req.GetClassNumber()) == 0 {
		return nil, errors.New("class number can not be empty")
	}

	userInfo, err := utils.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	if _, err := dal.FindClassById(ctx, db.DB, req.GetId()); err != nil {
		klog.CtxErrorf(ctx, "find class by id %v error: %v", req.GetClassNumber(), err)
		return nil, err
	}

	if err := dal.EditClass(ctx, db.DB, &db.Class{
		ID:            req.GetId(),
		ClassNumber:   req.GetClassNumber(),
		UpdatedAt:     time.Now().Unix(),
		UpdatedOpID:   userInfo.Id,
		UpdatedOpName: userInfo.Name,
	}); err != nil {
		return nil, fmt.Errorf("edit class error: %v", err)
	}

	return &teaching_evaluate.EditClassResponse{
		BaseResp: handler.ConstructSuccessResp(),
	}, nil
}
