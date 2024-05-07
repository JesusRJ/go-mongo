package core

type PaginationQuery[T any] struct {
	Page    int
	PerPage int
	Order   OrderBy
	Query   *T
}

type OrderBy struct {
	Direction string
	Field     string
}

type Pagination[T any] struct {
	Page  int `json:"page"`
	Pages int `json:"pages"`
	Total int `json:"total"`
	Data  []T `json:"data"`
}
