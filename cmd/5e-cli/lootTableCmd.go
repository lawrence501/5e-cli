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

	chosen := traps[rand.Intn(len(traps))]
	log.Printf("Trap!\n%s: %s", chosen.Name, chosen.Description)
	return nil
}

var mundane = func() error {
	t := "standard"
	if rand.Intn(100) < 5 {
		t = "crit"
	}

	var mundanes []Generic
	mundanes, err := fetchMundanes(t)
	if err != nil {
		return err
	}

	chosen := mundanes[rand.Intn(len(mundanes))]
	log.Printf("Mundane\n%s: %s", chosen.Name, chosen.Description)
	return nil
}
