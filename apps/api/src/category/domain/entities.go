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
	Name        string `validate:"required,min=3,max=20"`
	Description string `validate:"required,max=1000"`
}
