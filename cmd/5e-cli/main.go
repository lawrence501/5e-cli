package main

import (
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/manifoldco/promptui"
)

var ROLL_RANGE_CEILINGS = map[int]func() error{
	9:   func() error { log.Println("Reroll twice and upgrade results with +1 colour!"); return nil },
	16:  wondrous,
	23:  func() error { log.Printf("Soul gem: %s", GEM_TAGS[rand.Intn(len(GEM_TAGS))]); return nil },
	30:  func() error { log.Printf("Crystal: %s", CREATURE_TYPES[rand.Intn(len(CREATURE_TYPES))]); return nil },
	37:  belt,
	44:  ring,
	51:  amulet,
	58:  shrine,
	65:  body,
	72:  tome,
	79:  func() error { log.Println("2x Tarot Cards"); return nil },
	86:  relic,
	93:  func() error { log.Println("Dream Mirror"); return nil },
	100: func() error { log.Println("Glyph"); return nil },
}

var COMMAND_MAP = map[string]func() error{
	"exit":          func() error { os.Exit(0); return nil },
	"q":             func() error { os.Exit(0); return nil },
	"colour":        colour,
	"wep":           weaponEnchant,
	"arm":           armourEnchant,
	"glyph":         glyph,
	"relic":         upgradeRelic,
	"skill":         skill,
	"dmg type":      dmgType,
	"creature type": creatureType,
	"ability":       ability,
	"loot":          loot,
	"harvest":       harvest,
	"condi":         condition,
	"mutate":        mutation,
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
	"empower":       empower,
	"class":         class,
	"tarot":         tarot,
	"gem":           gem,
	"craft":         craft,
	"target craft":  targetCraft,
	"dmg upgrade":   dmgUpgrade,
	"activity":      activity,
	"amulet":        amulet,
	"relic new":     relic,
	"chaos":         chaos,
	"wondrous":      wondrous,
	"ring":          ring,
	"combat":        combat,
}

func main() {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	rand.Seed(time.Now().UnixNano())

	for true {
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
