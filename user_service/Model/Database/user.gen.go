// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package Database

import (
	"time"
)

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID               int64     `gorm:"column:id;primaryKey" json:"id"`
	NIP              int64     `gorm:"column:NIP;not null" json:"NIP"`
	Email            string    `gorm:"column:email;not null" json:"email"`
	JabatanID        int64     `gorm:"column:jabatan_id;not null" json:"jabatan_id"`
	DivisiID         int64     `gorm:"column:divisi_id;not null" json:"divisi_id"`
	TanggalBergabung time.Time `gorm:"column:tanggal_bergabung;not null" json:"tanggal_bergabung"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
