package controllers

import (
	"lightstrategy/core"
	"lightstrategy/models"
	"net/http"
	"strconv"
)

type WebapiFindStockInfoController struct {
}

func (webapiController WebapiFindStockInfoController) Get(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	stockcode := r.FormValue("stock_code")
	exchangetype := r.FormValue("exchange_type")
	tradedate, err := strconv.Atoi(r.FormValue("trade_date"))
	simpleStockinfo := models.SimpleStockinfo{}
	if err != nil {
		simpleStockinfo.FindStockRecord(stockcode, exchangetype)
	} else {
		simpleStockinfo.FindStockHisDate(stockcode, exchangetype, tradedate)
	}
	core.OutputJson(w, simpleStockinfo)

}
func (webapiController WebapiFindStockInfoController) Post(w http.ResponseWriter, r *http.Request) {

}
