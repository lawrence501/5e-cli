package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/manifoldco/promptui"
)

var COMMAND_MAP = map[string]func() error{
	"exit":          func() error { os.Exit(0); return nil },
	"q":             func() error { os.Exit(0); return nil },
	"1":             trap,
	"7":             book,
	"3":             mundane,
	"4":             wondrous,
	"6":             singleEnchant,
	"5":             essence,
	"9":             cards,
	"14":            doubleEnchant,
	"11":            highGold,
	"10":            func() error { log.Println("Reroll twice/upgrade result with +1 colour!"); return nil },
	"12":            amulet,
	"2":             blessing,
	"13":            ring,
	"8":             doubleValueSingleEnchant,
	"16":            tripleEnchant,
	"15":            doubleValueDoubleEnchant,
	"17":            craftingStone,
	"18":            func() error { log.Println("Dream Mirror"); return nil },
	"19":            func() error { log.Println("Glyph"); return nil },
	"20":            relic,
	"colour":        colour,
	"wep":           weaponEnchant,
	"arm":           armourEnchant,
	"ring":          upgradeRing,
	"glyph":         glyph,
	"relic":         upgradeRelic,
	"skill":         skill,
	"dmg type":      dmgType,
	"creature type": creatureType,
	"card":          singleCard,
	"ability score": abilityScore,
	"loot":          loot,
	"harvest":       harvest,
	"condi":         condition,
	"mutate":        mutation,
	"tpk":           tpk,
	"plains":        plains,
	"forest":        forest,
	"mountain":      mountain,
	"aquatic":       aquatic,
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

		err = COMMAND_MAP[input]()
		if err != nil {
			log.Printf("Error occurred during running of %s", input)
			log.Fatal(err)
		}
		log.Println()
	}
}
