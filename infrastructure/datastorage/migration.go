package datastorage

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// VersionedDB main function for database migration
func VersionedDB(db *Data) error {

	err := db.DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	var instanceConfig database.Driver
	var errConfig error

	switch DbDriver := os.Getenv("DB_DRIVER"); DbDriver {
	case "postgres":
		instanceConfig, errConfig = postgres.WithInstance(db.DB, &postgres.Config{})
	}

	if errConfig != nil {
		log.Fatal(errConfig)
	}

	version, errVersion := migrationUp(instanceConfig)
	if errVersion != nil {
		if strings.Contains(errVersion.Error(), "no change") {
			errVersion = nil
		} else {
			errVersion = fmt.Errorf("error fatal in up migration version %d , %s", version, errVersion.Error())
		}
	}
	return errVersion
}

// migrationUp method that triggers versioning scripts
func migrationUp(instanceConfig database.Driver) (int, error) {
	pathScripts := os.Getenv("SCRIPTS_PATH")
	DBName := os.Getenv("DB_NAME")

	migration, err := migrate.NewWithDatabaseInstance(
		pathScripts,
		DBName, instanceConfig)
	if err != nil {
		log.Fatalf("Error Connection Drive Migration Up %s", err.Error())
	}

	var version uint
	var errVersion error

	err = migration.Up()

	if err != nil {
		if !strings.Contains(err.Error(), "no change") {
			version, _, errVersion = migration.Version()
			if errVersion != nil {
				log.Fatal("Error Get version in Up")
			}

			err = migration.Force(int(version) - 1)
			if err != nil {
				log.Fatalf("Error Forced Migration %s", err.Error())
			}
		}
	}
	return int(version), err
}
