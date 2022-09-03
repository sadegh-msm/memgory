package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"simple-redis/cmd/api"
	"simple-redis/cmd/databeses"
)

func main() {
	e := echo.New()
	api.NewRouter(e)

	db := databeses.NewDB()
	cmdParser := api.NewCmd(db)

	err := cmdParser.CmdHandler()
	if err != nil {
		fmt.Println(err)
	}
}
