package lightmvc

import (
	"net/http"
)

type ControllerInterface interface {
	Get(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
}

var Controllerdic map[string]interface{}

func RegisterController(key string, obj interface{}) bool {
	_, isexist := Controllerdic[key]
	if !isexist {
		Controllerdic[key] = obj
	} else {
		return false
	}
	return true
}
func init() {
	Controllerdic = make(map[string]interface{})
}
