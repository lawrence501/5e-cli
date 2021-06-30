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
