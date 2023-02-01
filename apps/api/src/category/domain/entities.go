package domain

type Category struct {
	Id          int
	Name        string
	Description string
}

type WriteCategoryBody struct {
	Name        string
	Description string
}
