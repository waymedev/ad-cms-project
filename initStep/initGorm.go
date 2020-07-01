package initStep

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)


var DB *gorm.DB

func InitGorm() {
	database, err := gorm.Open("sqlite3", "ad.db")

	if err != nil {
		panic("Failed to connect to sqlite3!")
	}

	DB = database
}

func InitGormWithPath(path string) {
	database, err := gorm.Open("sqlite3", path)

	if err != nil {
		panic("Failed to connect to sqlite3!")
	}

	DB = database
}
