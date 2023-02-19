package scopes

import (
	"api/src/common"
	"fmt"

	"gorm.io/gorm"
)

func Pagination[T any](model []*T, queries *common.BaseQuery, baseRes *common.BasePaginationResponse[T], db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var total int64
	db.Model(model).Count(&total)

	baseRes.Total = uint(total)
	baseRes.Page = queries.GetPage()
	baseRes.PageSize = queries.GetPageSize()

	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(int(queries.GetPageSize())).Offset(int(queries.GetOffset())).Order(fmt.Sprintf("%s %s", *queries.GetSortBy(), *queries.GetSort()))
	}
}
