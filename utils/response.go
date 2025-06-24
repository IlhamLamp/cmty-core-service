package utils

import (
	"net/http"

	"github.com/IlhamLamp/cmty-project-service/types"
	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context, data interface{}, message string) {
	ctx.JSON(http.StatusOK, types.ApiResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func Created(ctx *gin.Context, data interface{}, message string) {
	ctx.JSON(http.StatusCreated, types.ApiResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func Error(ctx *gin.Context, status int, err error, message string) {
	ctx.JSON(status, types.ApiResponse{
		Status:  "error",
		Message: message,
		Error:   err.Error(),
	})
}
