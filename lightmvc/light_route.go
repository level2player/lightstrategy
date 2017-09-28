package lightmvc

import (
	"html/template"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"
)

func Router(pattern string) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		log.Println("recive msg...")
		timeNow := time.Now()
		parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
		controller := Controllerdic[parts[len(parts)-1]]
		if controller != nil {
			v := reflect.ValueOf(controller.(ControllerInterface))
			method := v.MethodByName(methodlist[r.Method])
			if method.IsValid() {
				parms := []reflect.Value{reflect.ValueOf(w), reflect.ValueOf(r)}
				method.Call(parms)
			}
		} else {
			t, err := template.ParseFiles("template/html/notfound.html")
			if err != nil {
				log.Println(err)
			}
			t.Execute(w, nil)
		}
		defer func() {
			log.Printf("execute time:%s", time.Since(timeNow))
			r.Body.Close()
		}()
	})
}

var methodlist map[string]string

func init() {
	methodlist = make(map[string]string)
	methodlist["GET"] = "Get"
	methodlist["POST"] = "Post"
}
