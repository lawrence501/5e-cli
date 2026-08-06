package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"os"
	"path/filepath"
)

// DATA_DIR is the default directory name holding the JSON data files.
const DATA_DIR = "data"

// dataDir is the resolved path to the JSON data directory, set from the -data
// flag (defaulting to a `data` directory beside the executable or the cwd).
var dataDir = DATA_DIR

func randSelect[S []E, E any](s S) E {
	return s[rand.Intn(len(s))]
}

func fetchData[T any](fileName string, _ T) (T, error) {
	var data T

	f, err := os.ReadFile(filepath.Join(dataDir, fileName+".json"))
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(f, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func generateWeather() (string, error) {
	weathers, err := fetchData("weather", []string{})
	if err != nil {
		return "", err
	}
	return randSelect(weathers), nil
}

func fetchDreamPool(char string) ([]Affix, error) {
	pools, err := fetchData("dreamPool", map[string][]Affix{})
	if err != nil {
		return []Affix{}, err
	}
	return pools[char], nil
}

func fetchAugments(char string) ([]Augment, error) {
	augments, err := fetchData("augment", map[string][]Augment{})
	if err != nil {
		return []Affix{}, err
	}
	return augments[char], nil
}

var DIE_SIZES []float64 = []float64{2.5, 3.5, 4.5, 5.5, 6.5}
var DIE_FACES []string = []string{"d4", "d6", "d8", "d10", "d12"}

func dmgToDice(dmg float64) string {
	bestCount := math.MaxInt
	bestFace := 0
	lowestDiff := math.MaxFloat64

	for count := 1; count <= 15; count++ {
		for face := 4; face <= 12; face += 2 {
			if face == 4 && count > 5 {
				continue
			}
			average := float64(count) * (float64(face + 1)) / 2.0
			diff := math.Abs(float64(average) - dmg)
			if diff == 0.0 {
				bestCount = count
				bestFace = face
				lowestDiff = diff
				break
			}
			if diff >= 2.5 {
				continue
			}

			if diff < lowestDiff || (diff == lowestDiff && count < bestCount) {
				lowestDiff = diff
				bestCount = count
				bestFace = face
			}
		}
		if lowestDiff == 0.0 {
			break
		}
	}

	if bestFace == 0 {
		return fmt.Sprintf("No elegant dice set found (1-15 dice) for %.1f damage.", dmg)
	}
	return fmt.Sprintf("%dd%d", bestCount, bestFace)
}
