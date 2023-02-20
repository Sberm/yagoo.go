package main

import (
	"log"

	"yagoo/database"
	"yagoo/search"

	"gopkg.in/ini.v1"
)

var (
	user     string
	password string
	dbname   string
)

func loadIniFile(initPath string) {
	cfg, err := ini.Load(initPath)
	if err != nil {
		log.Fatal("Fail to read file: "+initPath, " ", err)
	}
	user = cfg.Section("database").Key("user").String()
	password = cfg.Section("database").Key("password").String()
	dbname = cfg.Section("database").Key("dbname").String()
}

func main() {
	initPath := "my.ini"
	loadIniFile(initPath)

	// create db instance
	db, err := database.ConnectDataBase(user, password, dbname)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	searchText := "欢迎光临会员制餐厅"
	search.SearchCutWord(searchText, db)

	// var wukong_test database.Wukong_data
	// db.First(&wukong_test, 3)
	// fmt.Println(wukong_test)

	return
}
