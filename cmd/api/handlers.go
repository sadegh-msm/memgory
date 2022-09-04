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
	dbname := c.QueryParam("name")
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

	err = cmd.del(req.DbName, req.Key)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Err:  true,
			Data: err,
		})
	}

	return c.JSON(http.StatusOK, Response{
		Err:     false,
		Message: fmt.Sprintf("deleted %s with value %s", req.Key, cmd.Container.CurrentDatabase.Data[req.Key]),
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
	storageName := c.QueryParam("name")

	log.Printf("%v\n", storageName)

	dbNames := cmd.listDBs()

	return c.JSON(http.StatusOK, Response{
		Err:     false,
		Message: fmt.Sprintf("all results from %s", storageName),
		Data:    dbNames,
	})
}

func (cmd *Command) ListData(c echo.Context) error {
	databaseName := c.QueryParam("name")
	log.Printf("%v\n", databaseName)

	dbData, err := cmd.listData(databaseName)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Err:  true,
			Data: err,
		})
	}

	return c.JSON(http.StatusOK, Response{
		Err:     false,
		Message: fmt.Sprintf("data from %s database", databaseName),
		Data:    dbData,
	})
}

func (cmd *Command) Load(c echo.Context) error {
	type request struct {
		FilePath string `json:"pattern"`
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

	err = cmd.load(req.FilePath)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Err:  true,
			Data: err,
		})
	}

	return c.JSON(http.StatusAccepted, Response{
		Err:     false,
		Message: fmt.Sprintf("file loaded from %s", req.FilePath),
	})
}

func (cmd *Command) Save(c echo.Context) error {
	type request struct {
		FilePath string `json:"pattern"`
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

	err = cmd.save(req.FilePath)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Err:  true,
			Data: err,
		})
	}

	return c.JSON(http.StatusAccepted, Response{
		Err:     false,
		Message: fmt.Sprintf("file saved in %s", req.FilePath),
	})
}
