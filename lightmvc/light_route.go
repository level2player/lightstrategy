package lightmvc

import "net/http"

type ControllerInterface interface {
	Handle_rquest(w http.ResponseWriter, r *http.Request)
}

func Router(pattern string, Controller ControllerInterface) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		Controller.Handle_rquest(w, r)
	})
}
