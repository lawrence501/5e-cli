package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"reflect"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
)

var carnival = func() error {
	allGames, err := fetchData("carnivalGame", []string{})
	if err != nil {
		return err
	}
	game := randSelect(allGames)
	log.Printf("Carnival game!\n%s", processString(game))
	return nil
}

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

var affix = func() error {
	allAffixes, err := fetchData("affix", []Affix{})
	if err != nil {
		return err
	}
	a := randSelect(allAffixes)
	log.Printf("### Affix\n%s *[%s; %s; %s]*", processString(a.Description), a.PointValue, a.Upgrade, strings.Join(a.Affinities, ", "))
	return nil
}

var targetAffix = func() error {
	allAffixes, err := fetchData("affix", []Affix{})
	if err != nil {
		return err
	}

	affinityP := promptui.Prompt{
		Label:    "Affinities (space-separated)",
		Validate: validateSpaceSeparatedAffinities,
	}
	affinitiesS, err := affinityP.Run()
	if err != nil {
		return err
	}
	affinities := strings.Split(affinitiesS, " ")
	chosenAffinity := randSelect(affinities)

	var a Affix
	for {
		a = randSelect(allAffixes)
		if slices.Contains(a.Affinities, chosenAffinity) {
			break
		}
	}

	log.Printf("### Affix\n%s *[%s; %s; %s]*", processString(a.Description), a.PointValue, a.Upgrade, strings.Join(a.Affinities, ", "))
	return nil
}

var glyph = func() error {
	paths, err := fetchData("glyph", []GlyphPath{})
	if err != nil {
		return err
	}

	chosen := randSelect(paths)
	discordSend(fmt.Sprintf("### Glyph path about **%s**\n%s", chosen.Theme, chosen.Tiers[0]))
	return nil
}

var mutate = func() error {
	mutations, err := fetchData("mutation", []Mutation{})
	if err != nil {
		return err
	}

	chosen := randSelect(mutations)
	return discordSend(fmt.Sprintf("### Mutation\n**Mutation of %s** *(Can be applied to %s)*:\n%s", chosen.Name, chosen.Target, processString(chosen.Description)))
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
	newDmg = math.Round(newDmg*10) / 10
	fmt.Printf("New damage dice: %s (if mod-based: %s + 5) (average difference: %.0f)", dmgToDice(newDmg), dmgToDice(newDmg-5), math.Floor(newDmg-currentDmg))

	return nil
}

var chaos = func() error {
	chaos, err := fetchData("chaos", Chaos{})
	if err != nil {
		return err
	}

	chaosTrigger := randSelect(chaos.Trigger)
	chaosTarget := randSelect(chaos.Target)
	mod := fmt.Sprintf("%s, cast [https://spies-and-spiders.github.io/spells.html#blankhash,flstschool:m=2] on %s", chaosTrigger, chaosTarget)

	log.Printf("Chaotic modifier: %s", processString(mod))
	return nil
}

var PERK_CHANCE = 20
var CARNIVAL_CHANCE = 11
var OTHERWORLDLY_CHANCE = 3
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
			if rand.Intn(100) < FLARE_CHANCES[char] {
				roundMap[i] = append(roundMap[i], char)
				flares = true
			}
		}
		if flares {
			log.Printf("Round %d: the crystals of %v flare", i, roundMap[i])
		}
	}

	log.Print("\nRewards:")

	if rand.Intn(100) < PERK_CHANCE {
		log.Print("- 2d4 perk points")
	}
	if rand.Intn(100) < CARNIVAL_CHANCE {
		log.Print("- Carnival ticket")
	}
	if rand.Intn(100) < OTHERWORLDLY_CHANCE {
		log.Printf("- Otherworldly gift (%s)", processString("$inactiveDm"))
	}
	return nil
}

var LEGENDARY_CHEST_CHANCE = 0
var HOSTILE_CHANCE = 0
var POSITIVE_CHANCE = 0

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
	if hostileRoll < HOSTILE_CHANCE {
		for hostile1 == hostile2 {
			hostile1 = rand.Intn(len(charSlice) + 1)
			hostile2 = rand.Intn(len(charSlice) + 1)
		}
	}
	positiveRoll := rand.Intn(100)
	positive := -1
	if positiveRoll < POSITIVE_CHANCE {
		for positive == -1 || (positive == hostile1 || positive == hostile2) {
			positive = rand.Intn(len(charSlice) + 1)
		}
	}

	event := 1
	for i := 0; i < len(charSlice)+1; i++ {
		switch i {
		case hostile1, hostile2:
			legendaryRoll := rand.Intn(100)
			if legendaryRoll < LEGENDARY_CHEST_CHANCE {
				log.Printf("%d. Legendary chest!\n", event)
			} else {
				log.Printf("%d. Hostile random encounter\n", event)
			}
			event++
		case positive:
			log.Printf("%d. Positive random encounter\n", event)
			event++
		}
		if i < len(charSlice) {
			log.Printf("%d. %s's activity", event, charSlice[i])
			event++
		}
	}
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
	if err != nil {
		return err
	}
	mod := randSelect(pool)
	return discordSend(fmt.Sprintf("### %s's dream\n%s *[%s; %s]*", char, mod.Description, mod.PointValue, mod.Upgrade))
}

var mission = func() error {
	missions, err := fetchData("mission", []string{})
	if err != nil {
		return err
	}
	fmt.Printf("Mission: %s", randSelect(missions))
	return nil
}

var crystal = func() error {
	crystals, err := fetchData("crystal", []Crystal{})
	if err != nil {
		return err
	}
	crystal := randSelect(crystals)
	passive := randSelect(crystal.Passives)
	power := randSelect(crystal.Powers)

	passives, err := fetchData("crystalPassive", map[string]string{})
	if err != nil {
		return err
	}
	passiveDescription := passives[passive]

	powers, err := fetchData("crystalPower", map[string]string{})
	if err != nil {
		return err
	}
	powerDescription := powers[power]

	return discordSend(fmt.Sprintf("### Crystal: %s\n- **Passive:** %s (%s)\n- **Starting power:** %s (%s)", crystal.Name, passive, passiveDescription, power, powerDescription))
}

var crystalPower = func() error {
	crystals, err := fetchData("crystal", []Crystal{})
	if err != nil {
		return err
	}
	powers, err := fetchData("crystalPower", map[string]string{})
	if err != nil {
		return err
	}

	crystalP := promptui.Prompt{
		Label:    "Crystal",
		Validate: validateSpaceSeparated,
	}
	crystalName, err := crystalP.Run()
	if err != nil {
		return err
	}

	var foundCrystal Crystal
	for _, c := range crystals {
		if c.Name == crystalName {
			foundCrystal = c
			break
		}
	}
	power := randSelect(foundCrystal.Powers)
	return discordSend(fmt.Sprintf("### Crystal Power (%s)\n**%s:** %s", crystalName, power, powers[power]))
}

var amulet = func() error {
	amulets, err := fetchData("amulet", []Amulet{})
	if err != nil {
		return err
	}

	a := randSelect(amulets)
	return discordSend(fmt.Sprintf("### Amulet\n**%s:** %s", a.Name, a.Base.Description))
}

var prize = func() error {
	prizes, err := fetchData("prize", []string{})
	if err != nil {
		return err
	}

	prize := randSelect(prizes)
	log.Printf("Prize\n%s", processString(prize))
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

	cards, err := fetchData("tarot", []Generic{})
	if err != nil {
		return err
	}
	c := cards[cardIdx]
	discordSend(fmt.Sprintf("### Tarot card\n**%s:** %s", c.Name, c.Description))
	return nil
}

var relic = func() error {
	relics, err := fetchData("relic", []Relic{})
	if err != nil {
		return err
	}

	chosen := randSelect(relics)
	var modDescriptions []string
	for _, m := range chosen.StartingAffixes {
		m.Description = processString(m.Description)
		modDescriptions = append(modDescriptions, fmt.Sprintf("%s *[%s; %s]*", m.Description, m.PointValue, m.Upgrade))
	}
	modString := strings.Join(modDescriptions, "\n- ")
	return discordSend(fmt.Sprintf("### Relic\n**%s**:\n- %s", chosen.Name, modString))
}

var loot = func() error {
	roll := rand.Intn(10)
	// Replace enchanted item with loot result
	if roll < len(SPECIALISATION_TYPES["REFERENCE"]) {
		charP := promptui.Prompt{
			Label:    "Looting character",
			Validate: validatePartyMember,
		}
		char, err := charP.Run()
		if err != nil {
			return err
		}

		log.Printf("Enchanted item replaced with %s!", randSelect(SPECIALISATION_TYPES[char]))
		return nil
	}

	return enchantedItem()
}

var enchantedItem = func() error {
	allAffixes, err := fetchData("affix", []Affix{})
	if err != nil {
		return err
	}

	affixes := []Affix{}
	for len(affixes) < 2 {
		a := randSelect(allAffixes)
		a.Description = processString(a.Description)
		affixes = append(affixes, a)
	}
	var modDescriptions []string
	for _, m := range affixes {
		modDescriptions = append(modDescriptions, fmt.Sprintf("%s *[%s; %s; %s]*", m.Description, m.PointValue, m.Upgrade, strings.Join(m.Affinities, ", ")))
	}

	return discordSend(fmt.Sprintf("### Enchanted Item\n- %s", strings.Join(modDescriptions, "\n- ")))
}

var shrine = func() error {
	shrines, err := fetchData("shrine", []Generic{})
	if err != nil {
		return err
	}
	s := randSelect(shrines)
	return discordSend(fmt.Sprintf("### Shrine\n**Shrine of %s:** %s", s.Name, s.Description))
}

var augment = func() error {
	augments, err := fetchData("augment", Augments{})
	if err != nil {
		return err
	}

	charP := promptui.Prompt{
		Label:    "Character to augment",
		Validate: validatePartyMember,
	}
	char, err := charP.Run()
	if err != nil {
		return err
	}

	var augmentList []Augment
	switch char {
	case "Quincy":
		augmentList = augments.Quincy
	case "Viktor":
		augmentList = augments.Viktor
	case "Arthur":
		augmentList = augments.Arthur
	case "Nathaniel":
		augmentList = augments.Nathaniel
	}
	a1 := randSelect(augmentList)
	a2 := a1
	a3 := a1
	for a1 == a2 || a1 == a3 || a2 == a3 {
		a2 = randSelect(augmentList)
		a3 = randSelect(augmentList)
	}
	var augmentDescriptions []string
	for _, a := range []Augment{a1, a2, a3} {
		augmentDescriptions = append(augmentDescriptions, fmt.Sprintf("**%s:** %s", a.Name, a.Description))
	}
	return discordSend(fmt.Sprintf("### %s's Augment Options\n- %s", char, strings.Join(augmentDescriptions, "\n- ")))
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

var affinity = func() error {
	log.Printf("Affinity: %s", randSelect(AFFINITIES))
	return nil
}

var weaponTrait = func() error {
	log.Printf("Weapon trait: %s", randSelect(WEAPON_TRAITS))
	return nil
}

var lootResult = func() error {
	log.Printf("Loot result: %s", randSelect(LOOT_RESULTS))
	return nil
}

var journeyActivity = func() error {
	log.Printf("Journey activity: %s", randSelect(JOURNEY_ACTIVITIES))
	return nil
}
