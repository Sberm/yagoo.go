package search

import (
	"fmt"

	"strings"

	"github.com/yanyiwu/gojieba"

	"gorm.io/gorm"
)

func SearchCutWord(text string, db *gorm.DB) {
	use_hmm := true
	x := gojieba.NewJieba()
	defer x.Free()

	words := x.CutForSearch(text, !use_hmm)
	fmt.Println(text)
	fmt.Println("searching:", strings.Join(words, "/"))

	// how to search in database with cut words?
	// is it sql? or is it something else
	// sql not fast enough
	// do the sql first
	// how to do searching with cut words in go but not separated words in database?

	// pre processing: after importing to the database, do preprocess to the database's description column,  save as string separated by ,?(what is the fastest way to search a database using cut words?)
}
