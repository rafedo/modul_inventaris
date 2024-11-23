// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package Database

import (
	"time"
)

const TableNameDetaiAsetAplikasi = "detai_aset_aplikasi"

// DetaiAsetAplikasi mapped from table <detai_aset_aplikasi>
type DetaiAsetAplikasi struct {
	ID                      int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	NamaAplikasi            string    `gorm:"column:nama aplikasi;not null" json:"nama aplikasi"`
	TanggalPembuatan        time.Time `gorm:"column:tanggal_pembuatan;not null" json:"tanggal_pembuatan"`
	TanggalTerima           time.Time `gorm:"column:tanggal_terima;not null" json:"tanggal_terima"`
	LokasiServerPenyimpanan string    `gorm:"column:lokasi_server_penyimpanan;not null" json:"lokasi_server_penyimpanan"`
	TipeAplikasi            string    `gorm:"column:tipe_aplikasi;not null" json:"tipe_aplikasi"`
	LinkAplikasi            string    `gorm:"column:link_aplikasi;not null" json:"link_aplikasi"`
	SertifikasiAplikasi     string    `gorm:"column:sertifikasi_aplikasi" json:"sertifikasi_aplikasi"`
	TanggalAktif            time.Time `gorm:"column:tanggal_aktif;not null" json:"tanggal_aktif"`
	TanggalKadaluarsa       time.Time `gorm:"column:tanggal_kadaluarsa;not null" json:"tanggal_kadaluarsa"`
	AssetID                 int64     `gorm:"column:asset_id;not null" json:"asset_id"`
	CreatedAt               time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt               time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName DetaiAsetAplikasi's table name
func (*DetaiAsetAplikasi) TableName() string {
	return TableNameDetaiAsetAplikasi
}
