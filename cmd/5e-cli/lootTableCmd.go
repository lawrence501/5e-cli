package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"

	"github.com/manifoldco/promptui"
	"golang.org/x/exp/slices"
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
	chosen.Description = processString(chosen.Description)

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

var ring = func() error {
	bases, err := fetchRingBases()
	if err != nil {
		return err
	}

	chosenStone := randSelect(RING_STONES)
	log.Printf("Ring\n%s ring: %s", chosenStone, processString(bases[chosenStone].Base))
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
	log.Printf("Positive encounter reward\n%s", processString(chosen.Description))
	return nil
}

var shrine = func() error {
	shrines, err := fetchGenerics("shrine")
	if err != nil {
		return err
	}

	chosen := shrines[rand.Intn(len(shrines))]
	log.Printf("Shrine of %s: %s", chosen.Name, processString(chosen.Description))
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

	// c, err := fetchTarot(cardIdx)
	cards, err := fetchGenerics("tarot")
	if err != nil {
		return err
	}
	c := cards[cardIdx]
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
	for _, m := range chosen.StartingAffixes {
		m.Description = processString(m.Description)
		modDescriptions = append(modDescriptions, m.Description)
	}
	modString := strings.Join(modDescriptions, "\n- ")
	log.Printf("Relic\n%s (%s):\n- %s", chosen.Name, t, modString)
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

	allAffixes, err := fetchAffixes("affix")
	if err != nil {
		return err
	}

	var affixes []Affix
	for len(affixes) < affixRoll {
		var a Affix
		for {
			a = randSelect(allAffixes)
			valid := true
			for _, t := range a.Tags {
				if !slices.Contains(tags, t) {
					valid = false
					break
				}
			}
			if valid {
				break
			}
		}
		a.Description = processString(a.Description)
		affixes = append(affixes, a)
	}
	var modDescriptions []string
	for _, m := range affixes {
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
