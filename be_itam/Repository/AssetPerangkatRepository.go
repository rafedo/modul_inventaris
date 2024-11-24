package Repository

import (
	"itam/Model/Database"

	"gorm.io/gorm"
)

type (
	AssetPerangkatRepositoryHandler interface {
		Save(data *Database.DetailAsetPerangkat) (id int64, err error)
		Update(data *Database.DetailAsetPerangkat) (id int64, err error)
		Delete(id int64) error
		FindById(id int64) (data Database.DetailAsetPerangkat, err error)
		FindAll() (data []Database.DetailAsetPerangkat, err error)
	}

	AssetPerangkatRepositoryImpl struct {
		DB *gorm.DB
	}
)

func AssetPerangkatRepositoryProvider(db *gorm.DB) *AssetPerangkatRepositoryImpl {
	return &AssetPerangkatRepositoryImpl{
		DB: db,
	}
}

func (h *AssetPerangkatRepositoryImpl) Save(data *Database.DetailAsetPerangkat) (id int64, err error) {
	err = h.DB.Model(&Database.DetailAsetPerangkat{}).
		Save(&data).Error

	return data.ID, err
}

func (h *AssetPerangkatRepositoryImpl) Update(data *Database.DetailAsetPerangkat) (id int64, err error) {
	err = h.DB.Model(&Database.DetailAsetPerangkat{}).
		Where("id = ?", data.ID).
		Updates(&data).Error

	return data.ID, err
}

func (h *AssetPerangkatRepositoryImpl) Delete(id int64) error {
	err := h.DB.Delete(&Database.DetailAsetPerangkat{}, id).Error
	return err
}

func (h *AssetPerangkatRepositoryImpl) FindById(id int64) (data Database.DetailAsetPerangkat, err error) {
	err = h.DB.Model(&Database.DetailAsetPerangkat{}).
		Where("id = ?", id).
		Take(&data).
		Error
	return data, err
}

func (h *AssetPerangkatRepositoryImpl) FindAll() (data []Database.DetailAsetPerangkat, err error) {
	err = h.DB.Model(&Database.DetailAsetPerangkat{}).
		Order("id asc").
		Find(&data).
		Error

	return data, err
}
