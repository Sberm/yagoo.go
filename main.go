package main

import (
	"fmt"
	"log"
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

	/* read csv ,  preprocess data(into inverted index).*/
	// csvPath := "resources/wukong_test.csv"
	// storage.PreProcess(csvPath)

	/* iterate and print leveldb */
	// storage.TestPrintDB("./data")

	/* get a certain database id*/
	// str := test.GetDBValue(1)
	// fmt.Println(str)

	/* test Leveldb value */
	// test.TestPrintValue("./data", "连衣裙")

	/*search for a token*/
	e := search.NewEngine("./data", 30)
	var token string
	fmt.Println("Enter what you what to search:")
	fmt.Scan(&token)
	e.SearchCutWord(token)

	return
}
