package models

import "time"

type RmsGroupRestriction struct {
	ID      uint `json:"id" gorm:"primaryKey;autoIncrement"`
	GroupID uint `json:"group_id" gorm:"not null;index"`

	TickerGroupBuy  bool `json:"ticker_group_buy" gorm:"default:false"`
	TickerGroupSell bool `json:"ticker_group_sell" gorm:"default:false"`

	CategoryBuy  bool `json:"category_buy" gorm:"default:false"`
	CategorySell bool `json:"category_sell" gorm:"default:false"`

	SymbolBuy  bool `json:"symbol_buy" gorm:"default:false"`
	SymbolSell bool `json:"symbol_sell" gorm:"default:false"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`

	// Relationships
	Group *RmsGroup `json:"group,omitempty" gorm:"foreignKey:GroupID"`

	TickerGroupRestrictions []RmsTickerGroupRestrictionList `json:"ticker_group_restriction_list" gorm:"foreignKey:GroupRestrictionID;constraint:OnDelete:CASCADE"`
	SymbolRestrictions      []RmsSymbolRestrictionList      `json:"symbol_restriction_list" gorm:"foreignKey:GroupRestrictionID;constraint:OnDelete:CASCADE"`
	CategoryRestrictions    []RmsCategoryRestrictionList    `json:"category_restriction_list" gorm:"foreignKey:GroupRestrictionID;constraint:OnDelete:CASCADE"`
}

// TableName overrides table name
func (RmsGroupRestriction) TableName() string {
	return "rms_group_restriction"
}
