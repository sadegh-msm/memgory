package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"simple-redis/cmd/api"
)

const port = ":8080"

func main() {
	e := echo.New()

	log.Fatal(api.NewRouter(e).Start(port))
}
