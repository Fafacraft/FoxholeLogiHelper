package stockpile

import (
	"database/sql"
	"html/template"
	"net/http"

	"FLH/pkg/utils"

	"FLH/internal/common"

	_ "github.com/mattn/go-sqlite3"
)

type DataStockpile struct {
	ItemBoxList01 []template.HTML
	ItemBoxList02 []template.HTML
	ItemBoxList03 []template.HTML
	ItemBoxList04 []template.HTML
	ItemBoxList05 []template.HTML
	ItemBoxList06 []template.HTML
	ItemBoxList07 []template.HTML
	ItemBoxList08 []template.HTML
}

type DataItemBox struct {
	Id        string
	Name      string
	Category  string
	Priority  int
	Nb        int
	NbNeeded  int
	ImageLink string
	BMats     int
	RMats     int
	EMats     int
	HeMats    int

	Class string
}

/*
Handle the stockpile page
"/stockpile"
*/
func HandleStockpile(w http.ResponseWriter, r *http.Request) {
	// TODO : bla bla fill indexData, then executeFinishPage
	dataIndex := common.DataIndex{}

	var moreOptions string
	sortingMethod := r.URL.Query().Get("sorting")
	// set the moreOptions
	if sortingMethod == "bMatsOnly" {
		moreOptions += " AND (rMats == 0 AND eMats == 0 AND heMats == 0)"
	}
	if sortingMethod == "rMats" {
		moreOptions += " AND (rMats != 0)"
	}
	if sortingMethod == "eMats" {
		moreOptions += " AND (eMats != 0)"
	}

	tpl, err := template.ParseFiles("static/templates/stockpile/stockpile.html")
	utils.CheckErr(err)

	stockpile := DataStockpile{}
	// prepare different list with different priorities
	stockpile.ItemBoxList01 = makeItemBoxList("priority == 1" + moreOptions)
	stockpile.ItemBoxList02 = makeItemBoxList("priority == 2" + moreOptions)
	stockpile.ItemBoxList03 = makeItemBoxList("priority == 3" + moreOptions)
	stockpile.ItemBoxList04 = makeItemBoxList("priority == 4" + moreOptions)
	stockpile.ItemBoxList05 = makeItemBoxList("priority == 5" + moreOptions)
	stockpile.ItemBoxList06 = makeItemBoxList("priority == 6" + moreOptions)
	stockpile.ItemBoxList07 = makeItemBoxList("priority == 7" + moreOptions)
	stockpile.ItemBoxList08 = makeItemBoxList("priority == 8" + moreOptions)

	dataIndex.Main = common.Execute(tpl, stockpile)

	common.ExecuteFinishPage(w, dataIndex)
}

/*
list all items.

can use condition to add a WHERE clause with this condition
ex:
condition = "rMats == 0 && priority == 1"
does ; "SELECT * FROM stockpile WHERE rMats == 0 && priority == 1"
*/
func makeItemBoxList(condition string) []template.HTML {
	finalList := []template.HTML{}

	// open the db
	db, err := sql.Open("sqlite3", "./database.db")
	utils.CheckErr(err)
	// do the request with appropriate option
	var rows *sql.Rows

	rawRequest := "SELECT * FROM stockpile"
	if condition != "" {
		rawRequest += " WHERE " + condition
	}
	preparedRequest, err := db.Prepare(rawRequest)
	utils.CheckErr(err)
	rows, err = preparedRequest.Query()

	defer rows.Close() // THE THING INSIDE DEFER IS EVALUATED HERE, EVEN IF IT RUNS AT THE END

	// get each item resulting
	for rows.Next() {
		itemBox := DataItemBox{}
		err = rows.Scan(&itemBox.Id, &itemBox.Name, &itemBox.Category, &itemBox.Priority, &itemBox.Nb, &itemBox.NbNeeded, &itemBox.ImageLink, &itemBox.BMats, &itemBox.RMats, &itemBox.EMats, &itemBox.HeMats)
		utils.CheckErr(err)

		// get the border color
		itemBox.Class = FindBorderColorClass(itemBox.Id)

		tplRaw, err := template.ParseFiles("static/templates/stockpile/itemBox.html")
		utils.CheckErr(err)
		tpl := *common.Execute(tplRaw, itemBox)
		finalList = append(finalList, tpl)
	}

	return finalList
}
