package api

import (
	"errors"
	"fmt"
	"os"
	"simple-redis/db"
)

var LastSetDB string

var NoDatabaseSelected = errors.New("no database selected")

// checks if database pointer is nil or not
func (cmd *Command) checkDB() error {
	if cmd.Container.CurrentDatabase == nil {
		return NoDatabaseSelected
	}

	return nil
}

// setting new data in database by database name, key and value
// if the name of database is empty string ("") will do the operation on last set database
// if no database have been set it will do operation on default database
func (cmd *Command) set(databaseName, key, value string) (interface{}, error) {
	if databaseName != "" {
		cmd.use(databaseName)
	}

	err := cmd.checkDB()
	if err != nil {
		return nil, err
	}

	returnedValue := cmd.Container.CurrentDatabase.Set(key, value)

	return returnedValue, nil
}

// getting new data in database by database name, key
// if the name of database is empty string ("") will do the operation on last set database
// if no database have been set it will do operation on default database
func (cmd *Command) get(databaseName, key string) (interface{}, error) {
	if databaseName != "" {
		cmd.use(databaseName)
	} else {
		cmd.use(LastSetDB)
	}

	err := cmd.checkDB()
	if err != nil {
		return nil, err
	}

	returnedValue, err := cmd.Container.CurrentDatabase.Get(key)
	if err != nil {
		return nil, err
	}

	return returnedValue, nil
}

// deleting new data in database by database name, key
// if the name of database is empty string ("") will do the operation on last set database
// if no database have been set it will do operation on default database
func (cmd *Command) del(databaseName, key string) error {
	if databaseName != "" {
		cmd.use(databaseName)
	} else {
		cmd.use(LastSetDB)
	}

	err := cmd.checkDB()
	if err != nil {
		return err
	}

	err = cmd.Container.CurrentDatabase.Delete(key)
	if err != nil {
		return err
	}

	return nil
}

// searching for data in database by database name, pattern
// search with the pattern through all data in a database and finds the matching data
// if the name of database is empty string ("") will do the operation on last set database
// if no database have been set it will do operation on default database
func (cmd *Command) keyRegex(databaseName, pattern string) (interface{}, error) {
	if databaseName != "" {
		cmd.use(databaseName)
	} else {
		cmd.use(LastSetDB)
	}

	err := cmd.checkDB()
	if err != nil {
		return nil, err
	}

	keys, err := cmd.Container.CurrentDatabase.Regex(pattern)
	if err != nil {
		return "", err
	}

	return keys, nil
}

// using new data in database by database name, key and value
// if the name of database is empty string ("") will do the operation on last set database
// if no database have been set it will do operation on default database
func (cmd *Command) use(dbName string) {
	cmd.Container.CurrentDatabase = cmd.Container.GetDb(dbName)
	LastSetDB = cmd.Container.GetDb(dbName).Name
}

// will list all databases from a container
func (cmd *Command) listDBs() []string {
	var databaseList []string

	dataBases := cmd.Container.ListAllDatabases()
	for _, dbase := range dataBases {
		databaseList = append(databaseList, dbase.Name)
	}
	fmt.Println("count: ", cmd.Container.Count)

	return databaseList
}

// will list all data in one database
func (cmd *Command) listData(databaseName string) ([]string, error) {
	if databaseName != "" {
		cmd.use(databaseName)
	} else {
		cmd.use(LastSetDB)
	}

	err := cmd.checkDB()
	if err != nil {
		return nil, err
	}

	data := cmd.Container.GetDb(databaseName).ListData()

	return data, nil
}

// will save all data to a file for other usages
func (cmd *Command) save(filePath string) error {
	err := cmd.checkDB()
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	err = db.SaveToFile(cmd.Container.CurrentDatabase, file)
	if err != nil {
		return err
	}

	return nil
}

// will load the saved file to the server
func (cmd *Command) load(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	database, err := db.ReadFromFile(file)
	if err != nil {
		return err
	}

	for _, oldDatabase := range cmd.Container.AllDatabases {
		if oldDatabase.Name == database.Name {
			oldDatabase.Data = database.Data
		}
	}
	cmd.Container.CurrentDatabase = database

	return nil
}
