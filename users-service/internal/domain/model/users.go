package model

import "time"

type User struct {
	UserID    int    `gorm:"column:user_id;primaryKey;autoIncrement"`
	Password  string `gorm:"column:password"`
	Email     string `gorm:"column:email"`
	FullName  string `gorm:"column:full_name"`
	PhoneNum  string `gorm:"column:phone_num"`
	Address   string `gorm:"column:address"`
	RoleIDFK  int    `gorm:"column:role_id_fk"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
