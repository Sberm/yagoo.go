package main

import (
	"encoding/csv"
	"log"
	"os"
	"fmt"
)
func readCsvFile(filePath string) [][]string{
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file " + filePath, err)
	}
	defer f.Close()
	r := csv.NewReader(f)

	// use make to allocate a slice
	records := make([][]string, 20)

	for i := range records {
		records[i] = make([]string, 20)
	}

	//i := 0 
	//records[i], err = r.Read()
	//i++ // can't write in square brackets
	//records[i], err = r.Read()

	for i := 0;i < 20;i++{
		records[i], err = r.Read()
	}


	if err != nil {
		log.Fatal("Unable to parse file as CSV for " + filePath, err)
	}

	return records
}

func printCsv(records [][]string) {
	n := len(records)
	fmt.Println("n = ",n)
	for key, record := range records {
		if key == 0 {continue}
		if key >= n {break}
		for _, value := range record {
			fmt.Println(value, " ")
		}
	}
}

func main() {
	records := readCsvFile("resources/wukong_test.csv")
	printCsv(records)
	//fmt.Println(records)
}
