package common

import (
	"bytes"
	"html/template"
	"net/http"

	"FLH/pkg/utils"
)

/*
Execute the finished Page, with index.html

w : http.ResponseWritter
data : DataIndex{} filled, Main included
*/
func ExecuteFinishPage(w http.ResponseWriter, data DataIndex) {
	tpl, err := template.ParseFiles("static/templates/index.html")
	utils.CheckErr(err)

	tpl.Execute(w, data)
}

/*
Execute a template.Template (from ParseFiles) to return a usable *template.HTML

tpl : *template.Template
data : any corresponding struct of data
*/
func Execute(tpl *template.Template, data any) *template.HTML {
	var buf bytes.Buffer
	err := tpl.Execute(&buf, data)
	utils.CheckErr(err)

	tplHTML := template.HTML(buf.String())
	return &tplHTML
}
