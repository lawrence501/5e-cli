package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strings"

	"github.com/manifoldco/promptui"
)

var tome = func() error {
	critRoll := rand.Intn(100)
	if critRoll < 5 {
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

var WONDROUS_WEIGHTS = []int{50, 80, 92, 98, 100}
var WONDROUS_RARITIES = []string{"gold", "uncommon", "rare", "very rare", "legendary"}
var wondrous = func() error {
	rarityRoll := rand.Intn(100)
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
	modDescriptions = append(modDescriptions, chosen.Effects...)
	modString := strings.Join(modDescriptions, "\n- ")
	log.Printf("%s ring\n\n%s:\n- %s", rarity, chosen.Name, modString)
	return nil
}

var amulet = func() error {
	sets, err := fetchAmulets()
	if err != nil {
		return err
	}

	set := randSelect(sets)
	chosen := randSelect(set.Amulets)
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
	log.Printf("Positive encounter reward\n%s", processMod(chosen.Description))
	return nil
}

var shrine = func() error {
	shrines, err := fetchGenerics("shrine")
	if err != nil {
		return err
	}

	chosen := shrines[rand.Intn(len(shrines))]
	log.Printf("Shrine of %s: %s", chosen.Name, processMod(chosen.Description))
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
		if strings.EqualFold(t, card) {
			cardIdx = i
			break
		}
	}

	c, err := fetchTarot(cardIdx)
	if err != nil {
		return err
	}
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
	mods = append(mods, chosen.Mods...)
	if len(chosen.Variables) > 0 {
		chosenVar := chosen.Variables[rand.Intn(len(chosen.Variables))]
		mods = append(mods, chosenVar.Mods...)
		if len(chosenVar.Variables) > 0 {
			mods = append(mods, chosenVar.Variables[rand.Intn(len(chosenVar.Variables))])
		}
	}
	modString := strings.Join(mods, "\n- ")

	log.Printf("Body armour\n%s (%s):\n- %s", chosen.Name, weight, modString)
	return nil
}

var magicItem = func() error {
	affixRoll := rand.Intn(3) + 1

	typeRoll := rand.Intn(3)
	tags := []string{}
	if typeRoll < 1 {
		tags = append(tags, "weapon")
	} else {
		tags = append(tags, "armour")
	}

	affixes, err := getEnchants(affixRoll, tags)
	if err != nil {
		return err
	}
	var modDescriptions []string
	for _, m := range affixes {
		m.Description = processMod(m.Description)
		modDescriptions = append(modDescriptions, fmt.Sprintf("%s [%s; %s]", m.Description, m.PointValue, m.Upgrade))
	}

	minPointText := ""
	switch affixRoll {
	case 1:
		minPointText = "(MIN 3 POINTS)"
	case 2:
		minPointText = "(MIN 2 POINTS)"
	}
	log.Printf("%de magic %s %s\n- %s", affixRoll, tags[0], minPointText, strings.Join(modDescriptions, "\n- "))
	return nil
}
