package scopes

import (
	"api/src/common"

	"gorm.io/gorm"
)

func Pagination(queries *common.BaseQuery) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(int(queries.GetPageSize())).Offset(int(queries.GetOffset()))
	}
}
