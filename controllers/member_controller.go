package controllers

import (
	"net/http"
	"strconv"

	"github.com/IlhamLamp/cmty-project-service/models"
	"github.com/IlhamLamp/cmty-project-service/services"
	"github.com/IlhamLamp/cmty-project-service/utils"
	"github.com/gin-gonic/gin"
)

type MemberController struct {
	service services.MemberService
}

func NewMemberController(s services.MemberService) *MemberController {
	return &MemberController{}
}

func (c *MemberController) Create(ctx *gin.Context) {
	var member models.Member
	if err := ctx.ShouldBindJSON(&member); err != nil {
		utils.Error(ctx, http.StatusBadRequest, err, "Invalid request payload")
		return
	}
}

func (c *MemberController) GetAll(ctx *gin.Context) {
	members, err := c.service.GetAll()
	if err != nil {
		utils.Error(ctx, http.StatusInternalServerError, err, "Failed to get all members")
		return
	}
	utils.Success(ctx, members, "Members retrieved successfully")
	ctx.JSON(http.StatusOK, members)
}

func (c *MemberController) Delete(ctx *gin.Context) {
	id, _ = strconv.Atoi(ctx.Param("id"))
	if err := c.service.Delete(uint(id)); err != nil {

	}
}
