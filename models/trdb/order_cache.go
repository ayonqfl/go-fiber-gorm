package models

type OrderCache struct {
    ID                     uint    `json:"id" gorm:"primaryKey;autoIncrement"`
    Exchange               string  `json:"exchange"`
    TradeSession           string  `json:"trade_session"`
    BrokerID               string  `json:"broker_id"`
    BoardType              string  `json:"board_type"`
    SymbolCode             string  `json:"symbol_code"`
    SymbolIsin             string  `json:"symbol_isin"`
    SymbolAssetclass       string  `json:"symbol_assetclass"`
    SymbolCategory         string  `json:"symbol_category"`
    CompulsorySpot         string  `json:"compulsory_spot"`
    BoAcc                  string  `json:"bo_acc"`
    ClientName             string  `json:"client_name"`
    ClientCode             string  `json:"client_code"`
    OrderAction            string  `json:"order_action"`
    OrderType              string  `json:"order_type"`
    OrderID                string  `json:"order_id"`
    ReforderID             string  `json:"reforder_id"`
    ChainID                string  `json:"chain_id"`
    OrderSide              string  `json:"order_side"`
    OrderQty               int64   `json:"order_qty"`
    OrderValue             float64 `json:"order_value"`
    OrderPrice             float64 `json:"order_price"`
    OrderDate              string  `json:"order_date"`
    OrderTime              string  `json:"order_time"`
    OrderValidity          string  `json:"order_validity"`
    DripQty                int64   `json:"drip_qty"`
    MinQty                 int64   `json:"min_qty"`
    DueQty                 int64   `json:"due_qty"`
    CumQty                 int64   `json:"cum_qty"`
    LastQty                int64   `json:"last_qty"`
    LastPx                 float64 `json:"last_px"`
    AvgPx                  float64 `json:"avg_px"`
    OrderStatus            string  `json:"order_status"`
    ExecStatus             string  `json:"exec_status"`
    EngineID               string  `json:"engine_id"`
    TradeMatchID           string  `json:"trade_match_id"`
    TimeInForce            string  `json:"time_in_force"`
    TradeDate              string  `json:"trade_date"`
    SettleDate             string  `json:"settle_date"`
    GrossTradeAmt          float64 `json:"gross_trade_amt"`
    AgressorIndicator      string  `json:"agressor_indicator"`
    LimitOrderDate         string  `json:"limit_order_date"`
    LimitOrderExpiryDate   string  `json:"limit_order_expiry_date"`
    LimitOrderType         string  `json:"limit_order_type"`
    PvtLimitOrder          string  `json:"pvt_limit_order"`
    PvtMktOrder            string  `json:"pvt_mkt_order"`
    StopLoss               float64 `json:"stop_loss"`
    TakeProfit             float64 `json:"take_profit"`
    TradeLogComment        string  `json:"trade_log_comment"`
    Emergency              string  `json:"emergency"`
    OrderRemarks           string  `json:"order_remarks"`
    AuthorizeType          string  `json:"authorize_type"`
    FixWorkstationID       string  `json:"fix_workstation_id"`
    TraderWsID             string  `json:"trader_ws_id"`
    OwnerWsID              string  `json:"owner_ws_id"`
    UserID                 string  `json:"user_id"`
    UserRole               string  `json:"user_role"`
    UserDevice             string  `json:"user_device"`
    UserIP                 string  `json:"user_ip"`
    RefUserID              string  `json:"ref_user_id"`
    WsGroupID              string  `json:"ws_group_id"`
    Branch                 string  `json:"branch"`
    FixMsg                 string  `json:"fix_msg"`
    OrderYield             float64 `json:"order_yield"`
    ExecYield              float64 `json:"exec_yield"`
    SeqNum                 int     `json:"seq_num"`
    SendToEngine           bool    `json:"send_to_engine"`
    OrderPlacedTime        string  `json:"order_placed_time"`
    Associate              string  `json:"associate"`
    TempDealer             string  `json:"temp_dealer"`
    ExternalUserID         string  `json:"external_user_id"`
    IsBasket               bool    `json:"is_basket" gorm:"default:false"`
    IsPriority             bool    `json:"is_priority" gorm:"default:false"`
    AccruedInterestAmt     float64 `json:"accrued_interest_amt" gorm:"default:0.0"`
}

// TableName overrides the table name
func (OrderCache) TableName() string {
    return "order_cache"
}