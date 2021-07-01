package main

import (
	"log"
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
