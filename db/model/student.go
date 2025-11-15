package db

var (
	StudentTableName = "student"
)

type Student struct {
	ID            int64  `gorm:"primary_key;column:id"`
	StudentNumber string `gorm:"column:student_number"`
	StudentName   string `gorm:"column:student_name"`
	ClassID       int64  `gorm:"column:class_id"`
	Major         int8   `gorm:"column:major"` // 0 - 计算机, 1 - 自动化
	Grade         int    `gorm:"column:grade"`
	Status        int8   `gorm:"column:status"` // 0 - 正常使用, 1 - 拒绝访问
	CreateAt      int64  `gorm:"column:create_at"`
	DeletedAt     int8   `gorm:"column:deleted_at"`
}

func (s Student) TableName() string {
	return StudentTableName
}
