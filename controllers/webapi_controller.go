package controllers

import (
	"encoding/json"
	"io/ioutil"
	"lightstrategy/core"
	"lightstrategy/models"
	"log"
	"net/http"
)

type WebapiController struct {
}

func (webapiController WebapiController) Get(w http.ResponseWriter, r *http.Request) {
	var actionReulst = models.InsertStockinfoReulst{}
	actionReulst.IsInsertSuccess = false
	actionReulst.InsertSum = 0
	core.OutputJson(w, actionReulst)
}
func (webapiController WebapiController) Post(w http.ResponseWriter, r *http.Request) {
	simpleStockinfo := models.SimpleStockinfo{}
	content, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	err := json.Unmarshal(content, &simpleStockinfo)
	var actionReulst = models.InsertStockinfoReulst{}
	if err != nil {
		actionReulst.IsInsertSuccess = false
		actionReulst.InsertSum = 0
		log.Println(err)
	} else {
		actionReulst.IsInsertSuccess = true
		actionReulst.InsertSum = 1
	}
	log.Println(simpleStockinfo)
	core.OutputJson(w, actionReulst)
}
