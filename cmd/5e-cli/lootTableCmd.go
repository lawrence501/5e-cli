package main

import (
	"log"

	"github.com/justinian/dice"
	"github.com/manifoldco/promptui"
)

var gold = func() error {
	goldP := promptui.Select{
		Label: "Gold amount",
		Items: []string{"1d1", "1d20", "1d20+50", "1d20+150"},
	}
	_, choice, err := goldP.Run()
	if err != nil {
		return err
	}

	actual, _, err := dice.Roll(choice)
	if err != nil {
		return err
	}

	log.Printf("%dgp\n", actual.Int())
	return nil
}
