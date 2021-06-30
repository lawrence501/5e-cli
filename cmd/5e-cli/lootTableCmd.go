package main

import (
	"log"
	"math/rand"
)

var lowGold = func() error {
	amount := rand.Intn(20) + 1
	log.Printf("Low gold: %dgp\n", amount)
	return nil
}

var trap = func() error {
	t := "standard"
	if rand.Intn(100) < 5 {
		t = "crit"
	}

	var traps []Generic
	traps, err := fetchTraps(t)
	if err != nil {
		return err
	}

	chosenTrap := traps[rand.Intn(len(traps))]
	log.Printf("Trap!\n%s: %s", chosenTrap.Name, chosenTrap.Description)
	return nil
}
