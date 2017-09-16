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
	if err != nil {
		log.Println(err)
	}
	log.Println(simpleStockinfo)
	var actionReulst = models.InsertStockinfoReulst{}
	actionReulst.IsInsertSuccess = true
	actionReulst.InsertSum = 1
	core.OutputJson(w, actionReulst)
}
