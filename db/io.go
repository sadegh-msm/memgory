package db

import (
	"encoding/gob"
	"io"
)

// writes a database to a file

func SaveToFile(store *Store, writer io.Writer) error {
	encoder := gob.NewEncoder(writer)

	return encoder.Encode(*store)
}

// reads from a file to database

func ReadFromFile(reader io.Reader) (*Store, error) {
	db := NewDb("")

	decoder := gob.NewDecoder(reader)
	if err := decoder.Decode(db); err != nil {
		return nil, err
	}

	return db, nil
}
