package class

import (
	"context"
	"teaching_evaluate_backend/dal"
	db "teaching_evaluate_backend/db/model"
	"teaching_evaluate_backend/handler"
	"teaching_evaluate_backend/kitex_gen/teaching_evaluate"
	"teaching_evaluate_backend/utils"
)

func QueryClass(ctx context.Context, req *teaching_evaluate.QueryClassRequest) (*teaching_evaluate.QueryClassResponse, error) {
	pageSize, pageNum := utils.SetPageDefault(req.PageSize, req.PageNum)

	classList, total, err := dal.QueryClass(ctx, db.DB, pageNum, pageSize, req.Condition)
	if err != nil {
		return nil, err
	}

	return &teaching_evaluate.QueryClassResponse{
		Total:         total,
		ClassInfoList: wrapClassInfoList(classList),
		BaseResp:      handler.ConstructSuccessResp(),
	}, nil
}

func wrapClassInfoList(classList []*db.Class) []*teaching_evaluate.ClassInfo {
	classInfoList := make([]*teaching_evaluate.ClassInfo, 0, len(classList))
	for _, class := range classList {
		classInfoList = append(classInfoList, &teaching_evaluate.ClassInfo{
			Id:           class.ID,
			ClassNumber:  class.ClassNumber,
			CreateTime:   class.CreatedAt,
			CreateOpId:   class.CreatedOpID,
			CreateOpName: class.CreatedOpName,
			UpdateTime:   class.UpdatedAt,
			UpdateOpId:   class.UpdatedOpID,
			UpdateOpName: class.UpdatedOpName,
		})
	}

	return classInfoList
}
