package models

type Member struct {
	ID         uint   `gorm:"primaryKey"`
	Types      string `gorm:"size:10" json:"types"`
	CoreID     uint   `json:"core_id"`
	ProfileID  string `json:"profile_id"`
	Experience string `json:"experience"`
	RoleID     string `json:"role_id"`
}
