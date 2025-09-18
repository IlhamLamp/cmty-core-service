package services

import (
	"github.com/IlhamLamp/cmty-core-service/dto"
	"github.com/IlhamLamp/cmty-core-service/models"
	"github.com/IlhamLamp/cmty-core-service/repository"
)

type ProjectService interface {
	Create(project *models.Project) error
	BulkCreate(projects []models.Project) error
	GetAll(filter dto.CoreFilter) ([]models.Project, int64, error)
	GetByID(id uint) (*models.Project, error)
	Update(project *models.Project) error
	Delete(id uint) error
	Clean() (int64, error)
}

type projectService struct {
	repo repository.ProjectRepository
}

func NewProjectService(r repository.ProjectRepository) ProjectService {
	return &projectService{r}
}

func (s *projectService) Create(project *models.Project) error {
	return s.repo.Create(project)
}

func (s *projectService) BulkCreate(projects []models.Project) error {
	return s.repo.BulkCreate(projects)
}

func (s *projectService) GetAll(filter dto.CoreFilter) ([]models.Project, int64, error) {
	return s.repo.GetAll(filter)
}

func (s *projectService) GetByID(id uint) (*models.Project, error) {
	return s.repo.GetByID(id)
}

func (s *projectService) Update(project *models.Project) error {
	return s.repo.Update(project)
}

func (s *projectService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *projectService) Clean() (int64, error) {
	return s.repo.Clean()
}
