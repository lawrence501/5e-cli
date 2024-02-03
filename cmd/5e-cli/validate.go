package main

import (
	"errors"
	"regexp"
	"strconv"

	"golang.org/x/exp/slices"
)

var validateBase = func(input string) error {
	if validateInt(input) == nil {
		return nil
	}
	if _, valid := COMMAND_MAP[input]; !valid {
		return errors.New("invalid command")
	}
	return nil
}

var validateRingStone = func(input string) error {
	if valid := slices.Contains(RING_STONES, input); !valid {
		return errors.New("invalid ring stone")
	}
	return nil
}

var validateActivity = func(input string) error {
	if valid := slices.Contains(JOURNEY_ACTIVITIES, input); !valid {
		return errors.New("invalid journey activity")
	}
	return nil
}

var validateEncounter = func(input string) error {
	if input == "" {
		return nil
	}
	if valid := slices.Contains(ENCOUNTER_TAGS, input); !valid {
		return errors.New("invalid encounter tag")
	}
	return nil
}

var validateColourUpgrade = func(input string) error {
	if _, valid := COLOUR_UPGRADE_DESCRIPTIONS[input]; !valid {
		return errors.New("invalid colour upgrade target. Must be a loot type")
	}
	return nil
}

var validateGem = func(input string) error {
	if valid := slices.Contains(GEM_TAGS, input); !valid {
		return errors.New("invalid gem tag")
	}
	return nil
}

var validateSpaceSeparated = func(input string) error {
	validator := regexp.MustCompile(`^[\w\s]+$`)
	if !validator.MatchString(input) {
		return errors.New("invalid space-separated string")
	}
	return nil
}

var validateInt = func(input string) error {
	if _, err := strconv.Atoi(input); err != nil {
		return errors.New("invalid integer input")
	}
	return nil
}

var validateFloat = func(input string) error {
	if _, err := strconv.ParseFloat(input, 64); err != nil {
		return errors.New("invalid float input")
	}
	return nil
}

var validateTarotCard = func(input string) error {
	if valid := slices.Contains(TAROT_CARDS, input); !valid {
		return errors.New("invalid tarot card")
	}
	return nil
}

var validatePartyMember = func(input string) error {
	if valid := slices.Contains(PARTY_MEMBERS, input); !valid {
		return errors.New("invalid party member")
	}
	return nil
}
