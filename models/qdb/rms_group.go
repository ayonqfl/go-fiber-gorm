package models

import "time"

type RmsGroup struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	GroupName string    `json:"name" gorm:"type:varchar(255)"`
	GroupType string    `json:"group_type" gorm:"type:varchar(50)"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`

	// Relationships
	GroupValues       []RmsGroupList        `json:"group_value" gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE"`
	GroupRestrictions []RmsGroupRestriction `json:"group_restriction" gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE"`
}

// TableName overrides table name
func (RmsGroup) TableName() string {
	return "rms_group"
}
