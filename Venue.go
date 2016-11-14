package main

type Venue struct {
	Id   int `db:"id" json: "id"`
	Name string `db:"name" json: "name"`
	Phone string `db:"phone" json: "phone"`
	Address string `db:"address" json: "address"`
	WebURL string `db:"weburl" json: "weburl"`
	Genre string `db:"genre" json: "genre"`
	Rooftop bool `db:"rooftop" json: "rooftop"`
	Patio bool `db:"patio" json: "patio"`
}

func initVenue(name string, phone string, address string, weburl string, genre string, rooftop bool, patio bool) *Venue {
	return &Venue{
		Name: name,
		Phone: phone,
		Address: address,
		WebURL: weburl,
		Genre: genre,
		Rooftop: rooftop,
		Patio: patio,
	}
}
