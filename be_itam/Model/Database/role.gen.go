// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package Database

import (
	"time"
)

const TableNameRole = "role"

// Role mapped from table <role>
type Role struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Nama      string    `gorm:"column:nama;not null" json:"nama"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName Role's table name
func (*Role) TableName() string {
	return TableNameRole
}
