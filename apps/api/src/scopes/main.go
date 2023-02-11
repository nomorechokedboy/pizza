package scopes

import (
	"api/src/common"
	"fmt"

	"gorm.io/gorm"
)

func Pagination(queries *common.BaseQuery) func(db *gorm.DB) *gorm.DB {
	fmt.Println("Debug pagination query: ", queries, queries.GetOffset(), queries.GetPageSize())
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(int(queries.GetPageSize())).Offset(int(queries.GetOffset()))
	}
}
