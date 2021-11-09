package main

import (
	"log"
	"math/rand"
	"strconv"
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
	t1 := tiers[0]
	hints := []string{}
	for i := 0; i < 3; i++ {
		hIdx := rand.Intn(len(chosen.Tiers)-1) + 1
		hints = append(hints, chosen.Tiers[hIdx])
		tiers = append(tiers[:hIdx], tiers[hIdx+1:]...)
	}
	log.Printf("Glyph path\n%s\n%s\n\nRELIGION CHECK HINTS:\nDC 13: Path is all about %s\n\nDC 15: %s\n\nDC 18: %s\n\nDC 20: %s", chosen.Name, t1, chosen.Description, hints[0], hints[1], hints[2])
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

var loot = func() error {
	rollsFound, err := getLootSearchResults("Investigation")
	if err != nil || rollsFound == 0 {
		return err
	}

	lootablesP := promptui.Prompt{
		Label:    "Number of lootable places in room",
		Validate: validateInt,
	}
	lootables, err := lootablesP.Run()
	if err != nil {
		return err
	}
	lootableCount, err := strconv.Atoi(lootables)
	if err != nil {
		return err
	}
	lootableMap := map[int]int{}
	for i := 1; i <= lootableCount; i++ {
		lootableMap[i] = 0
	}

	for i := 0; i < rollsFound; i++ {
		roll := rand.Intn(lootableCount) + 1
		lootableMap[roll] = lootableMap[roll] + 1
	}
	for i := 1; i <= lootableCount; i++ {
		log.Printf("Lootable %d: %d roll(s)", i, lootableMap[i])
	}
	return nil
}

var harvest = func() error {
	organsFound, err := getLootSearchResults("Survival")
	if err != nil || organsFound == 0 {
		return err
	}

	log.Printf("Harvested %d creature organ(s)", organsFound)
	return nil
}

var mutation = func() error {
	mutations, err := fetchMutations()
	if err != nil {
		return err
	}

	chosen := mutations[rand.Intn(len(mutations))]
	log.Printf("Mutation\n%s: %s", chosen.Name, chosen.Description)
	return nil
}

var tpk = func() error {
	tpkNotes, err := fetchGenerics("tpk")
	if err != nil {
		return err
	}

	chosen := tpkNotes[rand.Intn(len(tpkNotes))]
	log.Printf("TPK rescue offer from %s:\n\n%s", chosen.Name, processMod(chosen.Description))
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

var condition = func() error {
	log.Printf("Condition: %s", CONDITIONS[rand.Intn(len(CONDITIONS))])
	return nil
}
