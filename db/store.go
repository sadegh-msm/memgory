package db

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"sync"
)

type Store struct {
	Name string
	Data map[string]interface{}
	Lock *sync.RWMutex
}

// creating new db

func NewDb(name string) *Store {
	return &Store{
		Name: name,
		Data: make(map[string]interface{}),
	}
}

// setting and getting and deleting data in db and using lock to stop race conditions

func (store *Store) Set(key string, any interface{}) string {
	//	store.Lock.Lock()
	//	defer store.Lock.Unlock()

	store.Data[key] = any

	return ""
}

func (store *Store) Get(key string) (value interface{}, err error) {
	//store.Lock.RLock()
	//defer store.Lock.RUnlock()

	if value, exist := store.Data[key]; exist {
		return value, nil
	}

	return "", errors.New(fmt.Sprintf("Given key:%s doesn't exist", key))
}

func (store *Store) Delete(key string) error {
	//store.Lock.Lock()
	//defer store.Lock.Unlock()

	if _, exist := store.Data[key]; exist {
		delete(store.Data, key)
		return nil
	}

	return errors.New(fmt.Sprintf("Given key:%s doesn't exist", key))
}

// finding matches with regular expressions and returning a list of matches

func (store *Store) Regex(character string) (keys []string, err error) {
	//store.Lock.Lock()
	//defer store.Lock.Unlock()

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
