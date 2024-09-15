package httpresponse

import "github.com/mazharul-islam/utils"

type Pagination struct {
	Items      []interface{} `json:"items"`
	TotalItems int           `json:"totalItems"`
	TotalPages int           `json:"totalPages"`
}

type ResponsesPagination struct {
	Items      interface{} `json:"items"`
	TotalItems int         `json:"totalItems"`
	TotalPages int         `json:"totalPages"`
}

func ToResourcePaginationResponse(count, size uint, items []any) Pagination {
	return Pagination{
		TotalItems: int(count),
		TotalPages: utils.CalculatePages(count, size),
		Items:      items,
	}
}
