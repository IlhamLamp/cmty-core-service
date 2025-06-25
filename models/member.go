package models

type Member struct {
	ID         uint   `gorm:"primaryKey"`
	ProjectID  uint   `json:"project_id"`
	ProfileID  string `json:"profile_id"`
	Experience string `json:"experience"`
	RoleID     uint   `json:"role_id"`
}
