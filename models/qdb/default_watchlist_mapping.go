package models

type WatchlistMappingType string

const (
	WatchlistMappingUser  WatchlistMappingType = "user"
	WatchlistMappingGroup WatchlistMappingType = "group"
)

type DefaultWatchlistMapping struct {
	ID          uint                 `json:"id" gorm:"primaryKey;autoIncrement"`
	WatchlistID uint                 `json:"watchlist_id" gorm:"not null;index"`
	GroupID     uint                 `json:"group_id" gorm:"not null"`
	Type        WatchlistMappingType `json:"type" gorm:"type:varchar(20);not null"`

	Watchlist DefaultWatchlist `json:"-" gorm:"foreignKey:WatchlistID"`
}

func (DefaultWatchlistMapping) TableName() string {
	return "default_watchlist_mapping"
}