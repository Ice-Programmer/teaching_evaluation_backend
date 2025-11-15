package class

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"teaching_evaluate_backend/handler"
	"teaching_evaluate_backend/kitex_gen/teaching_evaluate"
	"teaching_evaluate_backend/utils"
)

func GetClassImportExcel(ctx context.Context, req *teaching_evaluate.GetClassImportExcelRequest) (*teaching_evaluate.GetClassImportExcelResponse, error) {
	headers := []string{"班级编号"}
	exampleRow := []string{"23320110"}

	fileBytes, err := utils.GenerateExampleExcel("班级导入数据", headers, exampleRow)
	if err != nil {
		hlog.CtxErrorf(ctx, "GetStudentClassExcelExample GenerateExampleExcel error: %s", err.Error())
		return nil, fmt.Errorf("generate example excel error: %s", err.Error())
	}

	return &teaching_evaluate.GetClassImportExcelResponse{
		ExcelFile: fileBytes,
		FileName:  "学生导入表格.xlsx",
		BaseResp:  handler.ConstructSuccessResp(),
	}, nil
}
