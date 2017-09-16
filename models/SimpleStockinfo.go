package models

type SimpleStockinfo struct {
	tradeDate            int64   `json:"trade_date"`
	stockCode            string  `json:"stock_code"`
	stockName            string  `json:"stock_name"`
	closePrice           float64 `json:"close_price"`
	highPrice            float64 `json:"high_price"`
	lowPrice             float64 `json:"low_price"`
	openPrice            float64 `json:"open_price"`
	preClolsePrice       float64 `json:"preClolse_price"`
	updownOver           float64 `json:"updown_over"`           //涨跌额
	updownRange          float64 `json:"updown_range"`          //涨跌幅
	turnoveRate          float64 `json:"turnove_rate"`          //换手率
	tradeVolume          float64 `json:"trade_volume"`          //成交量
	turnover             float64 `json:"turnover"`              //成交金额
	marketValue          float64 `json:"market_value"`          //总市值
	circulateMarketValue float64 `json:"circulate_marketValue"` //流通总市值
}
