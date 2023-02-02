package domain

type Category struct {
	Id          int
	Name        string
	Description string
}

type CategoryQuery struct {
	Page     int    `query:"page"`
	PageSize int    `query:"pageSize"`
	Q        string `query:"q"`
}

type WriteCategoryBody struct {
	Name        string
	Description string
}
