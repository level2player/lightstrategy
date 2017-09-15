package main

import (
	"lightstrategy/controllers"
	"lightstrategy/lightmvc"
	"log"
	"net/http"
)

func main() {
	lightmvc.Router("/")
	lightmvc.RegisterController("webapi", controllers.WebapiController{})
	err := http.ListenAndServe(":8098", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
