package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pynezz/pynezz_com/internal/server/middleware/models"
	"github.com/pynezz/pynezzentials/ansi"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	// database names
	Content = "content"
	Main    = "main"
)

var dbNames = map[string]string{
	Content: "content.db",
	Main:    "main.db", // The main database
	// Remember to add new databases here if needed in the future
}

// Database defines the structure of the database. We're using SQLite in our project.
type Database struct {
	Tables map[string]interface{}
	Driver *gorm.DB
}

type IContentsDB interface {
	GetPostsMetadata(limit int) ([]models.PostMetadata, error)
	GetPosts(limit int) models.Posts

	GenerateSlug(title string) string
}

type IMainDB interface {
	GetUser(token, username string) (string, uint)
}

var DBInstance *Database // The global database instance
var ContentsDB *Database // The content database instance (globals are bad, I'll fix this later

func init() {
	conf := gorm.Config{
		PrepareStmt: true,
	}
	initContentsDB(conf)
	initMainDB(conf)
}

func initContentsDB(conf gorm.Config) {
	contentsBase, err := InitDB(dbNames[Content], conf, models.PostMetadata{}, models.Post{})
	if err != nil {
		ansi.PrintError(err.Error())
	}

	if err := contentsBase.AutoMigrate(&models.PostMetadata{}, &models.Post{}); err != nil {
		ansi.PrintError(err.Error())
	}

	ContentsDB = &Database{
		Driver: contentsBase,
		Tables: make(map[string]interface{}),
	}

	ContentsDB.SetDriver(contentsBase)
	ContentsDB.AddTable(models.PostMetadata{}, "posts_metadata")
	ContentsDB.AddTable(models.Post{}, "posts")
}

func initMainDB(conf gorm.Config) {
	mainBase, err := InitDB(dbNames[Main], conf, models.User{}, models.Admin{})
	if err != nil {
		ansi.PrintError(err.Error())
	}

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

func isAuthorized(requestedUsername, token string) (valid bool, sameUser bool) {
	// Check if the request is valid
	if token == "" {
		ansi.PrintError("Request is not valid")
		return false, false
	}

	t, err := VerifyJWTToken(token)
	if err != nil || !t.Valid {
		if !t.Valid {
			ansi.PrintError("Token is not valid")
		} else {
			ansi.PrintError(err.Error())
		}

		return false, false
	}

	return true, t.Claims.(jwt.MapClaims)["sub"] == requestedUsername
}

func isValidUser(u *models.User) bool {
	if userExists(u) {
		return true
	}
	return false
}

func GetPosts(limit int) models.Posts {
	var posts models.Posts
	DBInstance.Driver.Limit(limit).Find(&posts)
	return posts
}

// getUser is a helper function that returns a user from the database
// Users are able to fetch all username, but are only able to fetch more information, unless they're an admin.
// The JWT token is used to verify the user's role.
func getUser(token, username string) (string, uint) {

	// quite nasty if statement, but works for now
	if sameUser, valid := isAuthorized(username, token); !valid {
		ansi.PrintError("Request is not valid")
		return "", http.StatusUnauthorized
	} else {
		if isValidUser(&models.User{Username: username}) && !sameUser { // Authorized users are able to fetch any username
			// Not sure if it's the right status code, but it makes it easier to differentiate between
			// a valid returned full user and an partially unauthorized request, but still valid
			return username, http.StatusAccepted
		}
	}

	// user is authorized to fetch the full user data of the requested user
	var user models.User
	DBInstance.Driver.Where("username = ?", username).First(&user)

	userJson, err := json.Marshal(user)
	if err != nil {
		ansi.PrintError(err.Error())
		return "", http.StatusInternalServerError
	}

	return string(userJson), http.StatusOK
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

	tx := DBInstance.Driver.Create(&u) // Create a new user
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
