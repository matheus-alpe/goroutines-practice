package main

import (
	"errors"
	"fmt"
	"os"
)

type Runner interface {
    Init([]string) error
    Run() error
    Name() string
}

func root(args []string) error {
    if len(args) < 1 {
	return errors.New("you must pass a sub-command")
    }
    
    cmds := []Runner{
	NewGreetCommand(),
    }

    subcommand := os.Args[1]

    for _, cmd := range cmds {
	if cmd.Name() == subcommand {
	    cmd.Init(os.Args[2:])
	    return cmd.Run()
	}
    }

    return fmt.Errorf("unknown sub-command: %s", subcommand)
}
