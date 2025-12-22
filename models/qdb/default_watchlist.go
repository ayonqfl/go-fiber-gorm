package models

type DefaultWatchlist struct {
	ID   uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"type:varchar(255);not null"`
	Type string `json:"type"`

	// Relationships
	Instruments []DefaultWatchlistDetail  `json:"instruments" gorm:"foreignKey:WatchlistID;constraint:OnDelete:CASCADE"`
	Mappings    []DefaultWatchlistMapping `json:"mappings" gorm:"foreignKey:WatchlistID;constraint:OnDelete:CASCADE"`
}

// TableName overrides table name
func (DefaultWatchlist) TableName() string {
	return "default_watchlist"
}
