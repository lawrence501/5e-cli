package main

import (
	"log"
	"math/rand"
	"strings"
)

var lowGold = func() error {
	amount := rand.Intn(20) + 1
	log.Printf("Low gold: %dgp\n", amount)
	return nil
}

var mediumGold = func() error {
	amount := rand.Intn(20) + 51
	log.Printf("Medium gold: %dgp\n", amount)
	return nil
}

var highGold = func() error {
	amount := rand.Intn(20) + 151
	log.Printf("High gold: %dgp\n", amount)
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

var WONDROUS_WEIGHTS = []int{50, 80, 92, 98, 100}
var WONDROUS_RARITIES = []string{"common", "uncommon", "rare", "very rare", "legendary"}
var wondrous = func() error {
	rarityRoll := rand.Intn(100)
	var rarity string
	for i := range WONDROUS_WEIGHTS {
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

var doubleEnchant = func() error {
	base, err := getMundane("standard")
	if err != nil {
		return err
	}

	enchants, err := getEnchants(2, base.Tags)
	if err != nil {
		return err
	}

	log.Printf("2E magic item\nBase: %s (%s)\n- %s [%spts; %s]\n- %s [%spts; %s]", base.Name, base.Description, enchants[0].Description, enchants[0].PointValue, enchants[0].Upgrade, enchants[1].Description, enchants[1].PointValue, enchants[1].Upgrade)
	return nil
}

var tripleEnchant = func() error {
	base, err := getMundane("standard")
	if err != nil {
		return err
	}

	enchants, err := getEnchants(3, base.Tags)
	if err != nil {
		return err
	}

	log.Printf("3E magic item\nBase: %s (%s)\n- %s [%spts; %s]\n- %s [%spts; %s]\n- %s [%spts; %s]", base.Name, base.Description, enchants[0].Description, enchants[0].PointValue, enchants[0].Upgrade, enchants[1].Description, enchants[1].PointValue, enchants[1].Upgrade, enchants[2].Description, enchants[2].PointValue, enchants[2].Upgrade)
	return nil
}

var essence = func() error {
	dmgType := getDamageType()
	log.Printf("Essence of %s", dmgType)
	return nil
}

var amulet = func() error {
	amulets, err := fetchAmulets()
	if err != nil {
		return err
	}

	chosen := amulets[rand.Intn(len(amulets))]
	modString := strings.Join(chosen.Mods, "\n- ")
	log.Printf("Amulet\n%s:\n- %s", chosen.Name, modString)
	return nil
}

var blessing = func() error {
	blessings, err := fetchBlessings()
	if err != nil {
		return err
	}

	chosen := blessings[rand.Intn(len(blessings))]
	chosen = processMod(chosen)
	log.Printf("Blessing\n%s", chosen)
	return nil
}

var ring = func() error {
	rings, err := fetchBasicRings()
	if err != nil {
		return err
	}

	chosen := rings[rand.Intn(len(rings))]
	chosen.Description = processMod(chosen.Description)
	log.Printf("Basic ring\n%s", chosen.Description)
	return nil
}

var doubleValueSingleEnchant = func() error {
	base, err := getMundane("standard")
	if err != nil {
		return err
	}

	enchants, err := getEnchants(1, base.Tags)
	if err != nil {
		return err
	}

	log.Printf("1E magic item\nBase: %s (%s)\n- %s [%spts; %s] (HARD DOUBLE)", base.Name, base.Description, enchants[0].Description, enchants[0].PointValue, enchants[0].Upgrade)
	return nil
}

var belt = func() error {
	belts, err := fetchGenerics("belt")
	if err != nil {
		return err
	}

	chosen := belts[rand.Intn(len(belts))]
	log.Printf("Belt\n%s", chosen.Description)
	return nil
}

var craftingStone = func() error {
	stones, err := fetchGenerics("craftingStone")
	if err != nil {
		return err
	}

	chosen := stones[rand.Intn(len(stones))]
	log.Printf("Crafting stone\n%s Stone: %s", chosen.Name, chosen.Description)
	return nil
}

var relic = func() error {
	relics, err := fetchRelics()
	if err != nil {
		return err
	}

	chosen := relics[rand.Intn(len(relics))]
	var modDescriptions []string
	for _, m := range chosen.StartingMods {
		m.Description = processMod(m.Description)
		modDescriptions = append(modDescriptions, m.Description)
	}
	modString := strings.Join(modDescriptions, "\n- ")
	log.Printf("Relic\n%s (%s):\n- %s", chosen.Name, chosen.Type, modString)
	return nil
}
