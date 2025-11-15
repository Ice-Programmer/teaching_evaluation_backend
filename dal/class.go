package dal

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	dbModel "teaching_evaluate_backend/db/model"
	eva "teaching_evaluate_backend/kitex_gen/teaching_evaluate"
	"teaching_evaluate_backend/utils"
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
	if err := db.WithContext(ctx).Where("class_number = ?", number).Where("deleted_at = 0").First(&class).Error; err != nil {
		klog.CtxErrorf(ctx, "find class by number failed: %v", err)
		return nil, err
	}
	return &class, nil
}

func FindClassById(ctx context.Context, db *gorm.DB, id int64) (*dbModel.Class, error) {
	var class dbModel.Class
	if err := db.WithContext(ctx).Where("id = ?", id).Where("deleted_at = 0").First(&class).Error; err != nil {
		klog.CtxErrorf(ctx, "find class by id failed: %v", err)
		return nil, err
	}
	return &class, nil
}

func EditClass(ctx context.Context, db *gorm.DB, class *dbModel.Class) error {
	if err := db.WithContext(ctx).Model(class).Updates(class).Error; err != nil {
		klog.CtxErrorf(ctx, "edit class failed: %v", err)
		return err
	}
	return nil
}

func QueryClass(ctx context.Context, db *gorm.DB, pageNum, pageSize int32, condition *eva.QueryClassCondition) ([]*dbModel.Class, int64, error) {
	query := buildCondition(db, condition).WithContext(ctx)

	var total int64
	if err := query.Count(&total).Error; err != nil {
		klog.CtxErrorf(ctx, "query class count failed: %v", err)
		return nil, 0, err
	}

	var classList []*dbModel.Class
	offset := int((pageNum - 1) * pageSize)
	err := query.
		Limit(int(pageSize)).Offset(offset).
		Order("created_at DESC").
		Find(&classList).Error
	if err != nil {
		klog.CtxErrorf(ctx, "query class list failed: %v", err)
		return nil, 0, err
	}

	return classList, total, nil
}

func buildCondition(db *gorm.DB, condition *eva.QueryClassCondition) *gorm.DB {
	db = db.Table(dbModel.ClassTableName).Where("deleted_at = 0")
	if condition == nil {
		return db
	}

	if condition.Id != nil {
		db = db.Where("id = ?", condition.Id)
	}

	if condition.SearchText != nil {
		db = db.Where("class_number like ?", utils.WrapLike(condition.GetSearchText()))
	}

	if len(condition.GetIds()) > 0 {
		db = db.Where("id IN (?)", condition.Ids)
	}

	return db
}
