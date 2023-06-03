package main

import (
	"fmt"
	"log"
	"net/http"

	"FLH/internal/stockpile"
)

/*
Start server and list handlers
*/
func main() {

	fs1 := http.FileServer(http.Dir("static/images"))
	http.Handle("/static/images/", http.StripPrefix("/static/images/", fs1))

	http.HandleFunc("/", handleHomePage)
	http.HandleFunc("/stockpile", stockpile.HandleStockpile)

	// Start the HTTP server
	fmt.Println("Server started on port :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

/*
Handle "/"
*/
func handleHomePage(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/stockpile", http.StatusSeeOther)
}
