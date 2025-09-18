package helpers

import (
	"github.com/IlhamLamp/cmty-core-service/dto"
)

func SanitizePagination(f dto.Paginatable) {
	if f.GetPage() <= 0 {
		f.SetPage(1)
	}
	if f.GetLimit() <= 0 {
		f.SetLimit(10)
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
