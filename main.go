package main

import (
	"lightstrategy/controllers"
	"lightstrategy/lightmvc"
	"log"
	"net/http"
)

func main() {
	lightmvc.RegisterController("webapinsertstockinfo", controllers.WebapInsertStockInfoController{})
	lightmvc.RegisterController("webapifindstockinfo", controllers.WebapiFindStockInfoController{})
	lightmvc.Router("/")
	err := http.ListenAndServe(":8098", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
