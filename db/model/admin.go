package db

var (
	AdminTableName = "admin"
)

type Admin struct {
	ID        int64  `gorm:"primary_key"`
	Username  string `gorm:"username"`
	Password  string `gorm:"password"`
	CreatedAt int64  `gorm:"created_at"`
	DeletedAt int64  `gorm:"deleted_at"`
}

func (s Admin) TableName() string {
	return AdminTableName
}
