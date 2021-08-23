package main

import (
	"log"
	"math/rand"
	"strings"

	"github.com/manifoldco/promptui"
)

var singleCard = func() error {
	cards, err := getCards(1)
	if err != nil {
		return err
	}

	log.Printf("Myth card\n[%s] %s (%s)", cards[0].Rarity, cards[0].Name, cards[0].Set)
	return nil
}

var colour = func() error {
	rollP := promptui.Prompt{
		Label:    "Loot roll",
		Validate: validateColourUpgrade,
	}
	r, err := rollP.Run()
	if err != nil {
		return err
	}

	log.Printf("Colour upgrade for %s:\n%s", r, COLOUR_UPGRADE_DESCRIPTIONS[r])
	return nil
}

var weaponEnchant = func() error {
	tags := []string{"weapon"}
	enchants, err := getEnchants(1, tags)
	if err != nil {
		return err
	}

	enchant := enchants[0]
	log.Printf("Weapon enchant\n%s [%s; %s]", enchant.Description, enchant.PointValue, enchant.Upgrade)
	return nil
}

var armourEnchant = func() error {
	tags := []string{"armour"}
	enchants, err := getEnchants(1, tags)
	if err != nil {
		return err
	}

	enchant := enchants[0]
	log.Printf("Armour enchant\n%s [%s; %s]", enchant.Description, enchant.PointValue, enchant.Upgrade)
	return nil
}

var upgradeRing = func() error {
	basicP := promptui.Prompt{
		Label:    "Basic ring type",
		Validate: validateBasicRing,
	}
	b, err := basicP.Run()
	if err != nil {
		return err
	}

	rings, err := fetchThematicRings()
	if err != nil {
		return err
	}

	var ring ThematicRing
	for true {
		ring = rings[rand.Intn(len(rings))]
		if sliceContains(ring.Tags, b) {
			modString := strings.Join(ring.Mods, "\n- ")
			log.Printf("Thematic ring (in addition to basic)\n- %s", modString)
			return nil
		}
	}
	return nil
}

var glyph = func() error {
	paths, err := fetchGlyphs()
	if err != nil {
		return err
	}

	chosen := paths[rand.Intn(len(paths))]
	tiers := chosen.Tiers
	hints := []string{}
	for i := 0; i < 3; i++ {
		hIdx := rand.Intn(len(chosen.Tiers)-1) + 1
		hints = append(hints, chosen.Tiers[hIdx])
		tiers = append(tiers[:hIdx], tiers[hIdx+1:]...)
	}
	log.Printf("Glyph path\n%s\n\nRELIGION CHECK HINTS:\nDC 13: %s\n\nDC 15: %s\n\nDC 18: %s\n\nDC 20: %s", chosen.Name, chosen.Description, hints[0], hints[1], hints[2])
	return nil
}

var upgradeRelic = func() error {
	log.Println("Upgrade options:")
	rolls := []int{rand.Intn(4), rand.Intn(4)}
	for _, r := range rolls {
		if r < 2 {
			log.Println("- Upgrade existing mod")
		} else if r < 3 {
			log.Println("- New random mod")
		} else {
			log.Println("- New thematic mod")
		}
	}
	return nil
}

var skill = func() error {
	log.Printf("Skill: %s", SKILLS[rand.Intn(len(SKILLS))])
	return nil
}

var dmgType = func() error {
	log.Printf("Damage type: %s", DAMAGE_TYPES[rand.Intn(len(DAMAGE_TYPES))])
	return nil
}

var creatureType = func() error {
	log.Printf("Creature type: %s", CREATURE_TYPES[rand.Intn(len(CREATURE_TYPES))])
	return nil
}

var abilityScore = func() error {
	log.Printf("Ability score: %s", ABILITY_SCORES[rand.Intn(len(ABILITY_SCORES))])
	return nil
}
