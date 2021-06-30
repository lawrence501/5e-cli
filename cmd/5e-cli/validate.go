package main

import (
	"errors"
)

var validateBase = func(input string) error {
	if _, valid := COMMAND_MAP[input]; !valid {
		return errors.New("Invalid command.")
	}
	return nil
}

var validateColourUpgrade = func(input string) error {
	if _, valid := COLOUR_UPGRADE_DESCRIPTIONS[input]; !valid {
		return errors.New("Invalid colour upgrade target. Must either be a loot roll or a crafting stone.")
	}
	return nil
}
