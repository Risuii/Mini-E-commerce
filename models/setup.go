package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDatabase() {
	db, err := gorm.Open(mysql.Open("root:babehlo123@tcp(localhost:3306)/project-golang"))

	if err != nil {
		fmt.Println("ERROR DATABASE CONNECTION")
	}

	db.AutoMigrate(&User{}, &Toko{}, &Barang{})

	DB = db
}
