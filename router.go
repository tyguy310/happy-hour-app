package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/coopernurse/gorp"
	"github.com/gorilla/mux"
)

func FileServerRouteG(m *mux.Router, path, dir string) {
	m.PathPrefix(path).Handler(
		http.StripPrefix(path, http.FileServer(http.Dir(dir))))
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Some front end jazz")
}

//fetchVenues takes in a GORP DbMap and fetches all venues in the
//venues database table into a Venue slice
func fetchVenues(dbMap *gorp.DbMap) []Venue {
	var venues []Venue

	_, err := dbMap.Select(&venues, "SELECT * FROM venues")

	if err != nil {
		log.Fatal(err)
	}

	return venues
}

//makeVenuesRoute takes in a GORP DbMap and makes a route that uses that
//DbMap to fetch all venues in the venues table and serve them as JSON.
func makeVenuesRoute(dbMap *gorp.DbMap) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		venues := fetchVenues(dbMap)

		venuesJSON, err := json.Marshal(venues)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(w, "%s", venuesJSON)
	}
}

//initRouter takes in a GORP DbMap and initializes the router's routes while
//using the DbMap to handle database functionality.
func initRouter(dbMap *gorp.DbMap) *mux.Router {
	r := mux.NewRouter()

	//Add the venues route API with makeVenuesRoute
	r.HandleFunc("/venues", makeVenuesRoute(dbMap))

	//Serve all other requests with index.html, and ultimately the front-end
	//Angular.js app.
	r.PathPrefix("/").HandlerFunc(indexRoute)

	return r
}
