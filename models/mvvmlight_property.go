package models

import (
	"errors"
	"strings"
	//"fmt"
)

type MvvmlightProperty struct {
	PropertyType   string
	PropertyName   string
	PropertyRemark string
}
type ConversionResults struct {
	IsOk               bool   `json:"isok"`
	ConversionsContent string `json:"conversions_content"`
}

func ProcessData(ProcessDataText string) (MvvmlightProperty, error) {
	tmptext := strings.TrimSpace(ProcessDataText)
	tmptext2 := strings.Replace(tmptext, "///", "", strings.Count(tmptext, "///"))
	tmptext2 = strings.Replace(tmptext2, "<summary>", "", strings.Count(tmptext, "<summary>"))
	tmptext2 = strings.Replace(tmptext2, "</summary>", "", strings.Count(tmptext, "<summary>"))
	tmptext2 = strings.Replace(tmptext2, "public", "", strings.Count(tmptext, "public"))
	tmptext2 = strings.Replace(tmptext2, "{ get; set; }", "", strings.Count(tmptext, "{ get; set; }"))
	arry := strings.Fields(tmptext2)
	reulst := MvvmlightProperty{}

	if len(arry)==2{
		reulst.PropertyType=arry[len(arry)-2]
		reulst.PropertyName=arry[len(arry)-1]
	}else if len(arry)==3{
		reulst.PropertyRemark = arry[len(arry)-3]
		reulst.PropertyType=arry[len(arry)-2]
		reulst.PropertyName=arry[len(arry)-1]
	}else if len(arry)>3{
		reulst.PropertyRemark = " "
		reulst.PropertyType=arry[len(arry)-2]
		reulst.PropertyName=arry[len(arry)-1]
	}else {
		return reulst, errors.New("ProcessData error")
	}

	return reulst, nil
}
