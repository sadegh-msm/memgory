package api

import (
	"errors"
	"fmt"
	"os"
	"simple-redis/db"
)

var LastSetDB string

var NoDatabaseSelected = errors.New("no database selected")

func (cmd *Command) checkDB() (interface{}, error) {
	if cmd.Container.CurrentDatabase == nil {
		return "", NoDatabaseSelected
	}

	return nil, nil
}

func (cmd *Command) set(databaseName, key, value string) (interface{}, error) {
	if databaseName != "" {
		cmd.use(databaseName)
	}

	data, err := cmd.checkDB()
	if err != nil {
		return data, err
	}

	returnedValue := cmd.Container.CurrentDatabase.Set(key, value)

	return returnedValue, nil
}

func (cmd *Command) get(databaseName, key string) (interface{}, error) {
	if databaseName != "" {
		cmd.use(databaseName)
	} else {
		cmd.use(LastSetDB)
	}

	data, err := cmd.checkDB()
	if err != nil {
		return data, err
	}

	returnedValue, err := cmd.Container.CurrentDatabase.Get(key)
	if err != nil {
		return "", err
	}

	return returnedValue, nil
}

func (cmd *Command) del(databaseName, key string) (interface{}, error) {
	if databaseName != "" {
		cmd.use(databaseName)
	} else {
		cmd.use(LastSetDB)
	}

	data, err := cmd.checkDB()
	if err != nil {
		return data, err
	}

	err = cmd.Container.CurrentDatabase.Delete(key)
	if err != nil {
		return "", err
	}

	return "", nil
}

func (cmd *Command) keyRegex(databaseName, pattern string) (interface{}, error) {
	if databaseName != "" {
		cmd.use(databaseName)
	} else {
		cmd.use(LastSetDB)
	}

	data, err := cmd.checkDB()
	if err != nil {
		return data, err
	}

	keys, err := cmd.Container.CurrentDatabase.Regex(pattern)
	if err != nil {
		return "", err
	}

	return keys, nil
}

func (cmd *Command) use(dbName string) {
	cmd.Container.CurrentDatabase = cmd.Container.GetDb(dbName)
	LastSetDB = cmd.Container.GetDb(dbName).Name
}

func (cmd *Command) listDBs() []string {
	var databaseList []string

	dataBases := cmd.Container.ListAllDatabases()
	for _, dbase := range dataBases {
		databaseList = append(databaseList, dbase.Name)
	}
	fmt.Println("count: ", cmd.Container.Count)

	return databaseList
}

func (cmd *Command) listData(databaseName string) ([]string, error) {
	if databaseName != "" {
		cmd.use(databaseName)
	} else {
		cmd.use(LastSetDB)
	}

	_, err := cmd.checkDB()
	if err != nil {
		return nil, err
	}

	data := cmd.Container.GetDb(databaseName).ListData()

	return data, nil
}

func (cmd *Command) save(filePath string) (interface{}, error) {
	data, err := cmd.checkDB()
	if err != nil {
		return data, err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	err = db.SaveToFile(cmd.Container.CurrentDatabase, file)
	if err != nil {
		return "", err
	}

	return "", nil
}

func (cmd *Command) load(filePath string) (interface{}, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	database, err := db.ReadFromFile(file)
	if err != nil {
		return "", err
	}

	for _, oldDatabase := range cmd.Container.AllDatabases {
		if oldDatabase.Name == database.Name {
			oldDatabase.Data = database.Data
		}
	}
	cmd.Container.CurrentDatabase = database

	return "", nil
}
