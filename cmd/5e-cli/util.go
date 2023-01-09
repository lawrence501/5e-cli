package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
	"golang.org/x/exp/slices"
)

const DATA_DIR = "data"

func isInt(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func fetchTarot(cardIdx int) (Generic, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return Generic{}, err
	}

	f, err := os.ReadFile(filepath.Join(cwd, DATA_DIR, "tarot.json"))
	if err != nil {
		return Generic{}, err
	}

	tarots := []Generic{}
	err = json.Unmarshal([]byte(f), &tarots)
	if err != nil {
		return Generic{}, err
	}

	return tarots[cardIdx], nil
}

func getEnchants(num int, tags []string) ([]Enchant, error) {
	allEnchants, err := fetchEnchants()
	if err != nil {
		return []Enchant{}, err
	}

	var enchants []Enchant
	for len(enchants) < num {
		var e Enchant
		for {
			e = allEnchants[rand.Intn(len(allEnchants))]
			valid := true
			for _, t := range e.Tags {
				if !slices.Contains(tags, t) {
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

func fetchGenericEnchants(fileName string) ([]Enchant, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return []Enchant{}, err
	}

	f, err := os.ReadFile(filepath.Join(cwd, DATA_DIR, fileName+".json"))
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

func fetchTomes() ([]Tome, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return []Tome{}, err
	}

	f, err := os.ReadFile(filepath.Join(cwd, DATA_DIR, "tome.json"))
	if err != nil {
		return []Tome{}, err
	}

	tomes := []Tome{}
	err = json.Unmarshal([]byte(f), &tomes)
	if err != nil {
		return []Tome{}, err
	}
	return tomes, nil
}

func generateEncounter() (string, error) {
	chestChance := 0
	positiveEncounterChance := 20 + chestChance
	positiveRoll := rand.Intn(100)
	if positiveRoll < chestChance {
		return "Legendary Chest", nil
	}

	allEncounters, err := fetchEncounters()
	if err != nil {
		return "", err
	}
	encounterList := allEncounters.Hostile
	if positiveRoll < positiveEncounterChance {
		encounterList = allEncounters.Positive
	}
	return processMod(encounterList[rand.Intn(len(encounterList))]), nil
}

func fetchEncounters() (Encounters, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return Encounters{}, err
	}

	f, err := os.ReadFile(filepath.Join(cwd, DATA_DIR, "encounter.json"))
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

func fetchWeather() (Weathers, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return Weathers{}, err
	}

	f, err := os.ReadFile(filepath.Join(cwd, DATA_DIR, "weather.json"))
	if err != nil {
		return Weathers{}, err
	}

	weathers := Weathers{}
	err = json.Unmarshal([]byte(f), &weathers)
	if err != nil {
		return Weathers{}, err
	}
	return weathers, nil
}

func generateWeather() (string, error) {
	weathers, err := fetchWeather()
	if err != nil {
		return "", err
	}

	weatherRoll := rand.Intn(100)
	if weatherRoll < 5 {
		chosen := weathers.Exotic[rand.Intn(len(weathers.Exotic))]
		return fmt.Sprintf("%s (+1 minimum hostile random encounter. %s)", chosen.Name, chosen.Description), nil
	}
	chosen := weathers.Common[rand.Intn(len(weathers.Common))]
	return chosen, nil
}

func fetchWondrous(rarity string) ([]Generic, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return []Generic{}, err
	}

	f, err := os.ReadFile(filepath.Join(cwd, DATA_DIR, "wondrous.json"))
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
		return []Generic{}, fmt.Errorf("invalid wondrous item rarity: %s", rarity)
	}
	return ret, nil
}

func fetchRings(rarity string) ([]Ring, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return []Ring{}, err
	}

	f, err := os.ReadFile(filepath.Join(cwd, DATA_DIR, "ring.json"))
	if err != nil {
		return []Ring{}, err
	}

	rings := Rings{}
	err = json.Unmarshal([]byte(f), &rings)
	if err != nil {
		return []Ring{}, err
	}

	var ret []Ring
	switch rarity {
	case "uncommon":
		ret = rings.Uncommon
	case "rare":
		ret = rings.Rare
	case "very rare":
		ret = rings.VeryRare
	case "legendary":
		ret = rings.Legendary
	case "artifact":
		ret = rings.Artifact
	default:
		return []Ring{}, fmt.Errorf("invalid ring rarity: %s", rarity)
	}
	return ret, nil
}

func fetchEnchants() ([]Enchant, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return []Enchant{}, err
	}

	f, err := os.ReadFile(filepath.Join(cwd, DATA_DIR, "enchant.json"))
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

func fetchAmulets() ([]AmuletSet, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return []AmuletSet{}, err
	}

	f, err := os.ReadFile(filepath.Join(cwd, DATA_DIR, "amulet.json"))
	if err != nil {
		return []AmuletSet{}, err
	}

	sets := []AmuletSet{}
	err = json.Unmarshal([]byte(f), &sets)
	if err != nil {
		return []AmuletSet{}, err
	}
	return sets, nil
}

func fetchSimpleGenerics(fileName string) ([]SimpleGeneric, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return []SimpleGeneric{}, err
	}

	f, err := os.ReadFile(filepath.Join(cwd, DATA_DIR, fileName+".json"))
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

func fetchGlyphs() ([]GlyphPath, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return []GlyphPath{}, err
	}

	f, err := os.ReadFile(filepath.Join(cwd, DATA_DIR, "glyph.json"))
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

func fetchRelics() (Relics, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return Relics{}, err
	}

	f, err := os.ReadFile(filepath.Join(cwd, DATA_DIR, "relic.json"))
	if err != nil {
		return Relics{}, err
	}

	relics := Relics{}
	err = json.Unmarshal([]byte(f), &relics)
	if err != nil {
		return Relics{}, err
	}
	return relics, nil
}

func fetchBodies() (Bodies, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return Bodies{}, err
	}

	f, err := os.ReadFile(filepath.Join(cwd, DATA_DIR, "body.json"))
	if err != nil {
		return Bodies{}, err
	}

	bodies := Bodies{}
	err = json.Unmarshal([]byte(f), &bodies)
	if err != nil {
		return Bodies{}, err
	}
	return bodies, nil
}

func fetchGenerics(fileName string) ([]Generic, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return []Generic{}, err
	}

	f, err := os.ReadFile(filepath.Join(cwd, DATA_DIR, fileName+".json"))
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

func fetchChaos() (Chaos, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return Chaos{}, err
	}

	f, err := os.ReadFile(filepath.Join(cwd, DATA_DIR, "chaos.json"))
	if err != nil {
		return Chaos{}, err
	}

	chaos := Chaos{}
	err = json.Unmarshal([]byte(f), &chaos)
	if err != nil {
		return Chaos{}, err
	}
	return chaos, nil
}

func fetchMutations() ([]Generic, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return []Generic{}, err
	}

	f, err := os.ReadFile(filepath.Join(cwd, DATA_DIR, "mutation.json"))
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

	// TODO: update to 2 average after 13*count
	return int(math.Floor(float64(totalResult-(13*playerCount))/float64(playerCount)) + 1), nil
}

var DIE_SIZES []float64 = []float64{2.5, 3.5, 4.5, 5.5, 6.5}
var DIE_FACES []string = []string{"d4", "d6", "d8", "d10", "d12"}

func dmgToDice(dmg float64) string {
	var lowestDie string
	lowestMod := -1.0
	lowestCount := 1.0
	for idx, size := range DIE_SIZES {
		mod := math.Mod(dmg, size)
		if mod <= lowestMod || lowestMod < 0 {
			count := math.Floor(dmg / size)
			face := DIE_FACES[idx]
			if (face == "d4" && count > 0 && count <= 5) || (face != "d4" && count > 0 && count <= 10) {
				lowestDie = face
				lowestMod = mod
				lowestCount = count
			}
		}
	}
	if lowestMod < 0 {
		return fmt.Sprintf("No elegant dice set found (1-10 dice) for %.1f damage.", dmg)
	}
	return fmt.Sprintf("%.0f%s", lowestCount, lowestDie)
}
