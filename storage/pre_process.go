package storage

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"yagoo/util"
)

// pre-process the data, store the inverted index into leveldb
func PreProcess(path string) {
	// read csv, id and description, id starts from 1
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)

	index := uint32(0)

	s, err := Init("./data", 30)
	if err != nil {
		log.Fatal(err)
	}

	tr := util.NewTokenizer()

	wordCount := 0
	for {
		// to do: skip first line
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if index == 0 {
			index++
			continue
		}
		if err != nil {
			log.Fatal(err)
		}
		descrip := record[1]
		fmt.Println("description:", descrip)
		words := tr.CutWord(descrip)

		// remove dups in the one string (for example, two "的")
		words = util.RemoveDuplicate[string](words)

		/*
			fmt.Println("cut words:")
			for _, word := range words {
				fmt.Print(word, " ")
			}
		*/

		for _, word := range words {

			s.Mu.Lock()
			ids, err := s.Get([]byte(word))

			if err != nil {
				fmt.Println(err)
				wordCount++
				fmt.Println("word count:", wordCount)
				var ids []uint32
				ids = append(ids, index)
				ids_enc, err := util.Encode(ids)
				err = s.Put([]byte(word), ids_enc)
				if err != nil {
					log.Fatal(err)
				}
			} else {
				var orig []uint32
				util.Decode(ids, &orig)

				fmt.Println("matches, length of ids:", len(orig))
				orig = append(orig, index)
				ids_enc, err := util.Encode(orig)
				if err != nil {
					log.Fatal(err)
				}
				s.Put([]byte(word), ids_enc)
			}
			s.Mu.Unlock()
		}
		index++
	}
}
