package dal

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	dbModel "teaching_evaluate_backend/db/model"
)

func CreateClass(ctx context.Context, db *gorm.DB, class *dbModel.Class) error {
	if err := db.WithContext(ctx).Create(class).Error; err != nil {
		klog.CtxErrorf(ctx, "create class failed: %v", err)
		return err
	}
	return nil
}

func FindClassByNumber(ctx context.Context, db *gorm.DB, number string) (*dbModel.Class, error) {
	var class dbModel.Class
	if err := db.WithContext(ctx).Where("class_number = ?", number).First(&class).Error; err != nil {
		klog.CtxErrorf(ctx, "find class by number failed: %v", err)
		return nil, err
	}
	return &class, nil
}
