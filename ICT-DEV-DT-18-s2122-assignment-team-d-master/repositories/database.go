package repositories

import (
	. "eindopdracht/types"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var _connection *gorm.DB

func connection() *gorm.DB {
	if _connection != nil {
		return _connection
	}

	fmt.Println(os.Getenv("CONNECTION_STRING"))
	dsn := os.Getenv("CONNECTION_STRING")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Group{}, &Student{}, &Module{}, &Grade{}, &Teacher{})

	_connection = db

	return db
}

func Connected() bool {
	return !connection().Config.DryRun
}
