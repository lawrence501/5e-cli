package main

import (
	"log"
	"math/rand"
	"strings"

	"github.com/manifoldco/promptui"
)

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
	tagP := promptui.Prompt{
		Label:    "Weapon tags",
		Validate: validateTagString,
	}
	t, err := tagP.Run()
	if err != nil {
		return err
	}

	tags := strings.Split(t, ",")
	tags = append(tags, "weapon")
	enchants, err := getEnchants(1, tags)
	if err != nil {
		return err
	}

	enchant := enchants[0]
	log.Printf("Weapon enchant\n%s [%s; %s]", enchant.Description, enchant.PointValue, enchant.Upgrade)
	return nil
}

var armourEnchant = func() error {
	tagP := promptui.Prompt{
		Label:    "Armour tags",
		Validate: validateTagString,
	}
	t, err := tagP.Run()
	if err != nil {
		return err
	}

	tags := strings.Split(t, ",")
	tags = append(tags, "armour")
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
