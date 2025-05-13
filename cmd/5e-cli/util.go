package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"strings"

	"github.com/gtuk/discordwebhook"
	"github.com/manifoldco/promptui"
	"golang.org/x/exp/slices"
)

const DATA_DIR = "data"

func randSelect[S []E, E any](s S) E {
	return s[rand.Intn(len(s))]
}

func fetchData[T any](fileName string, _ T) (T, error) {
	var data T

	cwd, err := os.Getwd()
	if err != nil {
		return data, err
	}

	f, err := os.ReadFile(filepath.Join(cwd, DATA_DIR, fileName+".json"))
	if err != nil {
		return data, err
	}

	err = json.Unmarshal([]byte(f), &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func genSimAffix(simType string, affixSlice []SimulationAffix) string {
	for {
		affix := randSelect(affixSlice)
		if slices.Contains(affix.SimulationTypes, strings.ToLower(simType)) {
			if affix.Display != "" {
				return fmt.Sprintf("%s *Details: %s*", affix.Display, affix.Description)
			}
			return affix.Description
		}
	}
}

func discordSend(content string) error {
	sendP := promptui.Prompt{
		Label: "Broadcast to Discord? (Empty for yes, anything for no)",
	}
	send, err := sendP.Run()
	if err != nil {
		return err
	}
	log.Print(content)
	if send != "" {
		return nil
	}

	content = fmt.Sprintf("> *%s*\n%s", randomWords(2, " "), content)
	DISCORD_USERNAME := "Broadcaster"
	message := discordwebhook.Message{
		Username: &DISCORD_USERNAME,
		Content:  &content,
	}
	return discordwebhook.SendMessage(
		"https://discord.com/api/webhooks/1340159157164179489/auj9LBrU9CIGtxE6BE1-lhHJcTIL0OQkwK6oT9i7KwmFamUyEZ311fHS5S66aN4a-l1U",
		message,
	)
}

func chooseSimulationType(simWeights SimulationWeights) (string, error) {
	dungeonChance := simWeights.Dungeon
	huntChance := simWeights.Dungeon + simWeights.Hunt
	journeyChance := simWeights.Dungeon + simWeights.Hunt + simWeights.Journey

	typeRoll := rand.Intn(100)
	var simType string
	if typeRoll < dungeonChance {
		simWeights.Dungeon -= 16
		simType = "Dungeon"
	} else if typeRoll < huntChance {
		simWeights.Hunt -= 16
		simType = "Hunt"
	} else if typeRoll < journeyChance {
		simWeights.Journey -= 16
		simType = "Journey"
	} else {
		simWeights.Puzzle -= 16
		simType = "Puzzle"
	}
	simWeights.Dungeon += 4
	simWeights.Hunt += 4
	simWeights.Journey += 4
	simWeights.Puzzle += 4
	jsonWeights, err := json.Marshal(simWeights)
	if err != nil {
		return "", err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return simType, os.WriteFile(filepath.Join(cwd, DATA_DIR, "simulationWeights.json"), jsonWeights, 0644)
}

func genSimCurrency() (string, error) {
	allCurrencies, err := fetchData("simulationCurrency", []string{})
	if err != nil {
		return "", err
	}
	outputBuilder := ""
	for range 3 {
		outputBuilder += fmt.Sprintf("* 1x %s\n", randSelect(allCurrencies))
	}
	return outputBuilder, nil
}

func generateWeather() (string, error) {
	weathers, err := fetchData("weather", []string{})
	if err != nil {
		return "", err
	}
	return randSelect(weathers), nil
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
	pools, err := fetchData("dreamPool", map[string][]Affix{})
	if err != nil {
		return []Affix{}, err
	}
	return pools[char], nil
}

var DIE_SIZES []float64 = []float64{2.5, 3.5, 4.5, 5.5, 6.5}
var DIE_FACES []string = []string{"d4", "d6", "d8", "d10", "d12"}

func dmgToDice(dmg float64) string {
	var lowestDie string
	lowestMod := -1.0
	lowestCount := 1.0
	for idx, size := range DIE_SIZES {
		if dmg < size-0.5 {
			continue
		}
		mod := math.Mod(dmg, size)
		// fmt.Println(fmt.Sprintf("Mod for %s is %f (dmg = %f, size = %f)", DIE_FACES[idx], mod, dmg, size))
		if mod <= lowestMod || lowestMod < 0 {
			count := math.Round(dmg / size)
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

func randomWords(count int, separator string) string {
	outputBuilder := []string{}
	for _ = range count {
		outputBuilder = append(outputBuilder, randSelect(WORDLIST))
	}
	return strings.Join(outputBuilder, separator)
}
