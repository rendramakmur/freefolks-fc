package model

type Pagination struct {
	Page       int
	Limit      int
	TotalPages int
	TotalItems int
	Items      []interface{}
}
