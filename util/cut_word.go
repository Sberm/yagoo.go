package util

import (
	"github.com/yanyiwu/gojieba"
)

type Tokenizer struct {
	tok *gojieba.Jieba
}

func NewTokenizer() *Tokenizer {
	tr := &Tokenizer{
		tok: gojieba.NewJieba(),
	}
	return tr
}

func (tr *Tokenizer) CutWord(word string) []string {
	words := tr.tok.CutForSearch(word, false)
	return words
}

func RemoveDuplicate[T string | int](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
