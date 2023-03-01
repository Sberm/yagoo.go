package search

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"strings"
	"yagoo/database"
	"yagoo/storage"
	"yagoo/util"
)

type Engine struct {
	s *storage.LevelDB_storage
}

type Record struct {
	id    uint32
	count uint64
}

func NewEngine(path string, timeOut int64) *Engine {
	s, err := storage.Init(path, timeOut)
	if err != nil {
		log.Fatal("Start engine failed because failed to open LevelDB, err:", err)
	}
	e := &Engine{
		s: s,
	}
	return e
}

func (e *Engine) SearchCutWord(text string) {
	tr := util.NewTokenizer()
	words := tr.CutWord(text)
	fmt.Println("text :", text)

	words = util.RemoveDuplicate[string](words)

	fmt.Println("searching :", strings.Join(words, "|"))

	records := make([]Record, 0)
	record_map := make(map[uint32]int, 0)

	for _, word := range words {
		ids_orig, err := e.s.Get([]byte(word))
		if err != nil {
			fmt.Printf("Search for word: %s, No records found. (%s)\n", word, err)
		}
		ids := make([]uint32, 0)

		util.Decode(ids_orig, &ids)
		for _, id := range ids {
			_, exist := record_map[id]
			if !exist {
				r := Record{
					id:    id,
					count: 0,
				}
				records = append(records, r)
				record_map[id] = len(records) - 1
			}
			records[record_map[id]].count++
		}
	}
	sort.Slice(records, func(i, j int) bool {
		return records[i].count > records[j].count
	})

	// use init file later
	db, err := database.ConnectDataBase("root", "112358", "wukongDB")
	if err != nil {
		log.Fatal("Open Mysql failed, err:", err)
	}

	data := database.Wukong_data{}

	fmt.Println("Search result:\n")

	records_count := 0
	if records_count < len(records) {
		records_count = len(records)
	}
	if records_count > 20 {
		records_count = 20
	}

	for i := 0; i < records_count; i++ {
		fmt.Printf("id : %d, count : %d\n", records[i].id, records[i].count)
		// clear data before every time SELECT ing
		data = database.Wukong_data{}
		db.Where("id=?", strconv.FormatUint(uint64(records[i].id), 10)).First(&data)
		fmt.Println("result:\n", data.Text)
	}
}
