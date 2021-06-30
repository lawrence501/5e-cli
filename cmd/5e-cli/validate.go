package main

import (
	"errors"
)

var validateBase = func(input string) error {
	if _, valid := COMMAND_MAP[input]; !valid {
		return errors.New("Invalid command. Type 'help' for a list of all commands.")
	}
	return nil
}
