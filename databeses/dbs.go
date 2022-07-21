package databeses

import (
	"errors"
	"fmt"
	"simple-redis/db"
)

type Databases struct {
	AllDatabases    []*db.Store
	CurrentDatabase *db.Store
	Count           int
}

// creates a new databases for all databases in one container

func New() *Databases {
	return &Databases{
		AllDatabases:    nil,
		CurrentDatabase: nil,
		Count:           0,
	}
}

// checks if database name is valid or not and if not it creates database

func (dbs *Databases) AddDatabase(name string) (*db.Store, error) {
	for _, database := range dbs.AllDatabases {
		if database.Name == name {
			return nil, errors.New(fmt.Sprintf("this database %s is already created", name))
		} else {
			database := db.NewDb(name)
			dbs.AllDatabases = append(dbs.AllDatabases, database)
			dbs.Count++

			return database, nil
		}
	}

	return nil, nil
}

// returns all databases names in a container

func (dbs *Databases) ListAllDatabases() []*db.Store {
	return dbs.AllDatabases
}

func (dbs *Databases) GetDb(name string) *db.Store {
	var flag bool
	var findedDb *db.Store

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

func (dbs *Databases) CreateDefaultDatabase() {
	database := db.NewDb("default")
	dbs.AllDatabases = append(dbs.AllDatabases, database)
	dbs.Count++
}
