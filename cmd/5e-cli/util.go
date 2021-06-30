package main

import (
	"encoding/json"
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
