package main

import (
	"lightstrategy/controllers"
	"lightstrategy/lightmvc"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
)
var RootCmd = &cobra.Command{
    Use: "-v",
    Run: func(cmd *cobra.Command, args []string) {
        println("lightstrategy version is 0.0.1")
	},
}
var UpCmd= &cobra.Command{
	Use: "up",
    Run: func(cmd *cobra.Command, args []string) {
        InitHttpServer()
	},
}
func init(){
    RootCmd.AddCommand(UpCmd)
}
func InitHttpServer(){
	lightmvc.RegisterController("webapinsertstockinfo", controllers.WebapInsertStockInfoController{})
	lightmvc.RegisterController("webapifindstockinfo", controllers.WebapiFindStockInfoController{})
	lightmvc.RegisterController("webapipropertytomvvmlight", controllers.WebapiPropertyToMvvmlightController{})
	lightmvc.RegisterController("webapihktradett", controllers.WebapiHKTradettController{})
	lightmvc.Router("/")
	err := http.ListenAndServe(":8098", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
func main() {
	if err := RootCmd.Execute(); err != nil {
        log.Println(err)
        os.Exit(1)
    }
}
