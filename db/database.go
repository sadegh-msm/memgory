package db

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"sync"
)

type Database struct {
	Name string
	Data map[string]interface{}
	*sync.RWMutex
}

// NewDb creating new db
func NewDb(name string) *Database {
	return &Database{
		Name: name,
		Data: make(map[string]interface{}),
	}
}

// Set setting and getting and deleting data in db and using lock to stop race conditions
func (db *Database) Set(key string, any interface{}) string {
	//db.Lock()
	//defer db.Unlock()

	db.Data[key] = any

	return ""
}

func (db *Database) Get(key string) (value interface{}, err error) {
	//db.RLock()
	//defer db.RUnlock()

	if value, exist := db.Data[key]; exist {
		return value, nil
	}

	return "", errors.New(fmt.Sprintf("Given key:%s doesn't exist", key))
}

func (db *Database) Delete(key string) error {
	//db.Lock()
	//defer db.Unlock()

	if _, exist := db.Data[key]; exist {
		delete(db.Data, key)
		return nil
	}

	return errors.New(fmt.Sprintf("Given key:%s doesn't exist", key))
}

// Regex finding matches with regular expressions and returning a list of matches
func (db *Database) Regex(character string) (keys []string, err error) {
	//store.Lock()
	//defer store.Unlock()

	for key := range db.Data {
		matchString, err := regexp.MatchString(character, key)
		if err != nil {
			return nil, err
		}

		if matchString {
			keys = append(keys, key)
		}
	}

	return keys, nil
}

// ListData lists all data from a database to a slice of string
func (db *Database) ListData() (values []string) {
	for key := range db.Data {
		values = append(values, fmt.Sprintf("%s -> %s\n", key, db.Data[key]))
	}

	return values
}

// Exit exiting from application
func Exit() {
	fmt.Println("exiting database !!")
	os.Exit(1)
}
