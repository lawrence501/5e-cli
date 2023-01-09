package main

import (
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
	"golang.org/x/exp/slices"
)

var colour = func() error {
	rollP := promptui.Prompt{
		Label:    "Loot type",
		Validate: validateColourUpgrade,
	}
	r, err := rollP.Run()
	if err != nil {
		return err
	}

	log.Printf("Colour upgrade:\n%s", COLOUR_UPGRADE_DESCRIPTIONS[r])
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
		hIdx := rand.Intn(len(tiers)-1) + 1
		hints = append(hints, chosen.Tiers[hIdx])
		tiers = append(tiers[:hIdx], tiers[hIdx+1:]...)
	}
	log.Printf("Glyph path\n%s\n%s\n\nRELIGION CHECK HINTS:\nDC 13: Path is all about %s.\n\nDC 15: %s\n\nDC 18: %s\n\nDC 20: %s", chosen.Name, t1, chosen.Description, hints[0], hints[1], hints[2])
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

var empower = func() error {
	empowerments, err := fetchGenerics("enemyEmpowerment")
	if err != nil {
		return err
	}

	chosen := empowerments[rand.Intn(len(empowerments))]
	log.Printf("The encounter will be %s:\n%s", chosen.Name, processMod(chosen.Description))
	return nil
}

var randomEncounter = func() error {
	encounter, err := generateEncounter()
	if err != nil {
		return err
	}
	log.Printf("Random encounter!\n%s", encounter)
	return nil
}

var insight = func() error {
	socialP := promptui.Prompt{
		Label:    "Social bonus of speaker",
		Validate: validateInt,
	}
	socialBonus, err := socialP.Run()
	if err != nil {
		return err
	}
	socialBonusInt, err := strconv.Atoi(socialBonus)
	if err != nil {
		return err
	}
	socialCheck := rand.Intn(20) + socialBonusInt
	log.Printf("SocialCheck: %d", socialCheck)
	insightReflect := reflect.ValueOf(INSIGHTS).MapKeys()
	var players []string
	for _, v := range insightReflect {
		players = append(players, v.Interface().(string))
	}
	sort.Strings(players)
	for _, p := range players {
		var results []string
		for i := 0; i < 2; i++ {
			insightCheck := rand.Intn(20) + INSIGHTS[p]
			// log.Printf("%s's Insight #%d: %d", p, i+1, insightCheck)
			var result string
			if insightCheck > socialCheck+2 {
				result = "succeeds"
			} else if insightCheck < socialCheck-2 {
				result = "fails"
			} else {
				result = "is unsure"
			}
			results = append(results, result)
		}
		log.Printf("%s %s (or %s with reroll)", p, results[0], results[1])
	}
	return nil
}

var gem = func() error {
	tagP := promptui.Prompt{
		Label:    "Soul gem tag",
		Validate: validateGem,
	}
	tag, err := tagP.Run()
	if err != nil {
		return err
	}

	gems, err := fetchGenericEnchants("gem")
	if err != nil {
		return err
	}

	var chosen Enchant
	for {
		chosen = gems[rand.Intn(len(gems))]
		if slices.Contains(chosen.Tags, tag) {
			log.Printf("Gem affix: %s", processMod(chosen.Description))
			return nil
		}
	}
}

var craft = func() error {
	crafts, err := fetchGenericEnchants("craft")
	if err != nil {
		return err
	}

	chosen := crafts[rand.Intn(len(crafts))]
	chosen.Description = processMod(chosen.Description)
	log.Printf("Crafted mod: %s (%spts - %s; %v)", chosen.Description, chosen.PointValue, chosen.Upgrade, chosen.Tags)
	return nil
}

var targetCraft = func() error {
	affP := promptui.Prompt{
		Label:    "Affinities (space-separated)",
		Validate: validateSpaceSeparated,
	}
	affString, err := affP.Run()
	if err != nil {
		return err
	}

	affinities := strings.Split(affString, " ")

	crafts, err := fetchGenericEnchants("craft")
	if err != nil {
		return err
	}

	var chosen Enchant
	for {
		chosen = crafts[rand.Intn(len(crafts))]
		for _, affinity := range affinities {
			if slices.Contains(chosen.Tags, affinity) {
				chosen.Description = processMod(chosen.Description)
				log.Printf("Crafted mod: %s (%spts - %s; %v)", chosen.Description, chosen.PointValue, chosen.Upgrade, chosen.Tags)
				return nil
			}
		}
	}
}

var dmgUpgrade = func() error {
	dmgP := promptui.Prompt{
		Label:    "Current average damage",
		Validate: validateFloat,
	}
	dmgString, err := dmgP.Run()
	if err != nil {
		return err
	}

	multiP := promptui.Prompt{
		Label:    "Damage multiplier",
		Validate: validateFloat,
	}
	multiString, err := multiP.Run()
	if err != nil {
		return err
	}

	multiplier, _ := strconv.ParseFloat(multiString, 64)
	currentDmg, _ := strconv.ParseFloat(dmgString, 64)
	newDmg := currentDmg * multiplier
	fmt.Printf("New damage dice: %s (if mod-based: %s + 5)", dmgToDice(newDmg), dmgToDice(newDmg-5))

	return nil
}

var chaos = func() error {
	chaos, err := fetchChaos()
	if err != nil {
		return err
	}

	chaosTrigger := chaos.Trigger[rand.Intn(len(chaos.Trigger))]
	chaosTarget := chaos.Target[rand.Intn(len(chaos.Target))]
	log.Printf("Chaotic modifier: %s, cast [https://5e.tools/spells.html#blankhash,flstsubschool:maneuver=2] on %s", processMod(chaosTrigger), chaosTarget)
	return nil
}

const FLARE_CHANCE = 10

var combat = func() error {
	roundMap := map[int][]string{
		1:  {},
		2:  {},
		3:  {},
		4:  {},
		5:  {},
		6:  {},
		7:  {},
		8:  {},
		9:  {},
		10: {},
	}

	log.Println("Combat breakdown:")
	for i := 1; i <= 10; i++ {
		flares := false
		for _, char := range PARTY_MEMBERS {
			if rand.Intn(100) < FLARE_CHANCE {
				roundMap[i] = append(roundMap[i], char)
				flares = true
			}
		}
		if flares {
			log.Printf("Round %d: the crystals of %v flare", i, roundMap[i])
		}
	}
	return nil
}

var travel = func() error {
	charSlice := make([]string, len(PARTY_MEMBERS))
	copy(charSlice, PARTY_MEMBERS)
	rand.Shuffle(len(charSlice), func(i, j int) { charSlice[i], charSlice[j] = charSlice[j], charSlice[i] })

	weather, err := generateWeather()
	if err != nil {
		return err
	}

	log.Printf("Journey travel day:\n\n")
	log.Printf("WEATHER: %s\n\n", weather)
	event := 1
	for i := 0; i < 5; i++ {
		encRoll := rand.Intn(100)
		if encRoll < 5 {
			encounter, err := generateEncounter()
			if err != nil {
				return err
			}
			log.Printf("%d. Random encounter: %s\n", event, encounter)
			event++
		}

		if i < 4 {
			log.Printf("%d. %s's activity", event, charSlice[i])
			event++
		}
	}
	return nil
}

var journeyActivity = func() error {
	activityP := promptui.Prompt{
		Label:    "Undertaken activity",
		Validate: validateActivity,
	}
	activity, err := activityP.Run()
	if err != nil {
		return err
	}
	scoreP := promptui.Prompt{
		Label:    "Score",
		Validate: validateInt,
	}
	s, err := scoreP.Run()
	if err != nil {
		return err
	}
	score, err := strconv.Atoi(s)
	if err != nil {
		return err
	}

	var result string
	if score == 100 {
		result = "crit succ"
	} else if score == -100 {
		result = "crit fail"
	} else if score >= 23 {
		result = "sub succ"
	} else if score >= 18 {
		result = "succ"
	} else if score >= 13 {
		result = "fail"
	} else {
		result = "sub fail"
	}
	fmt.Printf("%s result (%s): %s", activity, result, ACTIVITY_RESULTS[activity]["result"])
	return nil
}

var npc = func() error {
	log.Printf("NPC: %s %s", GENDERS[rand.Intn(len(GENDERS))], RACES[rand.Intn(len(RACES))])
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

var ability = func() error {
	log.Printf("Ability: %s", ABILITIES[rand.Intn(len(ABILITIES))])
	return nil
}

var condition = func() error {
	log.Printf("Condition: %s", CONDITIONS[rand.Intn(len(CONDITIONS))])
	return nil
}

var dmgPolarity = func() error {
	log.Printf("Damage polarity: %s", DAMAGE_POLARITIES[rand.Intn(len(DAMAGE_POLARITIES))])
	return nil
}

var partyMember = func() error {
	log.Printf("Party member: %s", PARTY_MEMBERS[rand.Intn(len(PARTY_MEMBERS))])
	return nil
}

var xiloan = func() error {
	log.Printf("Xiloan: %s", XILOANS[rand.Intn(len(XILOANS))])
	return nil
}

var weaponClass = func() error {
	log.Printf("Weapon class: %s", WEAPON_CLASSES[rand.Intn(len(WEAPON_CLASSES))])
	return nil
}

var physType = func() error {
	log.Printf("Physical damage type: %s", PHYS_TYPES[rand.Intn(len(PHYS_TYPES))])
	return nil
}

var nonPhysType = func() error {
	log.Printf("Non-physical damage type: %s", NON_PHYS_TYPES[rand.Intn(len(NON_PHYS_TYPES))])
	return nil
}

var class = func() error {
	log.Printf("Class: %s", CLASSES[rand.Intn(len(CLASSES))])
	return nil
}

var activity = func() error {
	log.Printf("Journey activity: %s", JOURNEY_ACTIVITIES[rand.Intn(len(JOURNEY_ACTIVITIES))])
	return nil
}
