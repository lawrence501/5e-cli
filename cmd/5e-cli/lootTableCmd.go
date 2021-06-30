package main

import (
	"log"
	"math/rand"

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

var trap = func() error {
	fName := "trap"
	if rand.Intn(100) < 5 {
		fName = "critTrap"
	}

	var traps []Generic
	traps, err := fetchGenerics(fName)
	if err != nil {
		return err
	}

	chosenTrap := traps[rand.Intn(len(traps))]
	log.Printf("Trap!\n%s: %s", chosenTrap.Name, chosenTrap.Description)
	return nil
}
