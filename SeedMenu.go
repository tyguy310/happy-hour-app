package main

import (
	"log"

	"github.com/coopernurse/gorp"
	_ "github.com/lib/pq"
)

func populateMenus(menusMap *gorp.DbMap) {

  err := menusMap.Insert(
    initMenu("Libations", "Drink", 1),
    initMenu("Food", "Food", 2),
    initMenu("Drink", "Drink", 2),
    initMenu("Swill", "Drink", 1),
    initMenu("Munchies", "Food", 3),
  )

	if err != nil {
		log.Fatal(err)
	}
}
