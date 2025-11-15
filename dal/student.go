package dal

import (
	"context"
	"gorm.io/gorm"
	dbModel "teaching_evaluate_backend/db/model"
)

func FindStudentByStudentNumber(ctx context.Context, db *gorm.DB, studentNumber string) (*dbModel.Student, error) {
	var student dbModel.Student
	err := db.WithContext(ctx).
		Where("student_number = ?", studentNumber).
		Where("deleted_at = 0").
		First(&student).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}
