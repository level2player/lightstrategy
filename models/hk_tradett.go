package models

import (
	"lightstrategy/domain"
	"log"

	"gopkg.in/mgo.v2/bson"
)

//港股交易日志
type HKTradett struct {
	StockCode string `json:"stock_code"`
	TradeLog  []struct {
		TickerId   int     `json:"ticker_id"`
		Cancelled  string  `json:"cancelled"`
		TradePrice float64 `json:"price"`
		AggrQty    int     `json:"aggr_qty"`
		TradeTime  string  `json:"trade_time"`
		TradeType  int     `json:"trade_type"`
	} `json:"trade_Log"`
}

func (hkTradett *HKTradett) InsertHKTradett() error {
	c := domain.Session.DB("stockdb").C("hktradett")
	err := c.Insert(&hkTradett)
	if err != nil {
		log.Printf("mgo insert, err=:%s", err.Error())
		return err
	}
	log.Printf("Insert mongo suc, ticked_id=%d", len(hkTradett.TradeLog))
	return nil
}

func (hkTradett *HKTradett) FindHKTradett(stock_code string) error {
	c := domain.Session.DB("stockdb").C("hktradett")
	err := c.Find(bson.M{"stockcode": stock_code}).One(&hkTradett)

	if err != nil {
		log.Println("mgo find err=:" + err.Error())
		return err
	}
	return nil
}
