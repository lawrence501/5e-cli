package main

import (
	"fmt"
	"log"
	"math"
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

var weaponAffix = func() error {
	tags := []string{"weapon"}
	allAffixes, err := fetchAffixes("affix")
	if err != nil {
		return err
	}

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

	log.Printf("Weapon enchant\n%s [%s; %s]", processMod(a.Description), a.PointValue, a.Upgrade)
	return nil
}

var armourAffix = func() error {
	tags := []string{"armour"}
	allAffixes, err := fetchAffixes("affix")
	if err != nil {
		return err
	}

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

	log.Printf("Armour enchant\n%s [%s; %s]", processMod(a.Description), a.PointValue, a.Upgrade)
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

var mutate = func() error {
	mutations, err := fetchGenerics("mutation")
	if err != nil {
		return err
	}

	chosen := randSelect(mutations)
	log.Printf("Mutation\n%s: %s", chosen.Name, chosen.Description)
	return nil
}

var challenge = func() error {
	challenges, err := fetchSimpleGenerics("challenge")
	if err != nil {
		return err
	}

	chosen := challenges[rand.Intn(len(challenges))]
	log.Printf("Challenge: %s", chosen)
	return nil
}

var randomEncounter = func() error {
	tagP := promptui.Prompt{
		Label:    "Encounter tag",
		Validate: validateEncounter,
	}
	tag, err := tagP.Run()
	if err != nil {
		return err
	}
	encounter, err := hostileEncounter(tag)
	if err != nil {
		return err
	}
	log.Printf("Random encounter!\n%s", encounter)
	return nil
}

var posiEnc = func() error {
	encounter, err := positiveEncounter()
	if err != nil {
		return err
	}
	log.Printf("Positive encounter: %s", encounter)
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

var craft = func() error {
	allCrafts, err := fetchAffixes("craft")
	if err != nil {
		return err
	}

	chosen := randSelect(allCrafts)
	log.Printf("Crafted affix: %s (%spts - %s; %v)", processMod(chosen.Description), chosen.PointValue, chosen.Upgrade, chosen.Tags)
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
	chosenAffinity := randSelect(affinities)

	allCrafts, err := fetchAffixes("craft")
	if err != nil {
		return err
	}

	var chosen Affix
	for {
		chosen = randSelect(allCrafts)
		if slices.Contains(chosen.Tags, chosenAffinity) {
			chosen.Description = processMod(chosen.Description)
			log.Printf("Crafted affix: %s (%spts - %s; %v)", chosen.Description, chosen.PointValue, chosen.Upgrade, chosen.Tags)
			return nil
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
	fmt.Printf("New damage dice: %s (if mod-based: %s + 5) (average difference: %.0f)", dmgToDice(newDmg), dmgToDice(newDmg-5), math.Floor(newDmg-currentDmg))

	return nil
}

var chaos = func() error {
	chaosTypeRoll := rand.Intn(4)
	var mod string
	switch chaosTypeRoll {
	case 0:
		chaos, err := fetchChaos()
		if err != nil {
			return err
		}

		chaosTrigger := chaos.Trigger[rand.Intn(len(chaos.Trigger))]
		chaosTarget := chaos.Target[rand.Intn(len(chaos.Target))]
		mod = fmt.Sprintf("%s, cast [https://5e.tools/spells.html#blankhash,flstsubschool:maneuver=2] on %s", chaosTrigger, chaosTarget)
	case 1:
		mod = "AI-generated [https://www.bing.com/search?q=Bing+AI&showconv=1&FORM=hpcodx] from prompt: 'Design a homebrew $class archetype for Dungeons & Dragons 5th Edition inspired by [something random from random inspiration].'"
	case 2:
		mod = "AI-generated [https://www.bing.com/search?q=Bing+AI&showconv=1&FORM=hpcodx] from prompt: 'Design a homebrew $class archetype for Dungeons & Dragons 5th Edition themed around [https://perchance.org/object]s.'"
	case 3:
		mod = "AI-generated [https://www.bing.com/search?q=Bing+AI&showconv=1&FORM=hpcodx] from prompt: 'Design a homebrew $class archetype for Dungeons & Dragons 5th Edition themed around [https://www.randomlists.com/random-animals?show_images=false&dup=false&qty=1]s.'"
	}

	log.Printf("Chaotic modifier: %s", processMod(mod))
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

	hostileRoll := rand.Intn(100)
	hostile1, hostile2 := -1, -1
	if hostileRoll < 15 {
		for hostile1 == hostile2 {
			hostile1 = rand.Intn(5)
			hostile2 = rand.Intn(5)
		}
	}
	positiveRoll := rand.Intn(100)
	positive := -1
	if positiveRoll < 10 {
		for positive == -1 || (positive == hostile1 || positive == hostile2) {
			positive = rand.Intn(5)
		}
	}

	event := 1
	for i := 0; i < 5; i++ {
		if i == hostile1 || i == hostile2 {
			ambush := ""
			tag := ""
			if event >= 6 {
				ambush = " (NIGHT AMBUSH)"
				tag = "night"
			}
			encounter, err := hostileEncounter(tag)
			if err != nil {
				return err
			}

			log.Printf("%d. Random encounter%s: %s\n", event, ambush, encounter)
			event++
		} else if i == positive {
			encounter, err := positiveEncounter()
			if err != nil {
				return err
			}
			log.Printf("%d. Positive encounter: %s\n", event, encounter)
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
	fmt.Printf("%s result (%s): %s", activity, result, ACTIVITY_RESULTS[activity][result])
	return nil
}

var dream = func() error {
	charP := promptui.Prompt{
		Label:    "Dreaming character",
		Validate: validatePartyMember,
	}
	char, err := charP.Run()
	if err != nil {
		return err
	}

	pool, err := fetchDreamPool(char)
	mod := randSelect(pool)
	fmt.Printf("%s's dream: %s [%s; %s]", char, mod.Description, mod.PointValue, mod.Upgrade)
	return nil
}

var blessing = func() error {
	charP := promptui.Prompt{
		Label:    "Blessed character",
		Validate: validatePartyMember,
	}
	char, err := charP.Run()
	if err != nil {
		return err
	}

	blessings, err := fetchBlessings(char)
	option1 := randSelect(blessings)
	option2 := option1
	for option1 == option2 {
		option2 = randSelect(blessings)
	}
	option3 := option2
	for option3 == option1 || option3 == option2 {
		option3 = randSelect(blessings)
	}

	fmt.Printf("%s's blessing options:\n- %s\n\n- %s\n\n- %s", char, option1, option2, option3)
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

var tag = func() error {
	log.Printf("Gem tag: %s", GEM_TAGS[rand.Intn(len(GEM_TAGS))])
	return nil
}

var feat = func() error {
	log.Printf("Feat: %s", FEATS[rand.Intn(len(FEATS))])
	return nil
}

var simpleWeapon = func() error {
	log.Printf("Simple Weapon: %s", SIMPLE_WEAPONS[rand.Intn(len(SIMPLE_WEAPONS))])
	return nil
}

var martialWeapon = func() error {
	log.Printf("Martial Weapon: %s", MARTIAL_WEAPONS[rand.Intn(len(MARTIAL_WEAPONS))])
	return nil
}

var language = func() error {
	log.Printf("Language: %s", LANGUAGES[rand.Intn(len(LANGUAGES))])
	return nil
}

var plane = func() error {
	log.Printf("Plane: %s", PLANES[rand.Intn(len(PLANES))])
	return nil
}

var follower = func() error {
	log.Printf("Follower: %s", randSelect(FOLLOWERS))
	return nil
}
