package storage

import (
	"fmt"
	"testing"
)

func TestLevelDB(t *testing.T) {
	path := "/Users/zhuhaowei/Desktop/Cabin/yagoo/data"
	timeOut := 30 // timeOut set for 30s
	db, err := Init(path, int64(timeOut))
	if err != nil {
		t.Fatal("Error :", err)
	}

	err = db.Put([]byte("Craig"), []byte("Jones"))
	if err != nil {
		t.Fatal("Error:", err)
	}

	value, err := db.Get([]byte("Craig"))
	if err != nil {
		t.Fatal("Error:", err)
	}

	fmt.Println("value:", string(value))

}
