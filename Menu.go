package main

type Menu struct {
	Id int `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Type string `db:"category" json:"category"`
	Venue_id int `db:"venue_id" json:"venue_id"`
}

func initMenu(name string, category string, venue_id int) *Menu {
	return &Menu{
		Name: name,
		Type: category,
		Venue_id: venue_id,
	}
}
