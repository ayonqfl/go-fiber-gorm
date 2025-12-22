package models

import "time"

type RmsSymbolRestrictionList struct {
	ID                 uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	GroupID            uint      `json:"group_id" gorm:"not null;index"`
	GroupRestrictionID uint      `json:"suspension_id" gorm:"not null;index"`
	Symbol             string    `json:"symbol" gorm:"type:varchar(255)"`
	CreatedAt          time.Time `json:"created_at" gorm:"autoCreateTime"`

	// Relationship
	GroupRestriction *RmsGroupRestriction `json:"group_restriction,omitempty" gorm:"foreignKey:GroupRestrictionID"`
}

// TableName overrides table name
func (RmsSymbolRestrictionList) TableName() string {
	return "rms_symbol_restriction_list"
}
