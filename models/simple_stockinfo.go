package models

import (
	"lightstrategy/domain"
	"log"

	"gopkg.in/mgo.v2/bson"
)

type SimpleStockinfo struct {
	StockCode    string `json:"stock_code"`
	ExchangeType string `json:"exchange_type"`
	StockName    string `json:"stock_name"`
	HisData      []struct {
		TradeDate            int64   `json:"trade_date"`
		ClosePrice           float64 `json:"close_price"`
		HighPrice            float64 `json:"high_price"`
		LowPrice             float64 `json:"low_price"`
		OpenPrice            float64 `json:"open_price"`
		PreClolsePrice       float64 `json:"preClolse_price"`
		UpdownOver           float64 `json:"updown_over"`           //涨跌额
		UpdownRange          float64 `json:"updown_range"`          //涨跌幅
		TurnoveRate          float64 `json:"turnove_rate"`          //换手率
		TradeVolume          float64 `json:"trade_volume"`          //成交量
		Turnover             float64 `json:"turnover"`              //成交金额
		MarketValue          float64 `json:"market_value"`          //总市值
		CirculateMarketValue float64 `json:"circulate_marketValue"` //流通总市值
	} `json:"his_data"`
}

func (StockInfo *SimpleStockinfo) InsertSimpleStockInfo() error {
	c := domain.Session.DB("stockdb").C("simplestocklist")
	err := c.Insert(&StockInfo)
	if err != nil {
		log.Println("mgo insert, err=:" + err.Error())
		return err
	}
	log.Println("Insert mongo suc, StockCode=" + StockInfo.StockCode)
	return nil
}

func (StockInfo *SimpleStockinfo) FindStockRecord(StockCode string, ExchangeType string) {
	c := domain.Session.DB("stockdb").C("simplestocklist")
	err := c.Find(bson.M{"stockcode": StockCode, "exchangetype": ExchangeType}).One(&StockInfo)
	if err != nil {
		log.Println("mgo find err=:" + err.Error())
	}
}

func (StockInfo *SimpleStockinfo) FindStockHisDate(StockCode string, ExchangeType string, tradedate int) {
	c := domain.Session.DB("stockdb").C("simplestocklist")
	err := c.Find(bson.M{"stockcode": StockCode, "exchangetype": ExchangeType}).Select(bson.M{"stockcode": 1, "exchangetype": 1, "stockname": 1, "hisdata": bson.M{"$elemMatch": bson.M{"tradedate": tradedate}}}).One(&StockInfo)
	if err != nil {
		log.Println("mgo find err=:" + err.Error())
	}
}

func GetStockRecordCount() int {
	c := domain.Session.DB("stockdb").C("simplestocklist")
	count, err := c.Count()
	if err != nil {
		log.Println("mgo find err=:" + err.Error())
	}
	return count
}
