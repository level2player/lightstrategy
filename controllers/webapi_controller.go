package controllers

import (
	"encoding/json"
	"lightstrategy/lightmvc"
	"log"
	"net/http"
)

type Webapi_controller struct {
	lightmvc.ControllerInterface
}
type insert_stockinfo_reulst struct {
	is_insert_success bool
	sum               int
}

func (this *Webapi_controller) Handle_rquest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	log.Println("handel request")
	var action_reulst = insert_stockinfo_reulst{}
	action_reulst.is_insert_success = true
	action_reulst.sum = 1
	OutputJson(w, action_reulst)
}
func OutputJson(w http.ResponseWriter, output_object interface{}) {
	b, err := json.Marshal(output_object)
	if err != nil {
		return
	}
	w.Write(b)
}
