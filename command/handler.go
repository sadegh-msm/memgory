package command

import (
	"errors"
	"fmt"
	"os"
	"simple-redis/db"
	"strings"
)

// handling commands

func (cmd *Command) Serve(command string) (interface{}, error) {
	cmdStr := strings.Split(command, " ")[0]

	switch cmdStr {
	case "set":
		if cmd.Container.CurrentDatabase == nil {
			return "", errors.New("no database selected")
		}
		key := strings.Split(command, " ")[1]
		value := strings.Split(command, " ")[2]
		returnedValue := cmd.Container.CurrentDatabase.Set(key, value)

		return returnedValue, nil

	case "get":
		if cmd.Container.CurrentDatabase == nil {
			return "", errors.New("no database selected")
		}
		key := strings.Split(command, " ")[1]
		returnedValue, err := cmd.Container.CurrentDatabase.Get(key)
		if err != nil {
			return "", err
		}

		return returnedValue, nil

	case "del":
		if cmd.Container.CurrentDatabase == nil {
			return "", errors.New("no database selected")
		}
		key := strings.Split(command, " ")[1]
		err := cmd.Container.CurrentDatabase.Delete(key)
		if err != nil {
			return "", err
		}
		return "", nil

	case "keys":
		if cmd.Container.CurrentDatabase == nil {
			return "", errors.New("no database selected")
		}
		pattern := strings.Split(command, " ")[1]
		keys, err := cmd.Container.CurrentDatabase.Regex(pattern)
		if err != nil {
			return "", err
		}

		return keys, nil

	case "use":
		var err error
		dbName := strings.Split(command, " ")[1]

		cmd.Container.CurrentDatabase = cmd.Container.GetDb(dbName)
		if err != nil {
			return "", err
		}

		return "", nil

	case "list":
		var databaseList []string

		dataBases := cmd.Container.ListAllDatabases()
		for _, dbase := range dataBases {
			databaseList = append(databaseList, dbase.Name)
		}
		fmt.Println("count: ", cmd.Container.Count)
		return databaseList, nil

	case "dump":
		if cmd.Container.CurrentDatabase == nil {
			return "", errors.New("no database selected")
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

	case "load":
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

	case "exit":
		db.Exit()

	default:
		return "", errors.New("invalid command")
	}

	return "", nil
}
