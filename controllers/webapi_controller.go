package controllers

import (
	"encoding/json"
	"net/http"
)

type WebapiController struct {
}

type insertStockinfoReulst struct {
	is_insert_success bool
	sum               int
}

func (webapiController WebapiController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var actionReulst = insertStockinfoReulst{}
	actionReulst.is_insert_success = true
	actionReulst.sum = 1
	OutputJson(w, actionReulst)
}
func (webapiController WebapiController) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var actionReulst = insertStockinfoReulst{}
	actionReulst.is_insert_success = true
	actionReulst.sum = 1
	OutputJson(w, actionReulst)
}

func OutputJson(w http.ResponseWriter, output_object interface{}) {
	b, err := json.Marshal(output_object)
	if err != nil {
		return
	}
	w.Write(b)
}
