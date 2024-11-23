// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package Database

const TableNameAsset = "asset"

// Asset mapped from table <asset>
type Asset struct {
	ID           int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	VendorID     int64  `gorm:"column:vendor_id;not null" json:"vendor_id"`
	SerialNumber string `gorm:"column:serial_number;not null" json:"serial_number"`
	Merk         string `gorm:"column:merk;not null" json:"merk"`
	Model        string `gorm:"column:model;not null" json:"model"`
	NomorNota    string `gorm:"column:nomor_nota;not null" json:"nomor_nota"`
}

// TableName Asset's table name
func (*Asset) TableName() string {
	return TableNameAsset
}
