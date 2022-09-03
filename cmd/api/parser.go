package api

import (
	"bufio"
	"fmt"
	"os"
	"simple-redis/cmd/databeses"
	"strings"
)

type Command struct {
	Container *databeses.Storage
	Reader    *bufio.Reader
}

func NewCmd(dbs *databeses.Storage) *Command {
	return &Command{
		Container: dbs,
		Reader:    bufio.NewReader(os.Stdin),
	}
}

// CmdHandler infinite for loop for reading and serving the commands
func (cmd *Command) CmdHandler() error {
	cmd.Container.CreateDefaultDatabase()
	for {
		input, err := cmd.Reader.ReadString('\n')
		if err != nil {
			return err
		}

		val, err := cmd.Serve(strings.TrimSuffix(input, "\n"))
		if err != nil {
			return err
		}
		if val != "" && val != nil {
			fmt.Println("> ", val)
		}
	}
}
