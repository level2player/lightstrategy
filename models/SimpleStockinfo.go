package models

type SimpleStockinfo struct {
	tradeDate            int
	stockCode            string
	stockName            string
	closePrice           float64
	highPrice            float64
	lowPrice             float64
	openPrice            float64
	preClolsePrice       float64
	updownOver           float64 //涨跌额
	updownRange          float64 //涨跌幅
	turnoveRate          float64 //换手率
	tradeVolume          float64 //成交量
	turnover             float64 //成交金额
	marketValue          float64 //总市值
	circulateMarketValue float64 //流通总市值
}
