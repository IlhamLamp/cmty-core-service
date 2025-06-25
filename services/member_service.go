package services

import (
	"github.com/IlhamLamp/cmty-project-service/models"
	"github.com/IlhamLamp/cmty-project-service/repository"
)

type MemberService interface {
	Create(member *models.Member) error
	BulkCreate(members []models.Member) error
	GetAll() ([]models.Member, error)
	GetByID(id uint) (*models.Member, error)
	Update(member *models.Member) error
	Delete(id uint) error
	Clean() (int64, error)
}

type memberService struct {
	repo repository.MemberRepository
}

func NewMemberService(r repository.MemberRepository) MemberService {
	return &memberService{r}
}

func (s *memberService) Create(member *models.Member) error {
	return s.repo.Create(member)
}

func (s *memberService) BulkCreate(members []models.Member) error {
	return s.repo.BulkCreate(members)
}

func (s *memberService) GetAll() ([]models.Member, error) {
	return s.repo.GetAll()
}

func (s *memberService) GetByID(id uint) (*models.Member, error) {
	return s.repo.GetByID(id)
}

func (s *memberService) Update(member *models.Member) error {
	return s.repo.Update(member)
}

func (s *memberService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *memberService) Clean() (int64, error) {
	return s.repo.Clean()
}
