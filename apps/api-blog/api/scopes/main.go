package scopes

import (
	"api-blog/pkg/common"
	"fmt"
	"gorm.io/gorm"
)

func Pagination[T any](db *gorm.DB, model T, query common.BaseQuery, response *common.BasePaginationResponse[T]) func(db *gorm.DB) *gorm.DB {
	var total int64

	db.Model(model).Count(&total)

	response.Total = total
	response.Page = query.GetPage()
	response.PageSize = query.GetPageSize()

	return func(db *gorm.DB) *gorm.DB {

		return db.
			Offset(query.GetOffset()).
			Limit(query.GetPageSize()).
			Order(fmt.Sprintf("%s %s", query.GetSortBy(), query.GetSort()))
	}
}
