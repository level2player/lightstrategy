package controllers

import (
	"encoding/json"
	"io/ioutil"
	"lightstrategy/core"
	"lightstrategy/models"
	"net/http"
)

type WebapInsertStockInfoController struct {
}

func (webapiController WebapInsertStockInfoController) Get(w http.ResponseWriter, r *http.Request) {
	var actionReulst = models.InsertStockinfoReulst{}
	actionReulst.IsInsertSuccess = false
	actionReulst.InsertSum = 0
	core.OutputJson(w, actionReulst)
}
func (webapiController WebapInsertStockInfoController) Post(w http.ResponseWriter, r *http.Request) {
	simpleStockinfo := models.SimpleStockinfo{}
	content, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	json.Unmarshal(content, &simpleStockinfo)
	err := simpleStockinfo.InsertSimpleStockInfo()
	var actionReulst = models.InsertStockinfoReulst{}
	if err != nil {
		actionReulst.IsInsertSuccess = false
		actionReulst.InsertSum = 0
		actionReulst.ErrorInfo = err.Error()
	} else {
		actionReulst.IsInsertSuccess = true
		actionReulst.InsertSum = 1
	}
	core.OutputJson(w, actionReulst)
}
