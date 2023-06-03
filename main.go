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

	fsImages := http.FileServer(http.Dir("static/images"))
	http.Handle("/static/images/", http.StripPrefix("/static/images/", fsImages))

	fsCSS := http.FileServer(http.Dir("static/css"))
	http.Handle("/static/css/", http.StripPrefix("/static/css/", fsCSS))

	fsJS := http.FileServer(http.Dir("static/js"))
	http.Handle("/static/js/", http.StripPrefix("/static/js/", fsJS))

	http.HandleFunc("/", handleHomePage)
	http.HandleFunc("/stockpile", stockpile.HandleStockpile)

	// api
	http.HandleFunc("/api/stockpileUpdateItem", stockpile.UpdateItemCountHandler)

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
