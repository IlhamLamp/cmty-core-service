package models

type Tag struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `json:"name"`
	Color string `json:"color"`
}
