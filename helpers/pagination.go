package helpers

import (
	"github.com/IlhamLamp/cmty-project-service/dto"
)

func SanitizePagination(filter *dto.CoreFilter) {
	if filter.Page <= 0 {
		filter.Page = 1
	}
	if filter.Limit <= 0 {
		filter.Limit = 10
	}
}

func BuildPaginationMeta(totalRecords int64, currentPage, limit int) *dto.PaginationMeta {
	totalPages := int((totalRecords + int64(limit) - 1) / int64(limit))
	var nextPage, prevPage *int

	if currentPage < totalPages {
		np := currentPage + 1
		nextPage = &np
	}
	if currentPage > 1 {
		pp := currentPage - 1
		prevPage = &pp
	}

	return &dto.PaginationMeta{
		TotalRecords: totalRecords,
		CurrentPage:  currentPage,
		TotalPages:   totalPages,
		NextPage:     nextPage,
		PrevPage:     prevPage,
	}
}
