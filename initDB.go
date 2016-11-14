package main

import (
	"database/sql"
	"log"

	"github.com/coopernurse/gorp"
	_ "github.com/lib/pq"
)

func initDB() *gorp.DbMap {
	db, err := sql.Open("postgres", "postgres://tylermaier:password@localhost/happy_hour?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

  venuesTable :=
  	dbMap.AddTableWithName(Venue{}, "venues").SetKeys(true, "Id")

	venuesTable.ColMap("Name").SetNotNull(true).SetUnique(true)

  menusTable := dbMap.AddTableWithName(Menu{}, "menus").SetKeys(true, "Id")

  menusTable.ColMap("Name").SetNotNull(true).SetUnique(true)

	err = dbMap.DropTablesIfExists()
	if err != nil {
		log.Fatal(err)
	}

	err = dbMap.CreateTablesIfNotExists()
	if err != nil {
		log.Fatal(err)
	}

	populateVenues(dbMap)
  populateMenus(dbMap)

	return dbMap
}
