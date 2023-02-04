package common

type BaseQuery struct {
	Page     *int    `query:"page"`
	PageSize *int    `query:"pageSize"`
	Q        *string `query:"q"`
}
