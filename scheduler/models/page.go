package models

type Page[T any] struct {
	Items      []T `json:"items"`
	TotalItems int `json:"totalItems"`
}

func NewPage[T any](items []T, totalItems int) Page[T] {
	return Page[T]{
		Items:      items,
		TotalItems: totalItems,
	}
}

type PageParams struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

func (p PageParams) Limit() int {
	return p.PageSize
}

func (p PageParams) Offset() int {
	return (p.Page - 1) * p.PageSize
}
