package model

import (
	"math"
)

const defaultSize = 12
const maxSize = 60

// Pagination pagination data
type Pagination struct {
	PageNumber int `json:"pageNumber"`
	PageCount  int `json:"pageCount"`
	PageSize   int `json:"pageSize"`
	TotalCount int `json:"totalCount"`
}

// Limit limit
func (p *Pagination) Limit() int {
	return p.PageSize
}

// Offset offset
func (p *Pagination) Offset() int {
	offset := (p.PageNumber - 1) * p.PageSize
	if offset < 0 {
		return 0
	}
	return offset
}

// SetTotal set total and recalculated data
func (p *Pagination) SetTotal(total int) {
	if total > 0 {
		pageCount := int(math.Ceil(float64(total) / float64(p.PageSize)))
		pageNumber := p.PageNumber
		if p.PageNumber <= 0 {
			pageNumber = 1
		}

		p.PageCount = pageCount
		p.PageNumber = pageNumber
		p.TotalCount = total
	}
}

// InitPagination new pagination with pageNumber and size
func InitPagination(pageNumber int, size int) Pagination {
	if pageNumber <= 0 {
		pageNumber = 0
	}
	if size <= 0 {
		size = defaultSize
	}
	if size > maxSize {
		size = maxSize
	}
	return Pagination{
		PageNumber: pageNumber,
		PageSize:   size,
	}
}

// PaginationWithTotal new pagination with total input
func PaginationWithTotal(pageNumber int, size int, total int) Pagination {
	if total <= 0 {
		return Pagination{
			PageNumber: pageNumber,
			PageSize:   size,
			TotalCount: 0,
			PageCount:  0,
		}
	}
	pageCount := int(math.Ceil(float64(total) / float64(size)))

	if pageNumber <= 0 {
		pageNumber = 1
	}

	if pageNumber > pageCount {
		pageNumber = pageCount
	}

	return Pagination{
		PageNumber: pageNumber,
		PageSize:   size,
		TotalCount: total,
		PageCount:  pageCount,
	}
}
