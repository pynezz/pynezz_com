package cms

import (
	"github.com/pynezz/pynezz_com/internal/server/middleware"
	"github.com/pynezz/pynezz_com/internal/server/middleware/models"
	"github.com/pynezz/pynezzentials/ansi"
	"gorm.io/gorm"
)

const dbPath = "users.db" // Testing purposes. This should be in the config file

func testDbConnection() bool {
	conf := gorm.Config{
		PrepareStmt: true,
	}

	usersDb, err := middleware.InitDB(dbPath, conf, models.User{})
	if err != nil {
		ansi.PrintError(err.Error())
		return false
	}

	// Migrate the database
	if err := usersDb.AutoMigrate(); err != nil {
		ansi.PrintError(err.Error())
		return false
	}

	ansi.PrintInfo("Testing db connection...")
	if err := usersDb.Exec("SELECT 1").Error; err != nil {
		ansi.PrintError(err.Error())
		return false
	}

	ansi.PrintColorBold(ansi.LightGreen, "🎉 Database connected!")

	return true
}
