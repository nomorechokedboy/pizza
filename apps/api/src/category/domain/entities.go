package domain

type Category struct {
	Id          int
	Name        string
	Description string
}

type CategoryQuery struct {
	Page     int
	PageSize int
	Q        string
}

type WriteCategoryBody struct {
	Name        string
	Description string
}
