package models

type Role struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}
