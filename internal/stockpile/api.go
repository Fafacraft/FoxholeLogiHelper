package stockpile

import (
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

	// Perform the necessary actions to update the item count in the database
	// You can replace this with your own logic

	// For example,
	if request.IsAdd {
		itemNumber = +1
	} else {
		itemNumber = -1
	}

	// Create the response
	response := ItemActionResponse{
		Status:     "success",
		ItemNumber: itemNumber,
	}

	// Return the response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
