package model

type Role struct {
	RoleID   int    `gorm:"column:role_id;primaryKey"`
	RoleName string `gorm:"column:role_name"`
	User     []User `gorm:"foreignKey:role_id_fk;references:role_id"`
}
