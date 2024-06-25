package middleware

import (
	"fmt"
	"strings"

	"github.com/pynezz/pynezz_com/internal/server/middleware/models"
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
	"main":  "main.db", // The main database
	// Remember to add new databases here if needed in the future
}

// Database defines the structure of the database. We're using SQLite in our project.
type Database struct {
	Tables map[string]interface{}
	Driver *gorm.DB
}

var DBInstance *Database // The global database instance

func init() {
	conf := gorm.Config{
		PrepareStmt: true,
	}

	mainBase, err := InitDB("main.db", conf, models.User{}, models.Admin{})
	if err != nil {
		ansi.PrintError(err.Error())
	}

	// Migrate the database
	if err := mainBase.AutoMigrate(&models.User{}, &models.Admin{}); err != nil {
		ansi.PrintError(err.Error())
	}

	DBInstance = &Database{
		Driver: mainBase,
		Tables: make(map[string]interface{}),
	}

	DBInstance.SetDriver(mainBase)
	DBInstance.AddTable(models.User{}, "users")
	DBInstance.AddTable(models.Admin{}, "admins")
}

func (d *Database) AddTable(model interface{}, name string) {
	d.Tables[name] = model
}

func (d *Database) SetDriver(db *gorm.DB) {
	d.Driver = db
}

// https://gosamples.dev/sqlite-intro/

// NewDatabase creates a new database. It returns a pointer to the database.
// func initDatabaseDriver(db *gorm.DB) *Database {
// 	ansi.PrintInfo("Initializing new database driver...")
// 	DBInstance = &Database{
// 		Driver: db,
// 		Tables: make(map[string]interface{}),
// 	}
// 	return DBInstance
// }

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

	// if DBInstance == nil {
	// 	DBInstance = initDatabaseDriver(db)
	// }

	// for _, table := range tables {
	// 	DBInstance.Tables[database] = table
	// }

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

func userExists(u *models.User) bool {
	if DBInstance == nil {
		ansi.PrintError("Database driver is nil")
		return false
	}

	var user models.User
	tx := DBInstance.Driver.Where("username = ?", u.Username).First(&user)
	if tx.Error != nil {
		ansi.PrintError(tx.Error.Error())
		return false
	}

	if tx.RowsAffected == 0 {
		ansi.PrintWarning("User does not exist")
		return false
	}

	ansi.PrintSuccess("User exists!")
	return true
}

func writeUser(u *models.User) error {
	fmt.Println("Creating a new user...")
	if DBInstance == nil {
		ansi.PrintError("Database driver is nil")
		return fmt.Errorf("database driver is nil")
	}

	if userExists(u) {
		ansi.PrintWarning("User already exists")
		return fmt.Errorf("user already exists")
	}

	tx := DBInstance.Driver.FirstOrCreate(u) // Create a new user
	if tx.Error != nil || tx.RowsAffected == 0 {
		if tx.Error != nil {
			ansi.PrintError(tx.Error.Error())
		}
		ansi.PrintWarning("User not created")
		ansi.PrintBold("Affected rows: " + fmt.Sprintf("%d", tx.RowsAffected))
		ansi.PrintWarning(tx.Statement.Explain(tx.Statement.SQL.String(), tx.Statement.Vars...))

		return tx.Error
	}

	ansi.PrintSuccess("User created successfully!")
	return nil
}
