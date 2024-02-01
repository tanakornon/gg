package gg

import "fmt"

type QueryParams struct {
	PaginationParams
	SearchParams
	SortParams
}

// PaginationParams is used to paginate entities

type PaginationParams struct {
	Page     int `form:"page"`
	PageSize int `form:"page_size"`
}

func (p *PaginationParams) IsPaginated() bool {
	return p.Page > 0 && p.PageSize > 0
}

func (p *PaginationParams) Offset() int {
	return (p.Page - 1) * p.PageSize
}

func (p *PaginationParams) Limit() int {
	return p.PageSize
}

// SearchParams is used to search entities by keyword

type SearchParams struct {
	Keyword string `form:"search"`
}

func (s *SearchParams) IsSearched() bool {
	return s.Keyword != ""
}

func (s *SearchParams) GetKeyword() string {
	return fmt.Sprintf("%%%s%%", s.Keyword)
}

// SortParams is used to sort entities

type SortParams struct {
	SortBy string `form:"sort_by"`
	Order  string `form:"order"   default:"asc"`
}

func (s *SortParams) IsSorted() bool {
	return s.SortBy != ""
}

func (s *SortParams) GetOrderBy() string {
	return fmt.Sprintf("%s %s", s.SortBy, s.Order)
}
