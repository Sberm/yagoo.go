package util

import (
	"bytes"
	"encoding/gob"
)

func Encode(original []uint32) ([]byte, error) {
	var dbData bytes.Buffer
	enc := gob.NewEncoder(&dbData)
	err := enc.Encode(original)
	if err != nil {
		return nil, err
	}
	return dbData.Bytes(), nil
}

func Decode(code []byte, v interface{}) error {
	buffer := bytes.NewBuffer(code)
	dec := gob.NewDecoder(buffer)
	err := dec.Decode(v)
	if err != nil {
		return err
	}
	return nil
}
