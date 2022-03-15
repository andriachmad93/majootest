package repository

import (
	"area/models"

	"gorm.io/gorm"
)

type AreaRepository struct {
	DB *gorm.DB
}

func (_r *AreaRepository) InsertArea(param1 int64, param2 int64, area_type string, ar *models.Area) (err error) {
	var area int64
	area = 0
	switch area_type {
	case "persegi panjang":
		area = param1 * param2
		ar.AreaValue = area
		ar.AreaType = "persegi panjang"
		err = _r.DB.Create(&ar).Error
		if err != nil {
			return err
		}
	case "persegi":
		var area = param1 * param2
		ar.AreaValue = area
		ar.AreaType = "persegi"
		err = _r.DB.Create(&ar).Error
		if err != nil {
			return err
		}
	case "segitiga":

		area = 1 * (param1 * param2)
		ar.AreaValue = area
		ar.AreaType = "segitiga"
		err = _r.DB.Create(&ar).Error
		if err != nil {
			return err
		}
	default:
		ar.AreaValue = 0
		ar.AreaType = "undefined data"
		err = _r.DB.Create(&ar).Error
		if err != nil {
			return err
		}

	}

	return
}
