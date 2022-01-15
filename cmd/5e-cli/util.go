package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
)

const DATA_DIR = "data"

func sliceContains(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}

var CARD_WEIGHTS = []int{57, 87, 97, 100}
var CARD_RARITIES = []string{"common", "uncommon", "rare", "foil"}

func getCards(num int) ([]Card, error) {
	allCards, err := fetchCards()
	if err != nil {
		return []Card{}, err
	}

	var cards []Card
	for len(cards) < num {
		rarityRoll := rand.Intn(100)
		var rarity string
		for i := range CARD_WEIGHTS {
			if rarityRoll < CARD_WEIGHTS[i] {
				rarity = CARD_RARITIES[i]
				break
			}
		}

		var rarityCards []Card
		switch rarity {
		case "common":
			rarityCards = allCards.Common
		case "uncommon":
			rarityCards = allCards.Uncommon
		case "rare":
			rarityCards = allCards.Rare
		case "foil":
			foilRarity := CARD_RARITIES[rand.Intn(3)]
			switch foilRarity {
			case "common":
				rarityCards = allCards.Common
			case "uncommon":
				rarityCards = allCards.Uncommon
			case "rare":
				rarityCards = allCards.Rare
			default:
				return []Card{}, fmt.Errorf("Invalid foil myth card rarity: %s", foilRarity)
			}
			rarity = "foil - " + foilRarity
		default:
			return []Card{}, fmt.Errorf("Invalid myth card rarity: %s", rarity)
		}

		newCard := rarityCards[rand.Intn(len(rarityCards))]
		newCard.Rarity = rarity
		cards = append(cards, newCard)
	}

	rarityPoints := map[string]int{
		"common":          1,
		"uncommon":        2,
		"rare":            3,
		"foil - common":   4,
		"foil - uncommon": 5,
		"foil - rare":     6,
	}
	sort.Slice(cards, func(i, j int) bool {
		iRarity := rarityPoints[cards[i].Rarity]
		jRarity := rarityPoints[cards[j].Rarity]
		return iRarity < jRarity
	})
	return cards, nil
}

func fetchCards() (Cards, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return Cards{}, err
	}

	f, err := ioutil.ReadFile(filepath.Join(cwd, DATA_DIR, "card.json"))
	if err != nil {
		return Cards{}, err
	}

	cards := Cards{}
	err = json.Unmarshal([]byte(f), &cards)
	if err != nil {
		return Cards{}, err
	}
	return cards, nil
}

func getEnchants(num int, tags []string) ([]Enchant, error) {
	allEnchants, err := fetchEnchants()
	if err != nil {
		return []Enchant{}, err
	}

	var enchants []Enchant
	for len(enchants) < num {
		var e Enchant
		for true {
			e = allEnchants[rand.Intn(len(allEnchants))]
			valid := true
			for _, t := range e.Tags {
				if !sliceContains(tags, t) {
					valid = false
					break
				}
			}
			if valid {
				break
			}
		}
		e.Description = processMod(e.Description)
		enchants = append(enchants, e)
	}

	return enchants, nil
}

func getMundane(t string) (Mundane, error) {
	mundanes, err := fetchMundanes(t)
	if err != nil {
		return Mundane{}, err
	}

	return mundanes[rand.Intn(len(mundanes))], nil
}

func getDamageType() string {
	return DAMAGE_TYPES[rand.Intn(len(DAMAGE_TYPES))]
}

func fetchTraps(t string) ([]Generic, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return []Generic{}, err
	}

	f, err := ioutil.ReadFile(filepath.Join(cwd, DATA_DIR, "trap.json"))
	if err != nil {
		return []Generic{}, err
	}

	traps := Traps{}
	err = json.Unmarshal([]byte(f), &traps)
	if err != nil {
		return []Generic{}, err
	}

	var ret []Generic
	switch t {
	case "standard":
		ret = traps.Standard
	case "crit":
		ret = traps.Crit
	default:
		return []Generic{}, fmt.Errorf("Invalid trap type: %s", t)
	}
	return ret, nil
}

func fetchBooks(t string) ([]Generic, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return []Generic{}, err
	}

	f, err := ioutil.ReadFile(filepath.Join(cwd, DATA_DIR, "book.json"))
	if err != nil {
		return []Generic{}, err
	}

	books := Books{}
	err = json.Unmarshal([]byte(f), &books)
	if err != nil {
		return []Generic{}, err
	}

	var bookList []Generic
	switch t {
	case "Tome":
		bookList = books.Tome
	case "Manual":
		bookList = books.Manual
	default:
		return []Generic{}, fmt.Errorf("Invalid book type: %s", t)
	}
	return bookList, nil
}

func fetchBlessings() ([]string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return []string{}, err
	}

	f, err := ioutil.ReadFile(filepath.Join(cwd, DATA_DIR, "blessing.json"))
	if err != nil {
		return []string{}, err
	}

	blessings := []string{}
	err = json.Unmarshal([]byte(f), &blessings)
	if err != nil {
		return []string{}, err
	}
	return blessings, nil
}

func fetchMundanes(t string) ([]Mundane, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return []Mundane{}, err
	}

	f, err := ioutil.ReadFile(filepath.Join(cwd, DATA_DIR, "mundane.json"))
	if err != nil {
		return []Mundane{}, err
	}

	mundanes := Mundanes{}
	err = json.Unmarshal([]byte(f), &mundanes)
	if err != nil {
		return []Mundane{}, err
	}

	var ret []Mundane
	typeRoll := rand.Intn(3)
	if typeRoll < 1 {
		ret = mundanes.Weapon
	} else {
		ret = mundanes.Armour
	}
	return ret, nil
}

func generateEncounter(location string) (string, error) {
	chestChance := 0
	positiveEncounterChance := 20 + chestChance
	positiveRoll := rand.Intn(100)
	if positiveRoll < chestChance {
		return "Legendary Chest", nil
	}
	if positiveRoll < positiveEncounterChance {
		location = "positive"
	}

	allEncounters, err := fetchEncounters()
	if err != nil {
		return "", err
	}
	var encounterList []string
	switch location {
	case "plains":
		encounterList = allEncounters.Plains
	case "forest":
		encounterList = allEncounters.Forest
	case "mountain":
		encounterList = allEncounters.Mountain
	case "aquatic":
		encounterList = allEncounters.Aquatic
	case "positive":
		encounterList = allEncounters.Positive
	}
	return processMod(encounterList[rand.Intn(len(encounterList))]), nil
}

func fetchEncounters() (Encounters, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return Encounters{}, err
	}

	f, err := ioutil.ReadFile(filepath.Join(cwd, DATA_DIR, "encounter.json"))
	if err != nil {
		return Encounters{}, err
	}

	encounters := Encounters{}
	err = json.Unmarshal([]byte(f), &encounters)
	if err != nil {
		return Encounters{}, err
	}
	return encounters, nil
}

func fetchWondrous(rarity string) ([]Generic, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return []Generic{}, err
	}

	f, err := ioutil.ReadFile(filepath.Join(cwd, DATA_DIR, "wondrous.json"))
	if err != nil {
		return []Generic{}, err
	}

	wondrous := Wondrous{}
	err = json.Unmarshal([]byte(f), &wondrous)
	if err != nil {
		return []Generic{}, err
	}

	var ret []Generic
	switch rarity {
	case "common":
		ret = wondrous.Common
	case "uncommon":
		ret = wondrous.Uncommon
	case "rare":
		ret = wondrous.Rare
	case "very rare":
		ret = wondrous.VeryRare
	case "legendary":
		ret = wondrous.Legendary
	default:
		return []Generic{}, fmt.Errorf("Invalid wondrous item rarity: %s", rarity)
	}
	return ret, nil
}

func fetchEnchants() ([]Enchant, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return []Enchant{}, err
	}

	f, err := ioutil.ReadFile(filepath.Join(cwd, DATA_DIR, "enchant.json"))
	if err != nil {
		return []Enchant{}, err
	}

	enchants := []Enchant{}
	err = json.Unmarshal([]byte(f), &enchants)
	if err != nil {
		return []Enchant{}, err
	}
	return enchants, nil
}

func fetchAmulets() ([]Amulet, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return []Amulet{}, err
	}

	f, err := ioutil.ReadFile(filepath.Join(cwd, DATA_DIR, "amulet.json"))
	if err != nil {
		return []Amulet{}, err
	}

	amulets := []Amulet{}
	err = json.Unmarshal([]byte(f), &amulets)
	if err != nil {
		return []Amulet{}, err
	}
	return amulets, nil
}

func fetchSimpleGenerics(fileName string) ([]SimpleGeneric, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return []SimpleGeneric{}, err
	}

	f, err := ioutil.ReadFile(filepath.Join(cwd, DATA_DIR, fileName+".json"))
	if err != nil {
		return []SimpleGeneric{}, err
	}

	simple := []SimpleGeneric{}
	err = json.Unmarshal([]byte(f), &simple)
	if err != nil {
		return []SimpleGeneric{}, err
	}
	return simple, nil
}

func fetchThematicRings() ([]ThematicRing, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return []ThematicRing{}, err
	}

	f, err := ioutil.ReadFile(filepath.Join(cwd, DATA_DIR, "thematicRing.json"))
	if err != nil {
		return []ThematicRing{}, err
	}

	rings := []ThematicRing{}
	err = json.Unmarshal([]byte(f), &rings)
	if err != nil {
		return []ThematicRing{}, err
	}
	return rings, nil
}

func fetchGlyphs() ([]GlyphPath, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return []GlyphPath{}, err
	}

	f, err := ioutil.ReadFile(filepath.Join(cwd, DATA_DIR, "glyph.json"))
	if err != nil {
		return []GlyphPath{}, err
	}

	paths := []GlyphPath{}
	err = json.Unmarshal([]byte(f), &paths)
	if err != nil {
		return []GlyphPath{}, err
	}
	return paths, nil
}

func fetchRelics() ([]Relic, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return []Relic{}, err
	}

	f, err := ioutil.ReadFile(filepath.Join(cwd, DATA_DIR, "relic.json"))
	if err != nil {
		return []Relic{}, err
	}

	relics := []Relic{}
	err = json.Unmarshal([]byte(f), &relics)
	if err != nil {
		return []Relic{}, err
	}
	return relics, nil
}

func fetchGenerics(fileName string) ([]Generic, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return []Generic{}, err
	}

	f, err := ioutil.ReadFile(filepath.Join(cwd, DATA_DIR, fileName+".json"))
	if err != nil {
		return []Generic{}, err
	}

	generics := []Generic{}
	err = json.Unmarshal([]byte(f), &generics)
	if err != nil {
		return []Generic{}, err
	}
	return generics, nil
}

func fetchMutations() ([]Generic, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return []Generic{}, err
	}

	f, err := ioutil.ReadFile(filepath.Join(cwd, DATA_DIR, "mutation.json"))
	if err != nil {
		return []Generic{}, err
	}

	mutations := Mutation{}
	err = json.Unmarshal([]byte(f), &mutations)
	if err != nil {
		return []Generic{}, err
	}

	typeP := promptui.Prompt{
		Label:    "Mutation type",
		Validate: validateMutationType,
	}
	mutationType, err := typeP.Run()
	if err != nil {
		return []Generic{}, err
	}

	var mutationSlice []Generic
	switch mutationType {
	case "powerful":
		mutationSlice = mutations.Powerful
	case "beneficial":
		mutationSlice = mutations.Beneficial
	case "distinctive":
		mutationSlice = mutations.Distinctive
	case "harmful":
		mutationSlice = mutations.Harmful
	}
	return mutationSlice, nil
}

func getLootSearchResults(skill string) (int, error) {
	rollP := promptui.Prompt{
		Label:    skill + " results (space separated)",
		Validate: validateSpaceSeparated,
	}
	rolls, err := rollP.Run()
	if err != nil {
		return 0, err
	}

	rollsSlice := strings.Split(rolls, " ")
	playerCount := len(rollsSlice)

	totalResult := 0
	for _, i := range rollsSlice {
		intRoll, err := strconv.Atoi(i)
		if err != nil {
			return 0, err
		}
		totalResult += intRoll
	}

	if totalResult < 13*playerCount {
		log.Println("Nothing extra found.")
		return 0, nil
	}

	return int(math.Floor(float64(totalResult-(13*playerCount))/float64(playerCount)) + 1), nil
}
