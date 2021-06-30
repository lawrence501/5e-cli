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

func fetchMundanes(t string) ([]Generic, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return []Generic{}, err
	}

	f, err := ioutil.ReadFile(filepath.Join(cwd, DATA_DIR, "mundane.json"))
	if err != nil {
		return []Generic{}, err
	}

	mundanesAll := MundanesAll{}
	err = json.Unmarshal([]byte(f), &mundanesAll)
	if err != nil {
		return []Generic{}, err
	}

	var mundanes Mundanes
	switch t {
	case "standard":
		mundanes = mundanesAll.Standard
	case "crit":
		mundanes = mundanesAll.Crit
	default:
		return []Generic{}, fmt.Errorf("Invalid mundane type: %s", t)
	}

	var ret []Generic
	typeRoll := rand.Intn(5)
	if typeRoll < 2 {
		ret = mundanes.Weapon
	} else {
		ret = mundanes.Armour
	}
	return ret, nil
}

func fetchGenerics(fName string) ([]Generic, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return []Generic{}, err
	}

	f, err := ioutil.ReadFile(filepath.Join(cwd, DATA_DIR, fName+".json"))
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
