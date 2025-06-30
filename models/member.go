package models

type Member struct {
	ID         uint   `gorm:"primaryKey"`
	ProjectID  uint   `json:"project_id"`
	ProfileID  string `json:"profile_id"`
	Experience string `json:"experience"`
	RoleID     string `json:"role_id"`
}
