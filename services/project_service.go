package services

import (
	"github.com/IlhamLamp/cmty-project-service/models"
	"github.com/IlhamLamp/cmty-project-service/repository"
)

type ProjectService interface {
	Create(project *models.Project) error
	BulkCreate(projects []models.Project) error
	GetAll() ([]models.Project, error)
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

func (s *projectService) GetAll() ([]models.Project, error) {
	return s.repo.GetAll()
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
