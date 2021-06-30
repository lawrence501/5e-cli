package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/manifoldco/promptui"
)

var COMMAND_MAP = map[string]func() error{
	"exit": func() error { os.Exit(0); return nil },
	"gold": gold,
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
			log.Printf("Error occurred during running of %s\n", input)
			log.Fatal(err)
		}
		log.Println()
	}
}
