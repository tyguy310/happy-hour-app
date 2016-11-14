package main

import (
	"log"

	"github.com/coopernurse/gorp"
	_ "github.com/lib/pq"
)

func populateVenues(venuesMap *gorp.DbMap) {

	err := venuesMap.Insert(
    initVenue("Steubens", "3038301001", "523 E 17th Ave, Denver, CO 80203", "http://steubens.com", "American", false, true),
    initVenue("Root Down", "3039934200", "1600 W 33rd Ave, Denver, CO 80211", "http://rootdowndenver.com", "American", false, true),
    initVenue("Work & Class", "3032920700", "2500 Larimer St, Denver, CO 80205", "http://workandclassdenver.com", "Latin American", false, false),
  )

	if err != nil {
		log.Fatal(err)
	}
}
