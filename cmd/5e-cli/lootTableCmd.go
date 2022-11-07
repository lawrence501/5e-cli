package main

import (
	"errors"
	"log"
	"math/rand"
	"strings"

	"github.com/manifoldco/promptui"
)

var tome = func() error {
	critRoll := rand.Intn(100)
	if critRoll < 10 {
		log.Println("Blank tome!")
		return nil
	}
	tomes, err := fetchTomes()
	if err != nil {
		return err
	}
	chosen := tomes[rand.Intn(len(tomes))]
	chosen.Description = processMod(chosen.Description)

	log.Printf("Tome of %s - Can be applied to %s\nEffect: %s", chosen.Name, chosen.Target, chosen.Description)
	return nil
}

var lowGold = func() error {
	amount := rand.Intn(4) + 3
	log.Printf("Low gold: %dgp\n", amount)
	return nil
}

var mediumGold = func() error {
	amount := rand.Intn(4) + 11
	log.Printf("Medium gold: %dgp\n", amount)
	return nil
}

var highGold = func() error {
	amount := rand.Intn(4) + 31
	log.Printf("High gold: %dgp\n", amount)
	return nil
}

var mundane = func() error {
	t := "standard"
	if rand.Intn(100) < 10 {
		t = "crit"
	}

	chosen, err := getMundane(t)
	if err != nil {
		return err
	}

	title := "Mundane"
	if t == "crit" {
		var bonus string
		switch chosen.Tags[0] {
		case "weapon":
			bonus = "+1 dmg die size"
		case "armour":
			bonus = "+1 phys res"
		default:
			bonus = ""
		}
		title = "Exotic mundane (" + bonus + ")!"
	}
	log.Printf("%s\n%s: %s", title, chosen.Name, chosen.Description)
	return nil
}

var WONDROUS_WEIGHTS = []int{50, 80, 92, 98, 100}
var WONDROUS_RARITIES = []string{"gold", "uncommon", "rare", "very rare", "legendary"}
var wondrous = func() error {
	rarityRoll := rand.Intn(100)
	log.Println(rarityRoll)
	var rarity string
	for i := range WONDROUS_WEIGHTS {
		if rarityRoll < WONDROUS_WEIGHTS[i] {
			rarity = WONDROUS_RARITIES[i]
			break
		}
	}

	if rarity == "gold" {
		if err := lowGold(); err != nil {
			return err
		}
		return nil
	}
	wondrous, err := fetchWondrous(rarity)
	if err != nil {
		return err
	}

	chosen := wondrous[rand.Intn(len(wondrous))]
	log.Printf("%s wondrous item\n%s: %s", rarity, chosen.Name, chosen.Description)
	return nil
}

var RING_WEIGHTS = []int{50, 79, 91, 97, 99, 100}
var RING_RARITIES = []string{"gold", "uncommon", "rare", "very rare", "legendary", "artifact"}
var ring = func() error {
	rarityRoll := rand.Intn(100)
	var rarity string
	for i := range RING_WEIGHTS {
		if rarityRoll < RING_WEIGHTS[i] {
			rarity = RING_RARITIES[i]
			break
		}
	}

	if rarity == "gold" {
		if err := mediumGold(); err != nil {
			return err
		}
		return nil
	}
	rings, err := fetchRings(rarity)
	if err != nil {
		return err
	}

	chosen := rings[rand.Intn(len(rings))]
	var modDescriptions []string
	for _, m := range chosen.Effects {
		modDescriptions = append(modDescriptions, m)
	}
	modString := strings.Join(modDescriptions, "\n- ")
	log.Printf("%s ring\n\n%s:\n- %s", rarity, chosen.Name, modString)
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

	log.Printf("2E magic item\nBase: %s (%s)\n- %s [%spts; %s] [MIN 2PTS]\n- %s [%spts; %s] [MIN 2PTS]", base.Name, base.Description, enchants[0].Description, enchants[0].PointValue, enchants[0].Upgrade, enchants[1].Description, enchants[1].PointValue, enchants[1].Upgrade)
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
	sets, err := fetchAmulets()
	if err != nil {
		return err
	}

	set := sets[rand.Intn(len(sets))]
	chosen := set.Amulets[rand.Intn(len(set.Amulets))]
	log.Printf("Amulet\n%s (%s): %s", chosen.Name, set.Name, chosen.Effect)
	return nil
}

var belt = func() error {
	belts, err := fetchGenerics("belt")
	if err != nil {
		return err
	}

	belt := belts[rand.Intn(len(belts))]
	log.Printf("Utility Belt\n%s: %s", belt.Name, belt.Description)
	return nil
}

var positiveReward = func() error {
	positiveRewards, err := fetchSimpleGenerics("positiveReward")
	if err != nil {
		return err
	}

	chosen := positiveRewards[rand.Intn(len(positiveRewards))]
	log.Printf("Positive encounter reward\n%s", chosen.Description)
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

var doubleValueDoubleEnchant = func() error {
	base, err := getMundane("standard")
	if err != nil {
		return err
	}

	enchants, err := getEnchants(2, base.Tags)
	if err != nil {
		return err
	}

	log.Printf("2E magic item\nBase: %s (%s)\n- %s [%spts; %s] (HARD DOUBLE)\n- %s [%spts; %s] (HARD DOUBLE)", base.Name, base.Description, enchants[0].Description, enchants[0].PointValue, enchants[0].Upgrade, enchants[1].Description, enchants[1].PointValue, enchants[1].Upgrade)
	return nil
}

var shrine = func() error {
	shrines, err := fetchGenerics("shrine")
	if err != nil {
		return err
	}

	chosen := shrines[rand.Intn(len(shrines))]
	log.Printf("Shrine of %s: %s", chosen.Name, chosen.Description)
	return nil
}

var tarot = func() error {
	cardP := promptui.Prompt{
		Label:    "Card",
		Validate: validateTarotCard,
	}
	card, err := cardP.Run()
	if err != nil {
		return err
	}

	var cardIdx int
	for i, t := range TAROT_CARDS {
		if strings.ToLower(t) == strings.ToLower(card) {
			cardIdx = i
			break
		}
	}

	c, err := fetchTarot(cardIdx)
	log.Printf("Tarot card\n%s: %s", c.Name, c.Description)
	return nil
}

var relic = func() error {
	relics, err := fetchRelics()
	if err != nil {
		return err
	}
	typeRoll := rand.Intn(3)
	options := relics.Armour
	t := "armour"
	if typeRoll < 1 {
		options = relics.Weapon
		t = "weapon"
	}

	chosen := options[rand.Intn(len(options))]
	var modDescriptions []string
	for _, m := range chosen.StartingMods {
		m.Description = processMod(m.Description)
		modDescriptions = append(modDescriptions, m.Description)
	}
	modString := strings.Join(modDescriptions, "\n- ")
	log.Printf("Relic\n%s (%s):\n- %s", chosen.Name, t, modString)
	return nil
}

var body = func() error {
	bodies, err := fetchBodies()
	if err != nil {
		return err
	}

	weight := ARMOUR_WEIGHTS[rand.Intn(len(ARMOUR_WEIGHTS))]
	var options []Body
	switch weight {
	case "unarmoured":
		options = bodies.Unarmoured
	case "light":
		options = bodies.Light
	case "medium":
		options = bodies.Medium
	case "heavy":
		options = bodies.Heavy
	default:
		return errors.New("Somehow rolled invalid body armour weight: " + weight)
	}

	chosen := options[rand.Intn(len(options))]
	var mods []string
	for _, m := range chosen.Mods {
		mods = append(mods, m)
	}
	if len(chosen.Variables) > 0 {
		chosenVar := chosen.Variables[rand.Intn(len(chosen.Variables))]
		for _, m := range chosenVar.Mods {
			mods = append(mods, m)
		}
		if len(chosenVar.Variables) > 0 {
			mods = append(mods, chosenVar.Variables[rand.Intn(len(chosenVar.Variables))])
		}
	}
	modString := strings.Join(mods, "\n- ")

	log.Printf("Body armour\n%s (%s):\n- %s", chosen.Name, weight, modString)
	return nil
}
