package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/manifoldco/promptui"
)

var COMMAND_MAP = map[string]func() error{
	"exit":   func() error { os.Exit(0); return nil },
	"q":      func() error { os.Exit(0); return nil },
	"1":      trap,
	"2":      lowGold,
	"3":      mundane,
	"4":      wondrous,
	"5":      singleEnchant,
	"6":      mediumGold,
	"7":      essence,
	"8":      doubleEnchant,
	"9":      highGold,
	"colour": colour,
	"wep":    weaponEnchant,
	"arm":    armourEnchant,
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

		if _, err = strconv.Atoi(input); err == nil {
			checkColourCrit()
		}
		err = COMMAND_MAP[input]()
		if err != nil {
			log.Printf("Error occurred during running of %s", input)
			log.Fatal(err)
		}
		log.Println()
	}
}
