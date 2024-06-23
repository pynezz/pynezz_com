package middleware

import (
	"fmt"
	"strings"

	"github.com/pynezz/pynezzentials/ansi"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	// database names
	Users   = "users"
	Content = "content"
)

var dbNames = map[string]string{
	Users:   "users.db",
	Content: "content.db",
	// Remember to add new databases here if needed in the future
}

// Database defines the structure of the database. We're using SQLite in our project.
type Database struct {
	name   string
	Driver *gorm.DB
}

var DBInstance *Database // The global database instance

// https://gosamples.dev/sqlite-intro/

// NewDatabase creates a new database. It returns a pointer to the database.
func InitDatabaseDriver(db *gorm.DB) *Database {
	ansi.PrintInfo("Initializing new database driver...")
	DBInstance = &Database{
		Driver: db,
	}
	return DBInstance
}

// Initialize the database with the given name and configuration, and automigrate the given tables
func InitDB(database string, conf gorm.Config, tables ...interface{}) (*gorm.DB, error) {
	if _, ok := isValidDb(database); !ok && database != "" {
		return nil, fmt.Errorf("database name missing or invalid. Format: <name>.db or <name")
	} else {
		database = chkExt(database)
	}

	db, err := gorm.Open(sqlite.Open(database+".db"), &conf)
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(tables...); err != nil {
		return nil, err
	}

	db = db.Session(&gorm.Session{CreateBatchSize: 100})

	return db, nil
}

// Migrate creates the tables in the database
func (d *Database) Migrate() error {
	if d.Driver == nil {
		d.Driver = DBInstance.Driver
	}

	ansi.PrintInfo("Migrating the database...")
	return nil
}

func isValidDb(database string) (string, bool) {
	database = chkExt(database)
	for _, db := range dbNames {
		if db == database+".db" {
			ansi.PrintSuccess("Database name is valid: " + database)
			return database, true
		}
	}
	ansi.PrintError("Database name is invalid: " + database)
	return "", false
}

// chkExt checks every file with the extension .db
func chkExt(database string) string {
	strParts := strings.Split(database, ".")
	if l := len(strParts); l > 1 && strParts[1] == "db" {
		return strings.Split(database, ".")[0]
	} else if l > 1 && strParts[1] != "db" {
		return ""
	}
	return database
}
