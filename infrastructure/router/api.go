package router

import (
	"database/sql"
	"log"

	"github.com/samuskitchen/beer-api-clean-arch/infrastructure/database"
)

// Start started api
func Start(port string) {

	// connection to the database.
	db := database.New()
	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
			log.Println("panic occurred:", err)
		}
	}(db.DB)

	//Versioning the database
	err := database.VersionedDB(db)
	if err != nil {
		log.Fatal(err)
	}

	server := newServer(port, db)

	// start the server.
	server.Start()
}
