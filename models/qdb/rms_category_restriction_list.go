package models

import "time"

type RmsCategoryRestrictionList struct {
	ID                 uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	GroupID            uint      `json:"group_id" gorm:"not null;index"`
	GroupRestrictionID uint      `json:"suspension_id" gorm:"not null;index"`
	Category            string    `json:"category" gorm:"type:varchar(255)"`
	CreatedAt           time.Time `json:"created_at" gorm:"autoCreateTime"`

	// Relationship
	GroupRestriction *RmsGroupRestriction `json:"group_restriction,omitempty" gorm:"foreignKey:GroupRestrictionID"`
}

// TableName overrides table name
func (RmsCategoryRestrictionList) TableName() string {
	return "rms_category_restriction_list"
}
