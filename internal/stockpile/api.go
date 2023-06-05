package stockpile

import (
	"FLH/pkg/utils"
	"database/sql"
	"encoding/json"
	"net/http"
)

type ItemActionRequest struct {
	IsAdd bool   `json:"isAdd"` // adding item or deleting item
	Id    string `json:"id"`    // id of the item we're changing
}

type ItemActionResponse struct {
	Status     string `json:"status"`
	ItemNumber int    `json:"itemNumber"` // new number of item
	Class      string `json:"class"`
	Error      string `json:"error,omitempty"`
}

/*
Handle the adding or deleting of an item
*/
func UpdateItemCountHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request payload
	var request ItemActionRequest
	var itemNumber int
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// do the thing
	itemNumber = addOrDeleteItem(request.IsAdd, request.Id)
	class := FindBorderColorClass(request.Id)

	// Create the response
	response := ItemActionResponse{
		Status:     "success",
		ItemNumber: itemNumber,
		Class:      class,
	}

	// Return the response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

/*
Add or Delete one to the stockpile, depending on id and isAdd
Return the updated value and the css class for color border
*/
func addOrDeleteItem(isAdd bool, id string) int {
	// open the db
	db, err := sql.Open("sqlite3", "./database.db")
	utils.CheckErr(err)

	// get Nb
	var nb int
	preparedRequest, err := db.Prepare("SELECT nb FROM stockpile WHERE id == ?")
	utils.CheckErr(err)
	preparedRequest.QueryRow(id).Scan(&nb)

	// set newNb
	var newNb int
	if isAdd {
		newNb = nb + 1
	} else {
		newNb = nb - 1
		if newNb < 0 {
			newNb = 0
		}
	}

	// update db
	preparedRequest, err = db.Prepare("UPDATE stockpile SET nb = ? WHERE id == ?")
	utils.CheckErr(err)
	preparedRequest.QueryRow(newNb, id).Scan(&nb)

	return newNb
}

/*
Return the css class which correspond to the need in the stockpile

Need the id of the item
*/
func FindBorderColorClass(id string) string {
	// open the db
	db, err := sql.Open("sqlite3", "./database.db")
	utils.CheckErr(err)

	// get Nb and NbNeeded
	var nb, nbNeeded int
	preparedRequest, err := db.Prepare("SELECT nb, nbNeeded FROM stockpile WHERE id == ?")
	utils.CheckErr(err)
	preparedRequest.QueryRow(id).Scan(&nb, &nbNeeded)

	// find the class
	var class string
	if nb == 0 {
		class = "borderCritical"
	} else if nb <= nbNeeded/2 {
		class = "borderLow"
	} else if nb < nbNeeded {
		class = "borderNotFull"
	} else if nb < nbNeeded*2 {
		class = "borderFull"
	} else {
		class = "borderOverfilled"
	}

	return class
}
