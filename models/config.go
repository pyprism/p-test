package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	var err error
	db, err = gorm.Open(sqlite.Open("db.sqlite"))
	if err != nil {
		panic("sqlite failed to create")
	}

	db.AutoMigrate(User{}, TagRelation{})

}

func Db() *gorm.DB {
	return db
}
