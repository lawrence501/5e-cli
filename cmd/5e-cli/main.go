package main

import (
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"

	"github.com/manifoldco/promptui"
)

var ROLL_RANGE_CEILINGS = map[int]func() error{
	11:  func() error { log.Println("Reroll and upgrade result with +1 colour!"); return nil },
	19:  mediumGold,
	27:  shrine,
	35:  func() error { log.Printf("Crystal: %s", CREATURE_TYPES[rand.Intn(len(CREATURE_TYPES))]); return nil },
	43:  magicItem,
	51:  tome,
	59:  amulet,
	67:  ring,
	75:  func() error { log.Println("2x Tarot Cards"); return nil },
	83:  relic,
	91:  func() error { log.Println("Dream Mirror"); return nil },
	99:  func() error { log.Println("Glyph"); return nil },
	100: func() error { log.Println("Player's choice and upgrade result with +1 colour!"); return nil },
}

var COMMAND_MAP = map[string]func() error{
	"exit":          func() error { os.Exit(0); return nil },
	"q":             func() error { os.Exit(0); return nil },
	"colour":        colour,
	"wep":           weaponAffix,
	"arm":           armourAffix,
	"glyph":         glyph,
	"relic":         upgradeRelic,
	"skill":         skill,
	"dmg type":      dmgType,
	"creature type": creatureType,
	"ability":       ability,
	"loot":          loot,
	"harvest":       harvest,
	"condi":         condition,
	"encounter":     randomEncounter,
	"insight":       insight,
	"dmg polarity":  dmgPolarity,
	"party member":  partyMember,
	"npc":           npc,
	"xiloan":        xiloan,
	"positive":      positiveReward,
	"weapon class":  weaponClass,
	"phys type":     physType,
	"non-phys type": nonPhysType,
	"class":         class,
	"tarot":         tarot,
	"craft":         craft,
	"target craft":  targetCraft,
	"dmg upgrade":   dmgUpgrade,
	"activity":      activity,
	"amulet":        amulet,
	"relic new":     relic,
	"chaos":         chaos,
	"ring":          ring,
	"combat":        combat,
	"travel":        travel,
	"ja":            journeyActivity,
	"feat":          feat,
	"simple wep":    simpleWeapon,
	"martial wep":   martialWeapon,
	"posi enc":      posiEnc,
	"language":      language,
	"dream":         dream,
	"plane":         plane,
	"perk":          perk,
	"mutate":        mutate,
	"follower":      follower,
	"mission":       mission,
	"affinity":      affinity,
	"trait":         weaponTrait,
	"mag":           magicItem,
	"ring upgrade":  ringUpgrade,
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
