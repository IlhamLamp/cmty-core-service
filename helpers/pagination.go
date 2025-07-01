package helpers

import (
	"strconv"

	"github.com/IlhamLamp/cmty-project-service/dto"
	"github.com/gin-gonic/gin"
)

func GetPaginationParams(ctx *gin.Context) (page, limit int) {
	page, _ = strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ = strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	return page, limit
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
