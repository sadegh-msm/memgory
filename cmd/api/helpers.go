package api

import (
	"errors"
	"fmt"
	"os"
	"simple-redis/db"
	"strings"
)

var NoDatabaseSelected = errors.New("no database selected")

func (cmd *Command) checkDB() (interface{}, error) {
	if cmd.Container.CurrentDatabase == nil {
		return "", NoDatabaseSelected
	}
	return nil, nil
}

func (cmd *Command) set(databaseName, key, value string) (interface{}, error) {
	cmd.use(databaseName)

	data, err := cmd.checkDB()
	if err != nil {
		return data, err
	}

	returnedValue := cmd.Container.CurrentDatabase.Set(key, value)

	return returnedValue, nil
}

func (cmd *Command) get(command string) (interface{}, error) {
	data, err := cmd.checkDB()
	if err != nil {
		return data, err
	}

	key := strings.Split(command, " ")[1]
	returnedValue, err := cmd.Container.CurrentDatabase.Get(key)
	if err != nil {
		return "", err
	}

	return returnedValue, nil
}

func (cmd *Command) del(command string) (interface{}, error) {
	data, err := cmd.checkDB()
	if err != nil {
		return data, err
	}

	key := strings.Split(command, " ")[1]
	err = cmd.Container.CurrentDatabase.Delete(key)
	if err != nil {
		return "", err
	}

	return "", nil
}

func (cmd *Command) keys(command string) (interface{}, error) {
	data, err := cmd.checkDB()
	if err != nil {
		return data, err
	}

	pattern := strings.Split(command, " ")[1]
	keys, err := cmd.Container.CurrentDatabase.Regex(pattern)
	if err != nil {
		return "", err
	}

	return keys, nil
}

func (cmd *Command) use(dbName string) (interface{}, error) {
	cmd.Container.CurrentDatabase = cmd.Container.GetDb(dbName)
	if err != nil {
		return "", err
	}

	return "", nil
}

func (cmd *Command) list() (interface{}, error) {
	var databaseList []string

	dataBases := cmd.Container.ListAllDatabases()
	for _, dbase := range dataBases {
		databaseList = append(databaseList, dbase.Name)
	}
	fmt.Println("count: ", cmd.Container.Count)

	return databaseList, nil
}

func (cmd *Command) save(command string) (interface{}, error) {
	data, err := cmd.checkDB()
	if err != nil {
		return data, err
	}

	filePath := strings.Split(command, " ")[1]
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

func (cmd *Command) load(command string) (interface{}, error) {
	filePath := strings.Split(command, " ")[1]
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
