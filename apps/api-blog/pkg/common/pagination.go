package common

type BaseQuery struct {
	Page     int    `query:"page"`
	PageSize int    `query:"pageSize"`
	Sort     string `query:"sort" validate:"oneof=asc ASC desc DESC"`
	SortBy   string `query:"sortBy"`
}

type BasePaginationResponse[T any] struct {
	Items    []T   `json:"items"`
	Page     int   `json:"page"`
	PageSize int   `json:"page_size"`
	Total    int64 `json:"total"`
}

func (query *BaseQuery) GetPage() int {
	if query.Page == 0 {
		query.Page = 1
	}

	return query.Page
}

func (query *BaseQuery) GetPageSize() int {
	switch {
	case query.PageSize <= 0:
		query.PageSize = 10
	case query.PageSize >= 100:
		query.PageSize = 100
	}

	return query.PageSize
}

func (query *BaseQuery) GetOffset() int {
	return (query.GetPage() - 1) * query.GetPageSize()
}

func (query *BaseQuery) GetSort() string {
	if query.Sort == "" {
		query.Sort = "ASC"
	}

	return query.Sort
}

func (query *BaseQuery) GetSortBy() string {
	if query.SortBy == "" {
		query.SortBy = "id"
	}

	return query.SortBy
}
