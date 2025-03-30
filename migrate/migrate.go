package main

import (
	"fmt"

	"github.com/DaigoSugiyama0317/map-memo-app/db"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate()
}