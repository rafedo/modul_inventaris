// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package Database

const TableNameJabatan = "jabatan"

// Jabatan mapped from table <jabatan>
type Jabatan struct {
	ID   int64  `gorm:"column:id ;primaryKey" json:"id "`
	Nama string `gorm:"column:nama;not null" json:"nama"`
}

// TableName Jabatan's table name
func (*Jabatan) TableName() string {
	return TableNameJabatan
}
