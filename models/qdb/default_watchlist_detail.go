package models

type DefaultWatchlistDetail struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	WatchlistID uint   `json:"watchlist_id" gorm:"not null;index"`
	Instrument  string `json:"instrument" gorm:"type:varchar(50);not null"`

	// Relationship
	Watchlist DefaultWatchlist `json:"-" gorm:"foreignKey:WatchlistID"`
}

// TableName overrides table name
func (DefaultWatchlistDetail) TableName() string {
	return "default_watchlist_details"
}
