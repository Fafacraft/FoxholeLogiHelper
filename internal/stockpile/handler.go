package stockpile

import (
	"database/sql"
	"fmt"
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

	tpl, err := template.ParseFiles("static/templates/stockpile/stockpile.html")
	utils.CheckErr(err)

	stockpile := DataStockpile{}
	// prepare different list with different priorities
	stockpile.ItemBoxList01 = makeItemBoxList("priority", 1)
	stockpile.ItemBoxList02 = makeItemBoxList("priority", 2)
	stockpile.ItemBoxList03 = makeItemBoxList("priority", 3)
	stockpile.ItemBoxList04 = makeItemBoxList("priority", 4)
	stockpile.ItemBoxList05 = makeItemBoxList("priority", 5)
	stockpile.ItemBoxList06 = makeItemBoxList("priority", 6)
	stockpile.ItemBoxList07 = makeItemBoxList("priority", 7)
	stockpile.ItemBoxList08 = makeItemBoxList("priority", 8)

	dataIndex.Main = common.Execute(tpl, stockpile)

	common.ExecuteFinishPage(w, dataIndex)
}

/*
list all items with appropriate :
- all
- priority  (int)
- category  (string)
*/
func makeItemBoxList(option string, value any) []template.HTML {
	finalList := []template.HTML{}

	// open the db
	db, err := sql.Open("sqlite3", "./database.db")
	utils.CheckErr(err)
	// do the request with appropriate option
	var rows *sql.Rows
	switch option {
	case "all":
		preparedRequest, err := db.Prepare("SELECT * FROM stockpile")
		utils.CheckErr(err)
		rows, err = preparedRequest.Query()
		break
	case "priority":
		preparedRequest, err := db.Prepare("SELECT * FROM stockpile WHERE priority == ?")
		utils.CheckErr(err)
		rows, err = preparedRequest.Query(value)
		break
	case "category":
		preparedRequest, err := db.Prepare("SELECT * FROM stockpile WHERE category == ?")
		utils.CheckErr(err)
		rows, err = preparedRequest.Query(value)
		break
	default:
		fmt.Println("ERROR : Unreconized option in makeItemBoxList()")
	}
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
	fmt.Println(finalList)

	return finalList
}
