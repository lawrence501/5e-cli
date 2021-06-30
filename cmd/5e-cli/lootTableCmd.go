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

	chosen, err := getMundane(t)
	if err != nil {
		return err
	}

	log.Printf("Mundane\n%s: %s", chosen.Name, chosen.Description)
	return nil
}

func getMundane(t string) (Mundane, error) {
	mundanes, err := fetchMundanes(t)
	if err != nil {
		return Mundane{}, err
	}

	return mundanes[rand.Intn(len(mundanes))], nil
}

var WONDROUS_WEIGHTS = []int{50, 80, 92, 98, 100}
var WONDROUS_RARITIES = []string{"common", "uncommon", "rare", "very rare", "legendary"}
var wondrous = func() error {
	rarityRoll := rand.Intn(100)
	var rarity string
	for i, _ := range WONDROUS_WEIGHTS {
		if rarityRoll < WONDROUS_WEIGHTS[i] {
			rarity = WONDROUS_RARITIES[i]
			break
		}
	}

	wondrous, err := fetchWondrous(rarity)
	if err != nil {
		return err
	}

	chosen := wondrous[rand.Intn(len(wondrous))]
	log.Printf("%s wondrous item\n%s: %s", rarity, chosen.Name, chosen.Description)
	return nil
}

var singleEnchant = func() error {
	base, err := getMundane("standard")
	if err != nil {
		return err
	}

	enchants, err := getEnchants(1, base.Tags)
	if err != nil {
		return err
	}

	log.Printf("1E magic item\nBase: %s (%s)\n- %s [%spts; %s]", base.Name, base.Description, enchants[0].Description, enchants[0].PointValue, enchants[0].Upgrade)
	return nil
}

func getEnchants(num int, tags []string) ([]Enchant, error) {
	allEnchants, err := fetchEnchants()
	if err != nil {
		return []Enchant{}, err
	}

	var enchants []Enchant
	for len(enchants) < num {
		var e Enchant
		for true {
			e = allEnchants[rand.Intn(len(allEnchants))]
			valid := true
			for _, t := range e.Tags {
				if !sliceContains(tags, t) {
					valid = false
					break
				}
			}
			if valid {
				break
			}
		}
		enchants = append(enchants, e)
	}

	return enchants, nil
}
