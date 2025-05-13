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

var simulations = func() error {
	simWeights, err := fetchData("simulationWeights", SimulationWeights{})
	if err != nil {
		return err
	}
	simulationAffixes, err := fetchData("simulationAffix", SimulationAffixes{})
	if err != nil {
		return err
	}

	simType, err := chooseSimulationType(simWeights)
	if err != nil {
		return err
	}

	currencyOutput, err := genSimCurrency()
	if err != nil {
		return err
	}

	ONE_AFFIX_CHANCE := 20
	TWO_AFFIX_CHANCE := 30
	THREE_AFFIX_CHANCE := 34
	UNIQUE_CHANCE := 35

	outputBuilder := ""
	outputBuilder += fmt.Sprintf("# %s Simulations:\n", simType)
	for idx := range 3 {
		outputBuilder += fmt.Sprintf("### Option %d\n", idx+1)
		affixRoll := rand.Intn(100)
		if affixRoll < ONE_AFFIX_CHANCE {
			outputBuilder += fmt.Sprintf(
				"\t* 🟩 %s\n\t* 🟥 %s\n",
				processString(genSimAffix(simType, simulationAffixes.Positive)),
				processString(genSimAffix(simType, simulationAffixes.Negative)),
			)
		} else if affixRoll < TWO_AFFIX_CHANCE {
			positive1 := processString(genSimAffix(simType, simulationAffixes.Positive))
			negative1 := processString(genSimAffix(simType, simulationAffixes.Negative))
			positive2, negative2 := positive1, negative1
			for positive1 == positive2 {
				positive2 = processString(genSimAffix(simType, simulationAffixes.Positive))
			}
			for negative1 == negative2 {
				negative2 = processString(genSimAffix(simType, simulationAffixes.Negative))
			}
			outputBuilder += fmt.Sprintf(
				"\t* 🟩 %s\n\t* 🟩 %s\n\t* 🟥 %s\n\t* 🟥 %s\n",
				positive1,
				positive2,
				negative1,
				negative2,
			)
		} else if affixRoll < THREE_AFFIX_CHANCE {
			positive1 := processString(genSimAffix(simType, simulationAffixes.Positive))
			negative1 := processString(genSimAffix(simType, simulationAffixes.Negative))
			positive2, negative2 := positive1, negative1
			for positive1 == positive2 {
				positive2 = processString(genSimAffix(simType, simulationAffixes.Positive))
			}
			for negative1 == negative2 {
				negative2 = processString(genSimAffix(simType, simulationAffixes.Negative))
			}
			positive3, negative3 := positive1, negative1
			for positive1 == positive3 {
				positive3 = processString(genSimAffix(simType, simulationAffixes.Positive))
			}
			for negative1 == negative3 {
				negative3 = processString(genSimAffix(simType, simulationAffixes.Negative))
			}
			outputBuilder += fmt.Sprintf(
				"\t* 🟩 %s\n\t* 🟩 %s\n\t* 🟩 %s\n\t* 🟥 %s\n\t* 🟥 %s\n\t* 🟥 %s\n",
				positive1,
				positive2,
				positive3,
				negative1,
				negative2,
				negative3,
			)
		} else if affixRoll < UNIQUE_CHANCE {
			outputBuilder += "\t* ❕ Unique simulation!\n"
		} else {
			outputBuilder += "*Mundane*\n"
		}
	}
	outputBuilder += fmt.Sprintf("\n## Currencies\n%s\n", currencyOutput)
	return discordSend(outputBuilder)
}

var simulationAffix = func() error {
	simTypeP := promptui.Prompt{
		Label:    "Simulation type",
		Validate: validateSimulationType,
	}
	simType, err := simTypeP.Run()
	if err != nil {
		return err
	}

	simulationAffixes, err := fetchData("simulationAffix", SimulationAffixes{})
	if err != nil {
		return err
	}

	discordSend(fmt.Sprintf(
		"* 🟩 %s\n* 🟥 %s",
		processString(genSimAffix(simType, simulationAffixes.Positive)),
		processString(genSimAffix(simType, simulationAffixes.Negative)),
	))
	return nil
}

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
		Validate: validateSpaceSeparated,
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
	discordSend(fmt.Sprintf("### Glyph path\n%s", chosen.Tiers[0]))
	return nil
}

var upgradeRelic = func() error {
	log.Println("Upgrade options:")
	rolls := []int{rand.Intn(2), rand.Intn(2)}
	for _, r := range rolls {
		if r < 1 {
			log.Println("- Upgrade existing mod")
		} else {
			log.Println("- New thematic mod")
		}
	}
	return nil
}

var mutate = func() error {
	mutations, err := fetchData("mutation", []Mutation{})
	if err != nil {
		return err
	}

	chosen := randSelect(mutations)
	return discordSend(fmt.Sprintf("### Mutation\n**Mutation of %s** *(Can be applied to %s)*:\n%s", chosen.Name, chosen.Target, chosen.Description))
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
	fmt.Printf("New damage dice: %s (if mod-based: %s + 5) (average difference: %.0f)", dmgToDice(newDmg), dmgToDice(newDmg-5), math.Floor(newDmg-currentDmg))

	return nil
}

var chaos = func() error {
	chaos, err := fetchChaos()
	if err != nil {
		return err
	}

	chaosTrigger := chaos.Trigger[rand.Intn(len(chaos.Trigger))]
	chaosTarget := chaos.Target[rand.Intn(len(chaos.Target))]
	mod := fmt.Sprintf("%s, cast [https://spies-and-spiders.github.io/spells.html#blankhash,flstschool:m=2] on %s", chaosTrigger, chaosTarget)

	log.Printf("Chaotic modifier: %s", processString(mod))
	return nil
}

var PERK_CHANCE = 0
var CARNIVAL_CHANCE = 10
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
	if hostileRoll < 19 {
		for hostile1 == hostile2 {
			hostile1 = rand.Intn(len(charSlice) + 1)
			hostile2 = rand.Intn(len(charSlice) + 1)
		}
	}
	positiveRoll := rand.Intn(100)
	positive := -1
	if positiveRoll < 5 {
		for positive == -1 || (positive == hostile1 || positive == hostile2) {
			positive = rand.Intn(len(charSlice) + 1)
		}
	}

	event := 1
	for i := 0; i < len(charSlice)+1; i++ {
		if i == hostile1 || i == hostile2 {
			legendaryRoll := rand.Intn(100)
			if legendaryRoll < LEGENDARY_CHEST_CHANCE {
				log.Printf("%d. Legendary chest!\n", event)
			} else {
				log.Printf("%d. Hostile random encounter\n", event)
			}
			event++
		} else if i == positive {
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

var augment = func() error {
	charP := promptui.Prompt{
		Label:    "Character",
		Validate: validatePartyMember,
	}
	char, err := charP.Run()
	if err != nil {
		return err
	}

	all, err := fetchData("augment", map[string][]string{})
	if err != nil {
		return err
	}
	augments := all[char]

	option1 := randSelect(augments)
	option2 := option1
	for option1 == option2 {
		option2 = randSelect(augments)
	}
	option3 := option2
	for option3 == option1 || option3 == option2 {
		option3 = randSelect(augments)
	}

	return discordSend(fmt.Sprintf("### %s's augment options:\n- %s\n- %s\n- %s", char, option1, option2, option3))
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
		modDescriptions = append(modDescriptions, m.Description)
	}
	modString := strings.Join(modDescriptions, "\n- ")
	return discordSend(fmt.Sprintf("### Relic\n**%s**:\n- %s", chosen.Name, modString))
}

var enchantedItem = func() error {
	lvlP := promptui.Prompt{
		Label:    "Crafting level",
		Validate: validateInt,
	}
	lvlS, err := lvlP.Run()
	if err != nil {
		return err
	}

	// Affixes
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

	// Crafts
	lvl, err := strconv.Atoi(lvlS)
	if err != nil {
		return err
	}

	allCrafts, err := fetchData("craft", []Craft{})
	if err != nil {
		return err
	}

	tiers := 0
	upgradeIdx := -1
	loopCapper := 0
	chosenNames := []string{}
	crafts := []Craft{}
	for tiers < lvl {
		if len(crafts) < 3 {
			c := randSelect(allCrafts)
			if c.Tier+tiers > lvl || slices.Contains(chosenNames, c.Name) {
				continue
			}
			c.Rank = c.Tier
			tiers += c.Tier
			chosenNames = append(chosenNames, c.Name)
			crafts = append(crafts, c)
			continue
		}

		if upgradeIdx == 2 {
			upgradeIdx = 0
		} else {
			upgradeIdx++
		}
		if tiers+crafts[upgradeIdx].Tier > lvl {
			loopCapper++
			if loopCapper == 3 {
				break
			}
			continue
		}
		loopCapper = 0
		crafts[upgradeIdx].Rank += crafts[upgradeIdx].Tier
		tiers += crafts[upgradeIdx].Tier
	}

	craftDescriptions := []string{}
	for _, c := range crafts {
		craftDescriptions = append(craftDescriptions, fmt.Sprintf("**%s (Rank %d):** %s *[%s]*", c.Name, c.Rank-c.Tier, c.Description, c.Upgrade))
	}
	if len(craftDescriptions) == 0 {
		craftDescriptions = append(craftDescriptions, "*None*")
	}

	return discordSend(fmt.Sprintf("## Enchanted Item\n- %s\n\n### Crafts\n- %s", strings.Join(modDescriptions, "\n- "), strings.Join(craftDescriptions, "\n* ")))
}

var FUMBLE_MODIFIER = float64(-5)

var fumble = func() error {
	roll := math.Max(float64(rand.Intn(100))+FUMBLE_MODIFIER, 0)
	fumbles, err := fetchData("fumble", []Fumble{})
	if err != nil {
		return err
	}
	f := fumbles[int64(roll)]
	log.Printf("Fumble:\n%s - %s", f.Trigger, f.Effect)
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

var simulationType = func() error {
	log.Printf("Simulation type: %s", randSelect(SIMULATION_TYPES))
	return nil
}

var orbOfChance = func() error {
	log.Printf("Chance results: %s", randSelect(CHANCE_RESULTS))
	return nil
}

var corruptionOrb = func() error {
	log.Printf("Corruption results: %s", randSelect(CORRUPTION_RESULTS))
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
