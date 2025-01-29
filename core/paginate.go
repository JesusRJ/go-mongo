package core

var (
	_defaultPage int32 = 1
	_defaultSize int64 = 10
)

type QueryOptions struct {
	Page     *int32
	PageSize *int64
	Order    *OrderBy
}

const (
	Asc  int32 = 1
	Desc int32 = -1
)

type OrderBy struct {
	Direction int32
	Field     string
}

type Pagination[T any] struct {
	Page  int `json:"page"`
	Pages int `json:"pages"`
	Total int `json:"total"`
	Data  []T `json:"data"`
}

// Find creates a new FindOptions instance.
func Options() *QueryOptions {
	return &QueryOptions{
		Page:     &_defaultPage,
		PageSize: &_defaultSize,
	}
}

func (q *QueryOptions) SetPage(i int32) *QueryOptions {
	q.Page = &i
	return q
}

func (q *QueryOptions) SetPageSize(i int64) *QueryOptions {
	q.PageSize = &i
	return q
}

func (q *QueryOptions) SetOrder(o OrderBy) *QueryOptions {
	q.Order = &o
	return q
}
