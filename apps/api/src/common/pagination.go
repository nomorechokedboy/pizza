package common

import (
	"api/src/utils"
)

type BaseQuery struct {
	Page     *uint   `query:"page"`
	PageSize *uint   `query:"pageSize"`
	Q        *string `query:"q"`
	Sort     *string `query:"sort"`
	SortBy   *string `query:"sortBy"`
}

type BasePaginationResponse[T any] struct {
	Items    *[]T  `js:"items"`
	Page     *uint `js:"page"`
	PageSize *uint `js:"pageSize"`
	Total    uint  `js:"total"`
}

func (baseQ *BaseQuery) GetPage() uint {
	if baseQ.Page == nil {
		baseQ.Page = utils.GetDataTypeAddress(uint(0))
	}

	return *baseQ.Page
}

func (baseQ *BaseQuery) GetPageSize() uint {
	if baseQ.PageSize == nil || *baseQ.PageSize == 0 {
		baseQ.PageSize = utils.GetDataTypeAddress(uint(10))
	}

	if *baseQ.PageSize > 100 {
		baseQ.PageSize = utils.GetDataTypeAddress(uint(100))
	}

	return *baseQ.PageSize
}

func (baseQ *BaseQuery) GetOffset() uint {
	return baseQ.GetPage() * baseQ.GetPageSize()
}
