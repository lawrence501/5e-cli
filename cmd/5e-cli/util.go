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

func discordSend(content string) error {
	// sendP := promptui.Prompt{
	// 	Label: "Broadcast to Discord? (Empty for yes, anything for no)",
	// }
	// send, err := sendP.Run()
	// if err != nil {
	// 	return err
	// }
	log.Print(content)
	// if send != "" {
	// 	return nil
	// }

	// content = fmt.Sprintf("> *%s*\n%s", randomWords(2, " "), content)
	// DISCORD_USERNAME := "Broadcaster"
	// message := discordwebhook.Message{
	// 	Username: &DISCORD_USERNAME,
	// 	Content:  &content,
	// }
	// return discordwebhook.SendMessage(
	// 	"https://discord.com/api/webhooks/1340159157164179489/auj9LBrU9CIGtxE6BE1-lhHJcTIL0OQkwK6oT9i7KwmFamUyEZ311fHS5S66aN4a-l1U",
	// 	message,
	// )
	return nil
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

var DIE_SIZES []float64 = []float64{2.5, 3.5, 4.5, 5.5, 6.5}
var DIE_FACES []string = []string{"d4", "d6", "d8", "d10", "d12"}

func dmgToDice(dmg float64) string {
	// fmt.Printf("Finding dice for %.5f damage...\n", dmg)
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
				// fmt.Println("Exact match")
				bestCount = count
				bestFace = face
				lowestDiff = diff
				break
			}
			if diff >= 2.5 {
				continue
			}
			// fmt.Printf("Count: %d, Face: d%d, Average: %.5f, Diff: %.5f\n", count, face, average, diff)

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

func randomWords(count int, separator string) string {
	outputBuilder := []string{}
	for _ = range count {
		outputBuilder = append(outputBuilder, randSelect(WORDLIST))
	}
	return strings.Join(outputBuilder, separator)
}
