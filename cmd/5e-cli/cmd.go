package main

import (
	"fmt"
	"math"
	"math/rand"
	"slices"
	"sort"
	"strings"
)

var carnival = func(_ Request) (ViewModel, error) {
	allGames, err := fetchData("carnivalGame", []string{})
	if err != nil {
		return ViewModel{}, err
	}
	return vmText("Carnival game!", processString(randSelect(allGames))), nil
}

var colour = func(req Request) (ViewModel, error) {
	target := req.str("type")
	desc, ok := COLOUR_UPGRADE_DESCRIPTIONS[target]
	if !ok {
		return ViewModel{}, fmt.Errorf("input %q must be a loot type: one of %s",
			"type", strings.Join(sortedKeys(COLOUR_UPGRADE_DESCRIPTIONS), ", "))
	}
	return vmText("Colour upgrade", desc), nil
}

var affix = func(_ Request) (ViewModel, error) {
	allAffixes, err := fetchData("affix", []Affix{})
	if err != nil {
		return ViewModel{}, err
	}
	return ViewModel{Title: "Affix", Sections: sectionOf(affixItem(randSelect(allAffixes)))}, nil
}

var targetAffix = func(req Request) (ViewModel, error) {
	allAffixes, err := fetchData("affix", []Affix{})
	if err != nil {
		return ViewModel{}, err
	}

	affinities := strings.Fields(req.str("affinities"))
	if len(affinities) == 0 {
		return ViewModel{}, fmt.Errorf("input %q (space-separated affinities) is required", "affinities")
	}
	for _, a := range affinities {
		if !slices.Contains(AFFINITIES, a) {
			return ViewModel{}, fmt.Errorf("invalid affinity %q", a)
		}
	}
	chosenAffinity := randSelect(affinities)

	var matching []Affix
	for _, a := range allAffixes {
		if slices.Contains(a.Affinities, chosenAffinity) {
			matching = append(matching, a)
		}
	}
	if len(matching) == 0 {
		return ViewModel{}, fmt.Errorf("no affix has affinity %q", chosenAffinity)
	}
	return ViewModel{Title: "Affix", Subtitle: "targeting " + chosenAffinity,
		Sections: sectionOf(affixItem(randSelect(matching)))}, nil
}

var glyph = func(_ Request) (ViewModel, error) {
	paths, err := fetchData("glyph", []GlyphPath{})
	if err != nil {
		return ViewModel{}, err
	}
	chosen := randSelect(paths)
	return ViewModel{
		Title:    fmt.Sprintf("Glyph path about %s", chosen.Theme),
		Sections: sectionOf(Item{Body: chosen.Tiers[0]}),
	}, nil
}

var mutate = func(_ Request) (ViewModel, error) {
	mutations, err := fetchData("mutation", []Mutation{})
	if err != nil {
		return ViewModel{}, err
	}
	chosen := randSelect(mutations)
	return ViewModel{
		Title:    fmt.Sprintf("Mutation of %s", chosen.Name),
		Subtitle: fmt.Sprintf("Can be applied to %s", chosen.Target),
		Sections: sectionOf(Item{Body: processString(chosen.Description)}),
	}, nil
}

var insight = func(req Request) (ViewModel, error) {
	socialBonus, ok := req.num("socialBonus")
	if !ok {
		return ViewModel{}, fmt.Errorf("input %q (the speaker's social bonus) is required", "socialBonus")
	}
	socialCheck := rand.Intn(20) + int(socialBonus)

	players := make([]string, 0, len(INSIGHTS))
	for p := range INSIGHTS {
		players = append(players, p)
	}
	sort.Strings(players)

	items := make([]Item, 0, len(players))
	for _, p := range players {
		results := make([]string, 2)
		for i := range results {
			insightCheck := rand.Intn(20) + INSIGHTS[p]
			switch {
			case insightCheck > socialCheck+2:
				results[i] = "succeeds"
			case insightCheck < socialCheck-2:
				results[i] = "fails"
			default:
				results[i] = "is unsure"
			}
		}
		items = append(items, Item{
			Title: p,
			Body:  fmt.Sprintf("%s (or %s with reroll)", results[0], results[1]),
		})
	}
	return ViewModel{
		Title:    "Insight checks",
		Subtitle: fmt.Sprintf("Social check: %d", socialCheck),
		Sections: []Section{{Items: items}},
	}, nil
}

var dmgUpgrade = func(req Request) (ViewModel, error) {
	currentDmg, ok := req.num("currentDmg")
	if !ok {
		return ViewModel{}, fmt.Errorf("input %q (current average damage) is required", "currentDmg")
	}
	multiplier, ok := req.num("multiplier")
	if !ok {
		return ViewModel{}, fmt.Errorf("input %q (damage multiplier) is required", "multiplier")
	}
	newDmg := math.Round(currentDmg*multiplier*10) / 10
	body := fmt.Sprintf("New damage dice: %s (if mod-based: %s + 5) (average difference: %.0f)",
		dmgToDice(newDmg), dmgToDice(newDmg-5), math.Floor(newDmg-currentDmg))
	return vmText("Damage upgrade", body), nil
}

var chaos = func(_ Request) (ViewModel, error) {
	chaos, err := fetchData("chaos", Chaos{})
	if err != nil {
		return ViewModel{}, err
	}
	mod := fmt.Sprintf("%s, cast [https://spies-and-spiders.github.io/spells.html#blankhash,flstschool:m=2] on %s",
		randSelect(chaos.Trigger), randSelect(chaos.Target))
	return vmText("Chaotic modifier", processString(mod)), nil
}

var PERK_CHANCE = 20
var CARNIVAL_CHANCE = 11
var OTHERWORLDLY_CHANCE = 3

var combat = func(_ Request) (ViewModel, error) {
	var flares []Item
	for i := 1; i <= 10; i++ {
		var flared []string
		for _, char := range PARTY_MEMBERS {
			if rand.Intn(100) < FLARE_CHANCES[char] {
				flared = append(flared, char)
			}
		}
		if len(flared) > 0 {
			flares = append(flares, Item{
				Title: fmt.Sprintf("Round %d", i),
				Body:  fmt.Sprintf("the crystals of %s flare", strings.Join(flared, ", ")),
			})
		}
	}

	var rewards []Item
	if rand.Intn(100) < PERK_CHANCE {
		rewards = append(rewards, Item{Body: "2d4 perk points"})
	}
	if rand.Intn(100) < CARNIVAL_CHANCE {
		rewards = append(rewards, Item{Body: "Carnival ticket"})
	}
	if rand.Intn(100) < OTHERWORLDLY_CHANCE {
		rewards = append(rewards, Item{Body: processString("Otherworldly gift ($inactiveDm)")})
	}

	var sections []Section
	if len(flares) > 0 {
		sections = append(sections, Section{Heading: "Crystal flares", Items: flares})
	}
	if len(rewards) > 0 {
		sections = append(sections, Section{Heading: "Rewards", Items: rewards})
	}
	if len(sections) == 0 {
		sections = sectionOf(Item{Body: "No crystal flares or bonus rewards this combat."})
	}
	return ViewModel{Title: "Combat breakdown", Sections: sections}, nil
}

var LEGENDARY_CHEST_CHANCE = 0
var HOSTILE_CHANCE = 0
var POSITIVE_CHANCE = 0

var travel = func(_ Request) (ViewModel, error) {
	charSlice := make([]string, len(PARTY_MEMBERS))
	copy(charSlice, PARTY_MEMBERS)
	rand.Shuffle(len(charSlice), func(i, j int) { charSlice[i], charSlice[j] = charSlice[j], charSlice[i] })

	weather, err := generateWeather()
	if err != nil {
		return ViewModel{}, err
	}

	hostile1, hostile2 := -1, -1
	if rand.Intn(100) < HOSTILE_CHANCE {
		for hostile1 == hostile2 {
			hostile1 = rand.Intn(len(charSlice) + 1)
			hostile2 = rand.Intn(len(charSlice) + 1)
		}
	}
	positive := -1
	if rand.Intn(100) < POSITIVE_CHANCE {
		for positive == -1 || positive == hostile1 || positive == hostile2 {
			positive = rand.Intn(len(charSlice) + 1)
		}
	}

	var items []Item
	event := 1
	for i := 0; i < len(charSlice)+1; i++ {
		switch i {
		case hostile1, hostile2:
			if rand.Intn(100) < LEGENDARY_CHEST_CHANCE {
				items = append(items, Item{Body: fmt.Sprintf("%d. Legendary chest!", event)})
			} else {
				items = append(items, Item{Body: fmt.Sprintf("%d. Hostile random encounter", event)})
			}
			event++
		case positive:
			items = append(items, Item{Body: fmt.Sprintf("%d. Positive random encounter", event)})
			event++
		}
		if i < len(charSlice) {
			items = append(items, Item{Body: fmt.Sprintf("%d. %s's activity", event, charSlice[i])})
			event++
		}
	}
	return ViewModel{
		Title:    "Journey travel day",
		Subtitle: "Weather: " + weather,
		Sections: []Section{{Heading: "Events", Items: items}},
	}, nil
}

var dream = func(req Request) (ViewModel, error) {
	char := req.str("character")
	if char == "" {
		char = randSelect(PARTY_MEMBERS)
	} else if !slices.Contains(PARTY_MEMBERS, char) {
		return ViewModel{}, fmt.Errorf("invalid party member %q", char)
	}

	pool, err := fetchDreamPool(char)
	if err != nil {
		return ViewModel{}, err
	}
	if len(pool) == 0 {
		return ViewModel{}, fmt.Errorf("no dream pool configured for %q", char)
	}
	mod := randSelect(pool)
	return ViewModel{
		Title:    fmt.Sprintf("%s's dream", char),
		Sections: sectionOf(Item{Body: processString(mod.Description), Metadata: []string{mod.PointValue, mod.Upgrade}}),
	}, nil
}

var augment = func(req Request) (ViewModel, error) {
	char := req.str("character")
	if char == "" {
		char = randSelect(PARTY_MEMBERS)
	} else if !slices.Contains(PARTY_MEMBERS, char) {
		return ViewModel{}, fmt.Errorf("invalid party member %q", char)
	}

	augments, err := fetchAugments(char)
	if err != nil {
		return ViewModel{}, err
	}
	if len(augments) == 0 {
		return ViewModel{}, fmt.Errorf("no dream pool configured for %q", char)
	}
	a1 := randSelect(augments)
	a2 := a1
	a3 := a1
	for a1 == a2 || a1 == a3 || a2 == a3 {
		a2 = randSelect(augments)
		a3 = randSelect(augments)
	}

	var augmentDescriptions []string
	for _, a := range []Augment{a1, a2, a3} {
		augmentDescriptions = append(augmentDescriptions, fmt.Sprintf("**%s:** %s", a.Name, a.Description))
	}

	return ViewModel{
		Title: fmt.Sprintf("%s's Augment Options", char),
		Sections: sectionOf(
			Item{Title: a1.Name, Body: a1.Description},
			Item{Title: a2.Name, Body: a2.Description},
			Item{Title: a3.Name, Body: a3.Description},
		),
	}, nil
}

var mission = func(_ Request) (ViewModel, error) {
	missions, err := fetchData("mission", []string{})
	if err != nil {
		return ViewModel{}, err
	}
	return vmText("Mission", randSelect(missions)), nil
}

var crystal = func(_ Request) (ViewModel, error) {
	crystals, err := fetchData("crystal", []Crystal{})
	if err != nil {
		return ViewModel{}, err
	}
	passives, err := fetchData("crystalPassive", map[string]string{})
	if err != nil {
		return ViewModel{}, err
	}
	powers, err := fetchData("crystalPower", map[string]string{})
	if err != nil {
		return ViewModel{}, err
	}

	chosen := randSelect(crystals)
	passive := randSelect(chosen.Passives)
	power := randSelect(chosen.Powers)
	return ViewModel{
		Title: "Crystal: " + chosen.Name,
		Sections: sectionOf(
			Item{Title: "Passive", Body: fmt.Sprintf("%s (%s)", passive, passives[passive])},
			Item{Title: "Starting power", Body: fmt.Sprintf("%s (%s)", power, powers[power])},
		),
	}, nil
}

var crystalPower = func(req Request) (ViewModel, error) {
	crystals, err := fetchData("crystal", []Crystal{})
	if err != nil {
		return ViewModel{}, err
	}
	powers, err := fetchData("crystalPower", map[string]string{})
	if err != nil {
		return ViewModel{}, err
	}

	crystalName := req.str("crystal")
	if crystalName == "" {
		return ViewModel{}, fmt.Errorf("input %q (the crystal's name) is required", "crystal")
	}
	idx := slices.IndexFunc(crystals, func(c Crystal) bool { return c.Name == crystalName })
	if idx < 0 {
		return ViewModel{}, fmt.Errorf("unknown crystal %q", crystalName)
	}
	power := randSelect(crystals[idx].Powers)
	return ViewModel{
		Title:    fmt.Sprintf("Crystal Power (%s)", crystalName),
		Sections: sectionOf(Item{Title: power, Body: powers[power]}),
	}, nil
}

var amulet = func(_ Request) (ViewModel, error) {
	amulets, err := fetchData("amulet", []Amulet{})
	if err != nil {
		return ViewModel{}, err
	}
	a := randSelect(amulets)
	return ViewModel{Title: "Amulet", Sections: sectionOf(Item{Title: a.Name, Body: a.Base.Description})}, nil
}

var prize = func(_ Request) (ViewModel, error) {
	prizes, err := fetchData("prize", []string{})
	if err != nil {
		return ViewModel{}, err
	}
	return vmText("Prize", processString(randSelect(prizes))), nil
}

var tarot = func(req Request) (ViewModel, error) {
	card := req.str("card")
	if card == "" {
		return ViewModel{}, fmt.Errorf("input %q (the tarot card's name) is required", "card")
	}
	cardIdx := slices.IndexFunc(TAROT_CARDS, func(t string) bool { return strings.EqualFold(t, card) })
	if cardIdx < 0 {
		return ViewModel{}, fmt.Errorf("unknown tarot card %q", card)
	}
	cards, err := fetchData("tarot", []Generic{})
	if err != nil {
		return ViewModel{}, err
	}
	c := cards[cardIdx]
	return ViewModel{Title: "Tarot card", Sections: sectionOf(Item{Title: c.Name, Body: c.Description})}, nil
}

var relic = func(_ Request) (ViewModel, error) {
	relics, err := fetchData("relic", []Relic{})
	if err != nil {
		return ViewModel{}, err
	}
	chosen := randSelect(relics)
	items := make([]Item, 0, len(chosen.StartingAffixes))
	for _, m := range chosen.StartingAffixes {
		m.Description = processString(m.Description)
		var modDescriptions = []string{fmt.Sprintf("[%s; %s]", m.PointValue, m.Upgrade)}
		items = append(items, Item{Body: processString(m.Description), Metadata: modDescriptions})

	}
	return ViewModel{Title: "Relic", Subtitle: chosen.Name, Sections: []Section{{Items: items}}}, nil
}

var loot = func(req Request) (ViewModel, error) {
	// Occasionally an enchanted item is replaced by a party specialisation type.
	roll := rand.Intn(10)
	if roll < len(SPECIALISATION_TYPES["REFERENCE"]) {
		char := req.str("character")
		if char == "" {
			char = randSelect(PARTY_MEMBERS)
		} else if !slices.Contains(PARTY_MEMBERS, char) {
			return ViewModel{}, fmt.Errorf("invalid party member %q", char)
		}
		return vmText("Loot", fmt.Sprintf("Enchanted item replaced with %s!", randSelect(SPECIALISATION_TYPES[char]))), nil
	}
	return enchantedItem(req)
}

var enchantedItem = func(_ Request) (ViewModel, error) {
	allAffixes, err := fetchData("affix", []Affix{})
	if err != nil {
		return ViewModel{}, err
	}
	items := make([]Item, 0, 2)
	for len(items) < 2 {
		items = append(items, affixItem(randSelect(allAffixes)))
	}
	return ViewModel{Title: "Enchanted Item", Sections: []Section{{Items: items}}}, nil
}

var shrine = func(_ Request) (ViewModel, error) {
	shrines, err := fetchData("shrine", []Generic{})
	if err != nil {
		return ViewModel{}, err
	}
	s := randSelect(shrines)
	return ViewModel{Title: "Shrine", Sections: sectionOf(Item{Title: "Shrine of " + s.Name, Body: s.Description})}, nil
}

// lootRoll resolves a d100 loot-roll: it reads an optional `roll` input (1-100)
// or rolls one itself, then dispatches to the matching band.
var lootRoll = func(req Request) (ViewModel, error) {
	roll := rand.Intn(100) + 1
	if n, ok := req.num("roll"); ok {
		roll = int(n)
	}
	for _, band := range rollTable {
		if roll <= band.ceiling {
			return band.fn(req)
		}
	}
	return ViewModel{}, fmt.Errorf("roll %d is out of range (expected 1-100)", roll)
}

// rollTable is the d100 loot-roll, ordered by inclusive ceiling.
var rollTable = []struct {
	ceiling int
	fn      CommandFunc
}{
	{8, staticVM("Loot roll", "Reroll and upgrade result with +1 colour!")},
	{21, amulet},
	{34, crystal},
	{47, staticVM("Loot roll", "Dream Mirror")},
	{60, staticVM("Loot roll", "Glyph")},
	{73, relic},
	{86, shrine},
	{99, staticVM("Loot roll", "2x Tarot Cards")},
	{100, staticVM("Loot roll", "Player's choice and upgrade result with +1 colour!")},
}

var npc = func(_ Request) (ViewModel, error) {
	return vmText("NPC", fmt.Sprintf("%s %s", randSelect(GENDERS), randSelect(RACES))), nil
}

var skill = func(_ Request) (ViewModel, error) { return vmText("Skill", randSelect(SKILLS)), nil }
var dmgType = func(_ Request) (ViewModel, error) { return vmText("Damage type", randSelect(DAMAGE_TYPES)), nil }
var creatureType = func(_ Request) (ViewModel, error) {
	return vmText("Creature type", randSelect(CREATURE_TYPES)), nil
}
var ability = func(_ Request) (ViewModel, error) { return vmText("Ability", randSelect(ABILITIES)), nil }
var condition = func(_ Request) (ViewModel, error) { return vmText("Condition", randSelect(CONDITIONS)), nil }
var dmgPolarity = func(_ Request) (ViewModel, error) {
	return vmText("Damage polarity", randSelect(DAMAGE_POLARITIES)), nil
}
var partyMember = func(_ Request) (ViewModel, error) {
	return vmText("Party member", randSelect(PARTY_MEMBERS)), nil
}
var xiloan = func(_ Request) (ViewModel, error) { return vmText("Xiloan", randSelect(XILOANS)), nil }
var weaponClass = func(_ Request) (ViewModel, error) {
	return vmText("Weapon class", randSelect(WEAPON_CLASSES)), nil
}
var physType = func(_ Request) (ViewModel, error) {
	return vmText("Physical damage type", randSelect(PHYS_TYPES)), nil
}
var nonPhysType = func(_ Request) (ViewModel, error) {
	return vmText("Non-physical damage type", randSelect(NON_PHYS_TYPES)), nil
}
var class = func(_ Request) (ViewModel, error) { return vmText("Class", randSelect(CLASSES)), nil }
var feat = func(_ Request) (ViewModel, error) { return vmText("Feat", randSelect(FEATS)), nil }
var simpleWeapon = func(_ Request) (ViewModel, error) {
	return vmText("Simple weapon", randSelect(SIMPLE_WEAPONS)), nil
}
var martialWeapon = func(_ Request) (ViewModel, error) {
	return vmText("Martial weapon", randSelect(MARTIAL_WEAPONS)), nil
}
var language = func(_ Request) (ViewModel, error) { return vmText("Language", randSelect(LANGUAGES)), nil }
var plane = func(_ Request) (ViewModel, error) { return vmText("Plane", randSelect(PLANES)), nil }
var affinity = func(_ Request) (ViewModel, error) { return vmText("Affinity", randSelect(AFFINITIES)), nil }
var weaponTrait = func(_ Request) (ViewModel, error) {
	return vmText("Weapon trait", randSelect(WEAPON_TRAITS)), nil
}
var lootResult = func(_ Request) (ViewModel, error) {
	return vmText("Loot result", randSelect(LOOT_RESULTS)), nil
}
var journeyActivity = func(_ Request) (ViewModel, error) {
	return vmText("Journey activity", randSelect(JOURNEY_ACTIVITIES)), nil
}

// sortedKeys returns a map's string keys in sorted order (for stable messages).
func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
