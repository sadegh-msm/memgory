package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type Response struct {
	Err     bool        `json:"err"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (cmd *Command) Set(c echo.Context) error {
	type request struct {
		DbName string `json:"dbname"`
		Key    string `json:"key"`
		Value  string `json:"value"`
	}

	req := request{}

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Err:     true,
			Message: "bad request",
			Data:    err,
		})
	}
	log.Printf("%v\n", req)

	_, err = cmd.set(req.DbName, req.Key, req.Value)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Err:     true,
			Message: "bad request",
			Data:    err,
		})
	}

	return c.JSON(http.StatusCreated, Response{
		Err:     false,
		Message: fmt.Sprintf("in %s : %s -> %s", cmd.Container.CurrentDatabase.Name, req.Key, req.Value),
	})
}

func (cmd *Command) Get(c echo.Context) error {
	dbname := c.QueryParam("db")
	key := c.QueryParam("key")

	log.Printf("%v %v\n", dbname, key)

	data, err := cmd.get(dbname, key)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Err:     true,
			Message: fmt.Sprintf("can not find %s key", key),
			Data:    err,
		})
	}

	return c.JSON(http.StatusOK, Response{
		Err:  false,
		Data: data,
	})
}

func (cmd *Command) Del(c echo.Context) error {
	type request struct {
		DbName string `json:"dbname"`
		Key    string `json:"key"`
	}

	req := request{}

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Err:     true,
			Message: "bad request",
		})
	}
	log.Printf("%v\n", req)

	data, err := cmd.del(req.DbName, req.Key)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Err:  true,
			Data: err,
		})
	}

	return c.JSON(http.StatusOK, Response{
		Err:     false,
		Message: fmt.Sprintf("deleted %s with value %s", req.Key, cmd.Container.CurrentDatabase.Data[req.Key]),
		Data:    data,
	})
}

func (cmd *Command) UseDB(c echo.Context) error {
	type request struct {
		DbName string `json:"dbname"`
	}

	req := request{}

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Err:     true,
			Message: "bad request",
		})
	}
	log.Printf("%v\n", req)

	cmd.use(req.DbName)
	return c.JSON(http.StatusOK, Response{
		Err:     false,
		Message: fmt.Sprintf("selected %s", req.DbName),
	})
}

func (cmd *Command) KeyRegex(c echo.Context) error {
	type request struct {
		DbName  string `json:"dbname"`
		Pattern string `json:"pattern"`
	}

	req := request{}

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Err:     true,
			Message: "bad request",
		})
	}
	log.Printf("%v\n", req)

	data, err := cmd.keyRegex(req.DbName, req.Pattern)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Err:  true,
			Data: err,
		})
	}

	return c.JSON(http.StatusOK, Response{
		Err:     false,
		Message: fmt.Sprintf("serched for %s inside %s", req.Pattern, req.DbName),
		Data:    data,
	})
}

func (cmd *Command) ListDBs(c echo.Context) error {
	type request struct {
		StorageName string `json:"dbname"`
	}

	req := request{}

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Err:     true,
			Message: "bad request",
		})
	}
	log.Printf("%v\n", req)

	dbNames := cmd.listDBs()

	return c.JSON(http.StatusOK, Response{
		Err:     false,
		Message: fmt.Sprintf("all results from %s", req.StorageName),
		Data:    dbNames,
	})
}

func (cmd *Command) ListData(c echo.Context) error {
	type request struct {
		DatabaseName string `json:"dbname"`
	}

	req := request{}

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Err:     true,
			Message: "bad request",
		})
	}
	log.Printf("%v\n", req)

	dbData := cmd.listData()
}

func (cmd *Command) Load(c echo.Context) error {
	return nil
}

func (cmd *Command) Save(c echo.Context) error {
	return nil
}
