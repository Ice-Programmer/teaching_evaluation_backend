package user

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"teaching_evaluate_backend/dal"
	db "teaching_evaluate_backend/db/model"
	"teaching_evaluate_backend/kitex_gen/teaching_evaluate"
	"teaching_evaluate_backend/utils"
	"time"
)

const (
	AdminLoginExpireTime   int64 = 3600 * 24 * 30
	StudentLoginExpireTime int64 = 3600 * 24
)

// ValidateLoginParam checks if required fields are non-empty.
func ValidateLoginParam(userAccount, userPassword string) error {
	if userAccount == "" {
		return fmt.Errorf("userAccount is empty")
	}
	if userPassword == "" {
		return fmt.Errorf("userPassword is empty")
	}
	return nil
}

// FindAndBuildUserInfo checks whether the user is a student or admin,
// returns corresponding userInfo and expireTime.
func FindAndBuildUserInfo(ctx context.Context, req *teaching_evaluate.UserLoginRequest) (*teaching_evaluate.UserInfo, int64, error) {
	userAccount := req.GetUserAccount()
	hashedPwd := utils.MD5(req.GetUserPassword())

	// Try student first
	student, err := dal.FindStudentByStudentNumber(ctx, db.DB, userAccount)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, err
	}

	if student != nil {
		return BuildUserInfo(student.ID, student.StudentName, teaching_evaluate.UserRole_Student), StudentLoginExpireTime, nil
	}

	// Try admin
	admin, err := dal.FindAdminByAccountAndPassword(ctx, db.DB, userAccount, hashedPwd)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, fmt.Errorf("user not found")
		}
		return nil, 0, err
	}

	return BuildUserInfo(admin.ID, admin.Username, teaching_evaluate.UserRole_Admin), AdminLoginExpireTime, nil
}

// BuildUserInfo constructs a UserInfo object.
func BuildUserInfo(id int64, name string, role teaching_evaluate.UserRole) *teaching_evaluate.UserInfo {
	return &teaching_evaluate.UserInfo{
		Id:       id,
		Name:     name,
		Role:     role,
		CreateAt: time.Now().Unix(),
	}
}
