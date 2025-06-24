package models

import (
	"gorm.io/datatypes"
	"time"
)

type Project struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	Logo          string         `json:"logo"`
	Owner         string         `gorm:"size:50" json:"owner"`
	Title         string         `gorm:"size:100" json:"title"`
	Company       string         `gorm:"size:100" json:"company"`
	StartDate     *time.Time     `json:"start_date"`
	EndDate       *time.Time     `json:"end_date"`
	Types         string         `gorm:"size:10" json:"types"`
	Duration      string         `gorm:"size:10" json:"duration"`
	Participation string         `gorm:"size:10" json:"participation"`
	Address       datatypes.JSON `json:"address"` // format: struct => json.Marshal
	Approval      string         `gorm:"size:5" json:"approval"`
	Description   string         `json:"description"`
	Tags          datatypes.JSON `json:"tags"` // format: array struct => json.Marshal
	Salary        float64        `gorm:"type:numeric(15,2)" json:"salary"`
	Priority      string         `gorm:"size:10" json:"priority"`
	IsPrivate     bool           `json:"is_private"`
	IsClosed      bool           `json:"is_closed"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
}
