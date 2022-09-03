package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Request struct {
	DbName string `json:"dbName"`
	Data   string `json:"data"`
}

type Response struct {
	Err     bool   `json:"err"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func Set(c echo.Context) error {
	req := Request{}

	err := c.Bind(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Err:     true,
			Message: "bad request",
		})
	}

}

func Get(c echo.Context) error {
	return nil
}

func Del(c echo.Context) error {
	return nil
}

func Keys(c echo.Context) error {
	return nil
}

func List(c echo.Context) error {
	return nil
}

func Load(c echo.Context) error {
	return nil
}

func Save(c echo.Context) error {
	return nil
}
