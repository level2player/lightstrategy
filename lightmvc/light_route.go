package lightmvc

import (
	"net/http"
	"reflect"
	"log"
	)



type ControllerInterface interface {
	Handle_rquest(w http.ResponseWriter, r *http.Request)
}

func Router(pattern string, action_name string,Controller ControllerInterface) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		log.Println("a")
		method_value:=reflect.ValueOf(&Controller).MethodByName(action_name+"_action")
		if method_value.IsValid() {
			params := make([]reflect.Value, 1)                
			params[0] = reflect.ValueOf(w)                   
			params[1] = reflect.ValueOf(r)  
			log.Println("b")
			method_value.Call(params)
		}
	})
}

