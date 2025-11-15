package db

var (
	ClassTableName = "student_class"
)

type Class struct {
	ID            int64  `gorm:"primary_key"`
	ClassNumber   string `gorm:"class_number"`
	CreatedAt     int64  `gorm:"create_at"`
	CreatedOpName string `gorm:"create_op_name"`
	CreatedOpID   int64  `gorm:"create_op_id"`
	UpdatedAt     int64  `gorm:"update_at"`
	UpdatedOpName string `gorm:"update_op_name"`
	UpdatedOpID   int64  `gorm:"update_op_id"`
	deletedAt     int64  `gorm:"deleted_at"`
}

func (s Class) TableName() string {
	return ClassTableName
}
