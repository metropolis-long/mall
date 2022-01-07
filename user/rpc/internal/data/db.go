package data

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func init() {
	db, err := gorm.Open("mysql", "root:a123456A@(127.0.0.1:3306)/mall?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	fmt.Println("db is ok  ,,, ")
	db.DB().SetMaxIdleConns(10)
    db.DB().SetMaxOpenConns(100)

	if !db.HasTable(&Like{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Like{}).Error; err != nil {
			panic(err)
		}
	}
	Db=db
	// defer Db.Close()
}
