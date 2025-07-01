package utils

import (
	"net/http"

	"github.com/IlhamLamp/cmty-project-service/dto"
	"github.com/gin-gonic/gin"
)

type ApiResponse struct {
	Status     string              `json:"status"`
	Message    string              `json:"message"`
	Data       interface{}         `json:"data,omitempty"`
	Pagination *dto.PaginationMeta `json:"pagination,omitempty"`
	Error      string              `json:"error,omitempty"`
}

func Success(ctx *gin.Context, data interface{}, message string, pagination *dto.PaginationMeta) {
	ctx.JSON(http.StatusOK, ApiResponse{
		Status:     "success",
		Message:    message,
		Data:       data,
		Pagination: pagination,
	})
}

func Created(ctx *gin.Context, data interface{}, message string) {
	ctx.JSON(http.StatusCreated, ApiResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func Error(ctx *gin.Context, status int, err error, message string) {
	ctx.JSON(status, ApiResponse{
		Status:  "error",
		Message: message,
		Error:   err.Error(),
	})
}
