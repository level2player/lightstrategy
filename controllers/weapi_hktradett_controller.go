package controllers

import (
	"encoding/json"
	"io/ioutil"
	"lightstrategy/core"
	"lightstrategy/models"
	"net/http"
	"strconv"
)

type WebapiHKTradettController struct {
}

func (webapiController WebapiHKTradettController) Get(w http.ResponseWriter, r *http.Request) {

	action_in, err := strconv.Atoi(r.FormValue("action_in"))
	if err != nil {
		var actionReulst = models.InsertStockinfoReulst{ErrorInfo: err.Error()}
		core.OutputJson(w, actionReulst)
		return
	} else {
		if action_in == 0 {
			stock_code := r.FormValue("stock_code")
			hkTradett := models.HKTradett{}
			hkTradett.FindHKTradett(stock_code)
			core.OutputJson(w, hkTradett)
		}
	}
}
func (webapiController WebapiHKTradettController) Post(w http.ResponseWriter, r *http.Request) {
	hkTradett := models.HKTradett{}
	content, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	json.Unmarshal(content, &hkTradett)
	err := hkTradett.InsertHKTradett()
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
