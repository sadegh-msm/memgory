package api

import (
	"simple-redis/cmd/databeses"
)

type Command struct {
	Container *databeses.Storage
}

func NewCmd(dbs *databeses.Storage) *Command {
	return &Command{
		Container: dbs,
	}
}
