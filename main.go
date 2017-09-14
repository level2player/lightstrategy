package main

import (
	"lightstrategy/controllers"
	"lightstrategy/lightmvc"
	"net/http"
)

func main() {
	//http.HandleFunc("/data/", NotFoundHandler)
	lightmvc.Router("/data/","Insert", &controllers.Webapi_controller{})
	http.ListenAndServe(":8098", nil)
}



// func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
// 	t, err := template.ParseFiles("template/html/hello.html")
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	t.Execute(w, nil)
// }
