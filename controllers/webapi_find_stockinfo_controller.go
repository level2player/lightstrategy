package controllers

import (
	"lightstrategy/core"
	"lightstrategy/models"
	"net/http"
)

type WebapiFindStockInfoController struct {
}

func (webapiController WebapiFindStockInfoController) Get(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	stockcode := r.Form["stock_code"]
	exchangetype := r.Form["exchange_type"]
	simpleStockinfo := models.SimpleStockinfo{}
	defer r.Body.Close()
	core.OutputJson(w, simpleStockinfo.FindStockRecord(stockcode[0], exchangetype[0]))
}
func (webapiController WebapiFindStockInfoController) Post(w http.ResponseWriter, r *http.Request) {

}