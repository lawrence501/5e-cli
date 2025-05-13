package main

import (
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/manifoldco/promptui"
)

var ROLL_RANGE_CEILINGS = map[int]func() error{
	8:   func() error { log.Println("Reroll and upgrade result with +1 colour!"); return nil },
	21:  amulet,
	34:  crystal,
	47:  func() error { log.Println("Dream Mirror"); return nil },
	60:  enchantedItem,
	73:  func() error { log.Println("Glyph"); return nil },
	86:  relic,
	99:  func() error { log.Println("2x Tarot Cards"); return nil },
	100: func() error { log.Println("Player's choice and upgrade result with +1 colour!"); return nil },
}

var COMMAND_MAP = map[string]func() error{
	"exit":             func() error { os.Exit(0); return nil },
	"q":                func() error { os.Exit(0); return nil },
	"colour":           colour,
	"affix":            affix,
	"glyph":            glyph,
	"relic up":         upgradeRelic,
	"skill":            skill,
	"dmg type":         dmgType,
	"creature type":    creatureType,
	"ability":          ability,
	"condi":            condition,
	"insight":          insight,
	"dmg polarity":     dmgPolarity,
	"party member":     partyMember,
	"npc":              npc,
	"xiloan":           xiloan,
	"prize":            prize,
	"weapon class":     weaponClass,
	"phys type":        physType,
	"non-phys type":    nonPhysType,
	"class":            class,
	"tarot":            tarot,
	"dmg up":           dmgUpgrade,
	"amulet":           amulet,
	"relic":            relic,
	"chaos":            chaos,
	"combat":           combat,
	"feat":             feat,
	"simple wep":       simpleWeapon,
	"martial wep":      martialWeapon,
	"carni":            carnival,
	"language":         language,
	"dream":            dream,
	"plane":            plane,
	"mutate":           mutate,
	"mission":          mission,
	"affinity":         affinity,
	"trait":            weaponTrait,
	"enchant":          enchantedItem,
	"sim":              simulations,
	"sim affix":        simulationAffix,
	"sim type":         simulationType,
	"travel":           travel,
	"chance":           orbOfChance,
	"corrupt":          corruptionOrb,
	"augment":          augment,
	"target affix":     targetAffix,
	"crystal power":    crystalPower,
	"crystal":          crystal,
	"loot result":      lootResult,
	"journey activity": journeyActivity,
	"fumble":           fumble,
}

func main() {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	for {
		log.Println()
		baseP := promptui.Prompt{
			Label:    "Command",
			Validate: validateBase,
		}
		input, err := baseP.Run()
		if err != nil {
			log.Fatal(err)
			return
		}

		inputInt, err := strconv.Atoi(input)
		if err == nil {
			ceilings := make([]int, len(ROLL_RANGE_CEILINGS))
			i := 0
			for k := range ROLL_RANGE_CEILINGS {
				ceilings[i] = k
				i++
			}
			sort.Ints(ceilings)
			for _, c := range ceilings {
				if inputInt <= c {
					err = ROLL_RANGE_CEILINGS[c]()
					break
				}
			}
		} else {
			err = COMMAND_MAP[input]()
		}
		if err != nil {
			log.Printf("Error occurred during running of %s", input)
			log.Fatal(err)
		}
		log.Println()
	}
}
