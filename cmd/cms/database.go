package cms

import (
	"github.com/pynezz/pynezzentials/ansi"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const dbPath = "main.db" // Testing purposes. This should be in the config file

func testDbConnection() bool {
	ansi.PrintColor(ansi.DarkYellow, "🔌 Checking database connection... ")
	conf := gorm.Config{
		PrepareStmt: true,
	}

	// Ping the database
	usersDb, err := gorm.Open(sqlite.Open(dbPath), &conf)
	if err != nil {
		ansi.PrintError(err.Error())
		return false
	}
	sqlObj, err := usersDb.DB()
	if err != nil {
		ansi.PrintError(err.Error())
		return false
	}

	if err := sqlObj.Ping(); err != nil {
		ansi.PrintError(err.Error())
		return false
	}

	ansi.PrintColorBold(ansi.LightGreen, "🎉 Database connected!")
	return true
}
