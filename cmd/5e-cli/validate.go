package main

import (
	"errors"
	"regexp"
	"strconv"
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

var validateTagString = func(input string) error {
	validator := regexp.MustCompile(`^\S*[^,]$`)
	if input != "" && !validator.MatchString(input) {
		return errors.New("Invalid tag list. Must be a comma-separated list of tags, with no spaces.")
	}
	return nil
}

var validateBasicRing = func(input string) error {
	if valid := sliceContains(RING_TAGS, input); !valid {
		return errors.New("Invalid basic ring.")
	}
	return nil
}

var validateMutationType = func(input string) error {
	if valid := sliceContains(MUTATION_TYPES, input); !valid {
		return errors.New("Invalid mutation type.")
	}
	return nil
}

var validateSpaceSeparated = func(input string) error {
	validator := regexp.MustCompile(`^[\w\s]+$`)
	if !validator.MatchString(input) {
		return errors.New("Invalid space-separated string.")
	}
	return nil
}

var validateInt = func(input string) error {
	if _, err := strconv.Atoi(input); err != nil {
		return errors.New("Invalid integer input.")
	}
	return nil
}
