package models

import "time"

type Project struct {
	ID            uint       `gorm:"primaryKey" json:"id"`
	Logo          string     `json:"logo"`
	Owner         string     `json:"owner"`
	Title         string     `json:"title"`
	Company       string     `json:"company"`
	StartDate     *time.Time `json:"start_date"`
	EndDate       *time.Time `json:"end_date"`
	Types         string     `json:"types"`
	Duration      string     `json:"duration"`      // "day" | "month" | "year"
	Participation string     `json:"participation"` // "remote" | "onsite" | "hybrid"
	Approval      string     `json:"approval"`      // "yes" | "no"
	Description   string     `json:"description"`
	Salary        float64    `json:"salary"`
	Priority      string     `json:"priority"`
	Status        string     `json:"status"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	Address       Address    `gorm:"foreignKey:ProjectID"`
	Members       []Member   `gorm:"foreignKey:ProjectID"`
	Tags          []Tag      `gorm:"many2many:project_tags"`
}
