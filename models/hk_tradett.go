package models

import (
	"lightstrategy/domain"
	"log"

	cache "github.com/patrickmn/go-cache"
	"gopkg.in/mgo.v2/bson"
)

type TradeLog struct {
	TickerId   int     `json:"ticker_id"`
	Cancelled  string  `json:"cancelled"`
	TradePrice float64 `json:"price"`
	AggrQty    int     `json:"aggr_qty"`
	TradeTime  string  `json:"trade_time"`
	TradeType  int     `json:"trade_type"`
}

//港股交易日志
type HKTradett struct {
	StockCode string     `json:"stock_code"`
	TradeLog  []TradeLog `json:"trade_Log"`
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
	hkTradett.StockCode = stock_code
	if x, found := domain.FindCache(stock_code); found {
		hkTradett.TradeLog = x.(*HKTradett).TradeLog
		return nil
	} else {
		c := domain.Session.DB("stockdb").C("hktradett")
		err := c.Find(bson.M{"stockcode": stock_code}).One(hkTradett)
		if err != nil {
			log.Println("mgo find err=:" + err.Error())
			return err
		}
		domain.InsertCache(stock_code, hkTradett, cache.DefaultExpiration)
		return nil
	}
}
