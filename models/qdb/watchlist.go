package models

type Watchlist struct {
    ID               uint    `json:"id" gorm:"primaryKey;autoIncrement"`
    ClnID            string  `json:"cln_id"`
    Instrument       string  `json:"instrument"`
    PriceAlertLevel  float64 `json:"price_alert_level"`
    AlertCondition   string  `json:"alert_condition"`
    NewsAlert        bool    `json:"news_alert"`
    NewsCondition    string  `json:"news_condition"`
    WatchlistName    string  `json:"watchlist_name"`
    PriceAlert       bool    `json:"price_alert"`
    Precedence       int     `json:"precedence" gorm:"default:0"`
}

// TableName overrides the table name
func (Watchlist) TableName() string {
    return "stock_watchlist"
}