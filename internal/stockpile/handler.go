package stockpile

import (
	"html/template"
	"net/http"

	"FLH/pkg/utils"

	"FLH/internal/common"
)

type DataStockpile struct {
	ItemBoxList []template.HTML
}

type DataItemBox struct {
	Id        string
	Title     string
	ItemNb    int
	ImageLink string
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
	stockpile.ItemBoxList = makeItemBoxList()
	dataIndex.Main = common.Execute(tpl, stockpile)

	common.ExecuteFinishPage(w, dataIndex)
}

func makeItemBoxList() []template.HTML {
	finalList := []template.HTML{}

	// get the db then for each item
	itemBox := DataItemBox{}
	itemBox.Id = "bMat"
	itemBox.Title = "Basic Materials"
	itemBox.ItemNb = 5
	itemBox.ImageLink = "./static/images"
	tplRaw, err := template.ParseFiles("static/templates/stockpile/itemBox.html")
	utils.CheckErr(err)
	tpl := *common.Execute(tplRaw, itemBox)
	finalList = append(finalList, tpl)

	itemBox2 := DataItemBox{}
	itemBox2.Id = "gas_mask"
	itemBox2.Title = "Gas Mask"
	itemBox2.ItemNb = 3
	itemBox2.ImageLink = "./static/images"
	tplRaw2, err := template.ParseFiles("static/templates/stockpile/itemBox.html")
	utils.CheckErr(err)
	tpl2 := *common.Execute(tplRaw2, itemBox2)
	finalList = append(finalList, tpl2)

	return finalList
}
