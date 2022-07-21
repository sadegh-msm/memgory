package command

import (
	"bufio"
	"fmt"
	"os"
	"simple-redis/databeses"
	"strings"
)

type Command struct {
	Container *databeses.Databases
	Reader    *bufio.Reader
}

func New(dbs *databeses.Databases) *Command {
	return &Command{
		Container: dbs,
		Reader:    bufio.NewReader(os.Stdin),
	}
}

// infinite for loop for reading and serving the commands

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
