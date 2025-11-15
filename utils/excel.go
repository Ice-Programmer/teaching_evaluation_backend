package utils

import (
	"bytes"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

// ExcelRowMapper 用于将 Excel 行映射成结构体或处理函数
type ExcelRowMapper[T any] func(row []string, rowIndex int) (T, error)

// GenerateExampleExcel 生成通用示例 Excel 文件。
// 参数说明：
//   - sheetName: 工作表名
//   - headers:   表头（按列顺序）
//   - exampleRow: 示例数据（按列顺序）
//
// 返回：Excel 二进制字节数组、错误
func GenerateExampleExcel(sheetName string, headers []string, exampleRow []string) ([]byte, error) {
	if len(headers) == 0 {
		return nil, fmt.Errorf("headers cannot be empty")
	}

	// 创建 Excel 文件
	excelFile := excelize.NewFile()
	defer func() { _ = excelFile.Close() }()

	// 创建工作表
	if _, err := excelFile.NewSheet(sheetName); err != nil {
		return nil, fmt.Errorf("create sheet: %w", err)
	}

	// 写入表头
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		if err := excelFile.SetCellValue(sheetName, cell, h); err != nil {
			return nil, fmt.Errorf("set header value: %w", err)
		}
	}

	// 写入示例行（第二行）
	for i, val := range exampleRow {
		cell, _ := excelize.CoordinatesToCellName(i+1, 2)
		if err := excelFile.SetCellValue(sheetName, cell, val); err != nil {
			return nil, fmt.Errorf("set example value: %w", err)
		}
	}

	// 第三行加提示
	tipCell := getLastColumnCell(len(exampleRow)+1, 2)
	_ = excelFile.SetCellValue(sheetName, tipCell, "⚠️ 示例数据，请勿删除，请紧跟在第二行输入")

	// 设置样式
	headerStyle, _ := excelFile.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true, Color: "#FFFFFF"},
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"#4F81BD"}, Pattern: 1},
		Alignment: &excelize.Alignment{Horizontal: "center"},
	})
	_ = excelFile.SetCellStyle(sheetName, "A1",
		getLastColumnCell(len(headers), 1), headerStyle)

	// 设置列宽
	_ = excelFile.SetColWidth(sheetName, "A", getColumnLetter(len(headers)), 20)

	// 设为当前 Sheet
	index, err := excelFile.GetSheetIndex(sheetName)
	if err != nil {
		return nil, fmt.Errorf("get sheet index: %w", err)
	}
	excelFile.SetActiveSheet(index)

	// 输出为字节
	buf, err := excelFile.WriteToBuffer()
	if err != nil {
		return nil, fmt.Errorf("write excel buffer: %w", err)
	}
	return buf.Bytes(), nil
}

// getColumnLetter 将列号（1,2,3...）转成列字母（A,B,C...）
func getColumnLetter(col int) string {
	letter, _ := excelize.ColumnNumberToName(col)
	return letter
}

// getLastColumnCell 返回某行的最后一个单元格位置，如 C1、E2
func getLastColumnCell(col, row int) string {
	name, _ := excelize.CoordinatesToCellName(col, row)
	return name
}

// ParseExcelWithMapper 泛型 + 回调解析 Excel
func ParseExcelWithMapper[T any](ctx context.Context, excelBytes []byte, startRow int, mapper ExcelRowMapper[T]) ([]T, error) {
	if len(excelBytes) == 0 {
		return nil, fmt.Errorf("excel file is empty")
	}

	excel, err := excelize.OpenReader(bytes.NewReader(excelBytes))
	if err != nil {
		hlog.CtxErrorf(ctx, "ParseExcelWithMapper read excel error: %s", err.Error())
		return nil, fmt.Errorf("failed to open excel file: %w", err)
	}

	sheetName := excel.GetSheetName(0)
	if sheetName == "" {
		return nil, fmt.Errorf("excel sheet is empty")
	}

	rows, err := excel.GetRows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("failed to get rows: %w", err)
	}

	var result []T
	for i, row := range rows {
		if startRow > i+1 {
			continue
		}
		if len(row) == 0 {
			continue
		}

		item, err := mapper(row, i)
		if err != nil {
			return nil, fmt.Errorf("parse row %d err: %w", i+1, err)
		}
		result = append(result, item)
	}

	return result, nil
}

// StrToInt32 辅助函数：字符串转 int32
func StrToInt32(str string, defaultVal int32) int32 {
	val, err := strconv.Atoi(strings.TrimSpace(str))
	if err != nil {
		return defaultVal
	}
	return int32(val)
}
