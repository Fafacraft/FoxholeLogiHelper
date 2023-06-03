package stockpile

import (
	"html/template"
	"net/http"

	"FLH/pkg/utils"

	"FLH/internal/common"
)

type DataStockpile struct {
	Id string
}

/*
Handle the stockpile page
"/stockpile"
*/
func HandleStockpile(w http.ResponseWriter, r *http.Request) {
	// TODO : bla bla fill indexData, then executeFinishPage
	dataIndex := common.DataIndex{}
	tpl, err := template.ParseFiles("static/templates/stockpile/stockpile.html")
	utils.CheckErr(err)

	stockpile := DataStockpile{}
	stockpile.Id = "bMat"
	dataIndex.Main = common.Execute(tpl, stockpile)

	common.ExecuteFinishPage(w, dataIndex)
}
