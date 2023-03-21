package main

import (
	"log"
	"pkg/database"
)

func main() {

	var err error

	db := database.Db{
		PathToFile: "db.json",
	}

	// example 1

	doc := map[string]interface{}{
		"title":    "The Fairy Tale",
		"author":   "Stephen King",
		"price":    20.5,
		"in_stock": true,
	}

	err = db.Insert(doc)
	if err != nil {
		log.Printf("Could not insert doc: %s\n", err)
		return
	}

	// example 2

	doc = map[string]interface{}{
		"title":    "Dagon",
		"author":   "Hovard Lovecraft",
		"price":    12.99,
		"in_stock": false,
	}

	err = db.Insert(doc)
	if err != nil {
		log.Printf("Could not insert doc: %s\n", err)
		return
	}
}