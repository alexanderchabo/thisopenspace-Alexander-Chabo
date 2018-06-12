package main

// M maps arrays to be render in browser
type M map[string]interface{}

// Page stores page number and if its currently active
type Page struct {
	Active    string `json:"active"`
	CurrentNr int    `json:"current_nr"`
}

// Pagination stores information regarding amount of pages
type Pagination struct {
	StartNr int    `json:"start_number"`
	EndNr   int    `json:"end_number"`
	PrevNr  int    `json:"prev_page_number"`
	NextNr  int    `json:"next_page_number"`
	Pages   []Page `json:"pages"`
}

// Data stores page info from server
type Data struct {
	Spaces   []Space `json:"data"`
	PageSize int     `json:"page_size"`
	Total    int     `json:"total"`
	Page     int     `json:"current_page"`
}

// Space store space info from page on server
type Space struct {
	Id                      int     `json:"id"`
	Name                    string  `json:"name"`
	Address                 string  `json:"address"`
	HourlyPrice             float64 `json:"hourly_price"`
	DailyPrice              float64 `json:"daily_price"`
	SquareFootage           int     `json:"square_footage"`
	Capacity                int     `json:"capacity"`
	ViewsCount              int     `json:"views_count"`
	PrimaryPhotoCssUrlSmall string  `json:"primary_photo_css_url_small"`
}
