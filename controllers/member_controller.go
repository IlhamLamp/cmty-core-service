package controllers

import (
	"encoding/json"
	"net/http"
	"os"
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
	return &MemberController{s}
}

func (c *MemberController) Create(ctx *gin.Context) {
	var member models.Member
	if err := ctx.ShouldBindJSON(&member); err != nil {
		utils.Error(ctx, http.StatusBadRequest, err, "Invalid request payload")
		return
	}
}

func (c *MemberController) GetAll(ctx *gin.Context) {
	// page, limit := helpers.GetPaginationParams(ctx)
	// members, total, err := c.service.GetAll(page, limit)
	// if err != nil {
	// 	utils.Error(ctx, http.StatusInternalServerError, err, "Failed to get all members")
	// 	return
	// }

	// meta := helpers.BuildPaginationMeta(total, page, limit)
	utils.Success(ctx, "test", "Members retrieved successfully", nil)
}

func (c *MemberController) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.service.Delete(uint(id)); err != nil {
		utils.Error(ctx, http.StatusInternalServerError, err, "Falied to delete member")
		return
	}

	utils.Success(ctx, nil, "Member deleted successfully", nil)
}

// -+-+-+-+ SEEDER HANDLER +-+-+-+-
func (c *MemberController) SeedMembers(ctx *gin.Context) {
	file, err := os.ReadFile("database/seeders/02_member.json")
	if err != nil {
		utils.Error(ctx, http.StatusInternalServerError, err, "Failed to read seed file")
		return
	}

	var members []models.Member
	if err := json.Unmarshal(file, &members); err != nil {
		utils.Error(ctx, http.StatusInternalServerError, err, "Failed to parse seed file")
		return
	}

	var validMembers []models.Member
	for _, m := range members {
		if m.ProjectID != 0 {
			validMembers = append(validMembers, m)
		}
	}

	if err := c.service.BulkCreate(validMembers); err != nil {
		utils.Error(ctx, http.StatusInternalServerError, err, "Failed to seed members")
		return
	}

	utils.Success(ctx, members, "Members seeded successfully", nil)
}

func (c *MemberController) CleanMembers(ctx *gin.Context) {
	rowsAffected, err := c.service.Clean()
	if err != nil {
		utils.Error(ctx, http.StatusInternalServerError, err, "Failed to clean members")
		return
	}

	message := "total rows affected: " + strconv.FormatInt(rowsAffected, 10)
	utils.Success(ctx, nil, "Members cleaned succesfully, "+message, nil)
}
