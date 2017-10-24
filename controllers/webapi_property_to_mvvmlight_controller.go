package controllers

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"lightstrategy/core"
	"lightstrategy/models"
	"log"
	"net/http"
	"strings"
)

type WebapiPropertyToMvvmlightController struct {
}

func (webapiController WebapiPropertyToMvvmlightController) Get(w http.ResponseWriter, r *http.Request) {

}
func (webapiController WebapiPropertyToMvvmlightController) Post(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println(err)
	}
	codetext := strings.TrimSpace(string(content))
	propertyArry := strings.SplitAfter(codetext, "}")
	var mvvmlightPropertySlice []models.MvvmlightProperty
	for _, property := range propertyArry {
		mvvm, err := models.ProcessData(property)
		if err != nil {
			log.Println(err)
		} else {
			mvvmlightPropertySlice = append(mvvmlightPropertySlice, mvvm)
		}
	}
	var conversionResult models.ConversionResults
	if len(mvvmlightPropertySlice) > 0 {
		tpl := template.Must(template.ParseFiles("./template/tpl/mvvmlightProperty.tpl"))
		conversionsContentbuffer := bytes.NewBuffer(make([]byte, 0))
		err := tpl.Execute(conversionsContentbuffer, mvvmlightPropertySlice)
		if err != nil {
			log.Println(err)
			conversionResult.IsOk = false
		} else {
			conversionResult.IsOk = true
			conversionResult.ConversionsContent = conversionsContentbuffer.String()
		}
	} else {
		conversionResult.IsOk = false
	}
	fmt.Printf("转换结果:\n\r %s \n\r", conversionResult.ConversionsContent)
	core.OutputJson(w, conversionResult)
}
