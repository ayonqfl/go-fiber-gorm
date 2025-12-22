package models

import "time"

type RmsGroupList struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	GroupID    uint      `json:"group_id" gorm:"not null;index"`
	GroupValue string    `json:"group_value" gorm:"type:varchar(255)"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`

	// Relationship
	Group *RmsGroup `json:"group,omitempty" gorm:"foreignKey:GroupID"`
}

// TableName overrides table name
func (RmsGroupList) TableName() string {
	return "rms_group_list"
}
