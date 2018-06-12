package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetPage from server and render spaces in page to browser view
func GetPage(w http.ResponseWriter, r *http.Request) {
	var page string

	// Get page nr from url
	vars := mux.Vars(r)
	page = vars["page"]
	defer r.Body.Close()

	// If no page number in url then change to first page
	if page == "" {
		page = strconv.Itoa(1)
	}

	// Get page information from server
	req, err := http.Get(baseUrl + page)
	if err != nil {
		log.Printf("Error getting http request: %s", err)
		return
	}

	// Create new struct d containing information from page on server
	d := new(Data)
	d.Page, _ = strconv.Atoi(page)

	// Decode JSON to struct d
	if err := json.NewDecoder(req.Body).Decode(&d); err != nil {
		log.Printf("Error decoding body: %s", err)
		return
	}

	// Set current_page active and the rest of the pages deactive
	for index, pageOnIndex := range p.Pages {
		if pageOnIndex.CurrentNr == d.Page {
			p.Pages[index].Active = "active"
		} else {
			p.Pages[index].Active = "deactive"
		}
	}

	// Set prev page
	p.PrevNr = d.Page - 1
	if p.PrevNr < p.StartNr {
		p.PrevNr = p.StartNr
	}

	// Set next page
	p.NextNr = d.Page + 1
	if p.NextNr > p.EndNr {
		p.NextNr = p.EndNr
	}

	// Write data from page on server and pagination information to browser view
	t := template.Must(template.ParseFiles("src/templates/index.html"))
	t.Execute(w, M{
		"Data":       d,
		"Pagination": p,
	})
}
