package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
)

const DATA_DIR = "data"

func checkColourCrit() {
	if rand.Intn(100) < 5 {
		log.Println("Colour crit! +1 colour!")
	}
}

func sliceContains(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
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

func fetchMundanes(t string) ([]Mundane, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return []Mundane{}, err
	}

	f, err := ioutil.ReadFile(filepath.Join(cwd, DATA_DIR, "mundane.json"))
	if err != nil {
		return []Mundane{}, err
	}

	mundanesAll := MundanesAll{}
	err = json.Unmarshal([]byte(f), &mundanesAll)
	if err != nil {
		return []Mundane{}, err
	}

	var mundanes Mundanes
	switch t {
	case "standard":
		mundanes = mundanesAll.Standard
	case "crit":
		mundanes = mundanesAll.Crit
	default:
		return []Mundane{}, fmt.Errorf("Invalid mundane type: %s", t)
	}

	var ret []Mundane
	typeRoll := rand.Intn(5)
	if typeRoll < 2 {
		ret = mundanes.Weapon
	} else {
		ret = mundanes.Armour
	}
	return ret, nil
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

func fetchBasicRings() ([]BasicRing, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return []BasicRing{}, err
	}

	f, err := ioutil.ReadFile(filepath.Join(cwd, DATA_DIR, "basicRing.json"))
	if err != nil {
		return []BasicRing{}, err
	}

	rings := []BasicRing{}
	err = json.Unmarshal([]byte(f), &rings)
	if err != nil {
		return []BasicRing{}, err
	}
	return rings, nil
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
