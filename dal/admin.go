package dal

import (
	"context"
	"gorm.io/gorm"
	dbModel "teaching_evaluate_backend/db/model"
)

func FindAdminByAccountAndPassword(ctx context.Context, db *gorm.DB, account, password string) (*dbModel.Admin, error) {
	var admin dbModel.Admin
	err := db.WithContext(ctx).
		Where("username = ?", account).
		Where("password = ?", password).
		Where("deleted_at = 0").
		First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}
