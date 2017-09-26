package core

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func OutputJson(w http.ResponseWriter, output_object interface{}) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")
	b, err := json.Marshal(output_object)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintf(w, string(b))
}
