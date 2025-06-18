package models

type Address struct {
	ID        uint   `gorm:"primaryKey"`
	ProjectID uint   `json:"project_id"`
	Street    string `json:"street"`
	City      string `json:"city"`
	State     string `json:"state"`
	ZipCode   string `json:"zip_code"`
}
