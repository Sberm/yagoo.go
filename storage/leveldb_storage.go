package storage

import (
	"log"
	"sync"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
)

type LevelDB_storage struct {
	path       string
	db         *leveldb.DB
	lastUpdate int64
	timeOut    int64
	closed     bool
	Mu         sync.Mutex
}

func Init(path string, timeOut int64) (*LevelDB_storage, error) {
	db, err := OpenDB(path)
	if err != nil {
		log.Fatal("err:", err)
	}
	s := &LevelDB_storage{
		path:       path,
		db:         db,
		lastUpdate: time.Now().Unix(),
		timeOut:    timeOut,
		closed:     true,
	}
	// preprocessing, don't check timeout
	// go s.checkTimeOut()
	return s, err
}

func (s *LevelDB_storage) checkTimeOut() {
	for {
		n := time.Now().Unix() // returns int64
		if !s.closed && n-s.lastUpdate > s.timeOut {
			// s.db.Close() // use interface to open, close and set, get the db.
			err := s.closeDB()
			if err != nil {
				log.Fatal("closing leveldb failed:", err)
			}
			s.closed = true
		}
		time.Sleep(time.Second * 7)
	}
}

func (s *LevelDB_storage) reOpenDB() {
	if s.closed {
		db, err := OpenDB(s.path)
		if err != nil {
			log.Fatal("err:", err)
		}
		s.db = db
		s.closed = false
	}
}

// public method to test
func OpenDB(path string) (*leveldb.DB, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}
	return db, err
}

func (s *LevelDB_storage) closeDB() error {
	err := s.db.Close()
	return err
}

func (s *LevelDB_storage) Put(key []byte, value []byte) error {
	return s.db.Put(key, value, nil)
}

func (s *LevelDB_storage) Get(key []byte) ([]byte, error) {
	val, err := s.db.Get(key, nil)
	if err != nil {
		return nil, err
	}
	return val, err
}
