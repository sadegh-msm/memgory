package databeses

import (
	"errors"
	"fmt"
	"simple-redis/db"
)

type Storage struct {
	AllDatabases    []*db.Database
	CurrentDatabase *db.Database
	Count           int
}

// TODO: add default

// NewStorage creates a new databases for all databases in one container
func NewStorage() *Storage {
	defaultDb := db.NewDb("default")

	storage := &Storage{
		AllDatabases:    nil,
		CurrentDatabase: defaultDb,
		Count:           1,
	}

	storage.AddDatabase("default")

	return storage
}

// AddDatabase checks if database name is valid or not and if not it creates database
func (dbs *Storage) AddDatabase(name string) (*db.Database, error) {
	if dbs.CheckDBExists(name) {
		return nil, errors.New(fmt.Sprintf("this database %s is already created", name))
	}

	database := db.NewDb(name)
	dbs.AllDatabases = append(dbs.AllDatabases, database)
	dbs.Count++

	return database, nil
}

// ListAllDatabases returns all databases names in a container
func (dbs *Storage) ListAllDatabases() []*db.Database {
	return dbs.AllDatabases
}

func (dbs *Storage) GetDb(name string) *db.Database {
	var flag bool
	var findedDb *db.Database

	for _, database := range dbs.AllDatabases {
		if database.Name == name {
			flag = true
			findedDb = database
		}
	}

	if flag {
		return findedDb
	}
	// 	 database doesn't exist, create a new one
	newDb := db.NewDb(name)
	dbs.AllDatabases = append(dbs.AllDatabases, newDb)
	return newDb
}

// this will add default database to databases list

func (dbs *Storage) CreateDefaultDatabase() {
	database := db.NewDb("default")
	dbs.AllDatabases = append(dbs.AllDatabases, database)
	dbs.Count++
}

func (dbs *Storage) CheckDBExists(name string) bool {
	for _, database := range dbs.AllDatabases {
		if database.Name == name {
			return true
		}
	}

	return false
}
