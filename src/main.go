package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Global constants
const baseUrl = "https://thisopenspace.com/lhl-test?page="

// Global array to store Pagination
var p *Pagination

func init() {
	// Run newPagination at init to get pagination info
	// (e.g. total amount of pages)
	p = newPagination()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", GetPage).Methods("GET")
	r.HandleFunc("/{page}", GetPage).Methods("GET")
	r.PathPrefix("/styles/").Handler(http.StripPrefix("/styles/",
		http.FileServer(http.Dir("src/templates/styles/"))))
	http.Handle("/", r)

	fmt.Println("Running server!")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// create pagination
func newPagination() *Pagination {
	p := new(Pagination)

	// Set start page to 1
	page := 1
	p.StartNr = page
	valid := true

	// Loop through pages until we get a bad response from server
	for valid {

		url := baseUrl + strconv.Itoa(page)

		// Get response from url
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		// Valid response store page number to Array p
		// Not valid response means there is no more pages -> break loop
		if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
			p.Pages = append(p.Pages, Page{CurrentNr: page})
			page++

		} else {
			valid = false
		}
	}

	// Set last page number
	p.EndNr = page - 1
	return p
}
