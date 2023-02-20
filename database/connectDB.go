package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDataBase(user string, password string, dbname string) (*gorm.DB, error) { // capital the first letter to make it public
	dsn := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, password, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}
