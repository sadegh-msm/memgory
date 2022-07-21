package main

import (
	"fmt"
	"simple-redis/command"
	"simple-redis/databeses"
)

func main() {
	db := databeses.New()
	cmdParser := command.New(db)

	err := cmdParser.CmdHandler()
	if err != nil {
		fmt.Println(err)
	}
}
