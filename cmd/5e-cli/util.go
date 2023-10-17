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

func randSelect[S []E, E any](s S) E {
	return s[rand.Intn(len(s))]
}

func fetchAffixes(fileName string) ([]Affix, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return []Affix{}, err
	}

	f, err := os.ReadFile(filepath.Join(cwd, DATA_DIR, fileName+".json"))
	if err != nil {
		return []Affix{}, err
	}

	affixes := []Affix{}
	err = json.Unmarshal([]byte(f), &affixes)
	if err != nil {
		return []Affix{}, err
	}
	return affixes, nil
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

func positiveEncounter() (string, error) {
	allEncounters, err := fetchEncounters()
	if err != nil {
		return "", err
	}
	return processMod(randSelect(allEncounters.Positive)), nil
}

func hostileEncounter(tag string) (string, error) {
	allEncounters, err := fetchEncounters()
	if err != nil {
		return "", err
	}
	var encounter HostileEncounter
	for true {
		encounter = randSelect(allEncounters.Hostile)
		if tag != "" {
			if slices.Contains(encounter.Tags, tag) {
				break
			}
		} else {
			break
		}
	}
	return fmt.Sprintf("%s (%d)", encounter.Name, encounter.ID), nil
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
		chosen := randSelect(weathers.Exotic)
		return fmt.Sprintf("%s (+2 minimum hostile random encounters [before %d and %d]. At least one must be a combat. %s)", chosen.Name, rand.Intn(5)+1, rand.Intn(5)+1, chosen.Description), nil
	}
	chosen := randSelect(weathers.Common)
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

func fetchDreamPool(char string) ([]Affix, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return []Affix{}, err
	}

	f, err := os.ReadFile(filepath.Join(cwd, DATA_DIR, "dreamPool.json"))
	if err != nil {
		return []Affix{}, err
	}

	pools := map[string][]Affix{}
	err = json.Unmarshal([]byte(f), &pools)
	if err != nil {
		return []Affix{}, err
	}
	return pools[char], nil
}

func fetchBlessings(char string) ([]string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return []string{}, err
	}

	f, err := os.ReadFile(filepath.Join(cwd, DATA_DIR, "blessing.json"))
	if err != nil {
		return []string{}, err
	}

	blessings := map[string][]string{}
	err = json.Unmarshal([]byte(f), &blessings)
	if err != nil {
		return []string{}, err
	}
	return blessings[char], nil
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

	return int(math.Floor(float64(totalResult-(13*playerCount))/float64(2.5*float64(playerCount))) + 1), nil
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
