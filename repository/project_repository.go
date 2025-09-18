package repository

import (
	"github.com/IlhamLamp/cmty-core-service/dto"
	"github.com/IlhamLamp/cmty-core-service/helpers"
	"github.com/IlhamLamp/cmty-core-service/models"
	"gorm.io/gorm"
)

type ProjectRepository interface {
	Create(project *models.Project) error
	BulkCreate(projects []models.Project) error
	GetAll(filter dto.CoreFilter) ([]models.Project, int64, error)
	GetByID(id uint) (*models.Project, error)
	Update(project *models.Project) error
	Delete(id uint) error
	Clean() (int64, error)
}

type projectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &projectRepository{db}
}

func (r *projectRepository) Create(project *models.Project) error {
	return r.db.Create(project).Error
}

func (r *projectRepository) BulkCreate(projects []models.Project) error {
	return r.db.Create(&projects).Error
}

func (r *projectRepository) GetAll(filter dto.CoreFilter) ([]models.Project, int64, error) {
	var projects []models.Project
	var total int64

	query := r.db.Model(&models.Project{}).Preload("Members")
	query = helpers.FilterProjectsByType(query, filter)
	query.Count(&total)

	offset := (filter.Page - 1) * filter.Limit
	err := query.Limit(filter.Limit).Offset(offset).Find(&projects).Error

	return projects, total, err

}

func (r *projectRepository) GetByID(id uint) (*models.Project, error) {
	var project models.Project
	err := r.db.Preload("Members").Preload("Tags").First(&project, id).Error
	return &project, err
}

func (r *projectRepository) Update(project *models.Project) error {
	return r.db.Save(project).Error
}

func (r *projectRepository) Delete(id uint) error {
	return r.db.Delete(&models.Project{}, id).Error
}

func (r *projectRepository) Clean() (int64, error) {
	result := r.db.Exec("DELETE FROM projects")
	return result.RowsAffected, result.Error
}
