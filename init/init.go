package init

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"teaching_evaluate_backend/consts"
	"teaching_evaluate_backend/db/model"
	"teaching_evaluate_backend/utils"
)

func Init(ctx context.Context) error {
	if err := InitDBGorm(); err != nil {
		klog.CtxErrorf(ctx, "InitDBGorm err: %v", err)
		panic(err)
	}
	klog.CtxInfof(ctx, "InitDBGorm success")

	if err := utils.InitIdGeneratorClient(); err != nil {
		klog.CtxErrorf(ctx, "InitIdGeneratorClient err: %v", err)
		panic(err)
	}

	return nil
}

func InitDBGorm() error {
	gormDB, err := newInit()
	if err != nil {
		return err
	}
	db.DB = gormDB
	return nil
}

func newInit() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		consts.DBUser, consts.DBPassword,
		consts.DBHost, consts.DBPort, consts.DBName)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
