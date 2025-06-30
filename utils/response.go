package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiResponse struct {
	Status  string
	Message string
	Data    interface{}
	Error   string
}

func Success(ctx *gin.Context, data interface{}, message string) {
	ctx.JSON(http.StatusOK, ApiResponse{
		Status:  "success",
		Message: message,
		Data:    data,
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
