package repository

import (
	"github.com/IlhamLamp/cmty-project-service/models"
	"gorm.io/gorm"
)

type MemberRepository interface {
	Create(member *models.Member) error
	BulkCreate(members []models.Member) error
	GetAll() ([]models.Member, error)
	GetByID(id uint) (*models.Member, error)
	Update(member *models.Member) error
	Delete(id uint) error
	Clean() (int64, error)
}

type memberRepository struct {
	db *gorm.DB
}

func NewMemberRepository(db *gorm.DB) MemberRepository {
	return &memberRepository{db}
}

func (r *memberRepository) Create(member *models.Member) error {
	return r.db.Create(member).Error
}

func (r *memberRepository) BulkCreate(members []models.Member) error {
	return r.db.Create(&members).Error
}

func (r *memberRepository) GetAll() ([]models.Member, error) {
	var members []models.Member
	err := r.db.Find(&members).Error
	return members, err
}

func (r *memberRepository) GetByID(id uint) (*models.Member, error) {
	var member models.Member
	err := r.db.First(&member, id).Error
	return &member, err
}

func (r *memberRepository) Update(member *models.Member) error {
	return r.db.Save(member).Error
}

func (r *memberRepository) Delete(id uint) error {
	return r.db.Delete(&models.Member{}, id).Error
}

func (r *memberRepository) Clean() (int64, error) {
	result := r.db.Exec("DELETE FROM members")
	return result.RowsAffected, result.Error
}
