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

// creating new db

func NewDb(name string) *Database {
	return &Database{
		Name: name,
		Data: make(map[string]interface{}),
	}
}

// setting and getting and deleting data in db and using lock to stop race conditions

func (store *Database) Set(key string, any interface{}) string {
	//store.Lock()
	//defer store.Unlock()

	store.Data[key] = any

	return ""
}

func (store *Database) Get(key string) (value interface{}, err error) {
	//store.RLock()
	//defer store.RUnlock()

	if value, exist := store.Data[key]; exist {
		return value, nil
	}

	return "", errors.New(fmt.Sprintf("Given key:%s doesn't exist", key))
}

func (store *Database) Delete(key string) error {
	//store.Lock()
	//defer store.Unlock()

	if _, exist := store.Data[key]; exist {
		delete(store.Data, key)
		return nil
	}

	return errors.New(fmt.Sprintf("Given key:%s doesn't exist", key))
}

// finding matches with regular expressions and returning a list of matches

func (store *Database) Regex(character string) (keys []string, err error) {
	//store.Lock()
	//defer store.Unlock()

	for key, _ := range store.Data {
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

// exiting from application

func Exit() {
	fmt.Println("exiting database !!")
	os.Exit(1)
}
