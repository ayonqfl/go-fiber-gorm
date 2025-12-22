package models

import "time"

type RmsTickerGroupRestrictionList struct {
	ID                  uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	GroupID             uint      `json:"group_id" gorm:"not null;index"`
	GroupRestrictionID  uint      `json:"suspension_id" gorm:"not null;index"`
	TickerGroup         string    `json:"ticker_group" gorm:"type:varchar(255)"`
	CreatedAt           time.Time `json:"created_at" gorm:"autoCreateTime"`

	// Relationship
	GroupRestriction *RmsGroupRestriction `json:"group_restriction,omitempty" gorm:"foreignKey:GroupRestrictionID"`
}

// TableName overrides table name
func (RmsTickerGroupRestrictionList) TableName() string {
	return "rms_ticker_group_restriction_list"
}
