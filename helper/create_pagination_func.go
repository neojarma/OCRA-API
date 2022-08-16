package helper

import "gorm.io/gorm"

func GetPaginationFunc(db *gorm.DB, offset, limit int) func(d *gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		return d.Offset(offset).Limit(limit)
	}
}
