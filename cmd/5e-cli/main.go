package main

import (
	"fmt"
	"log"
	"os"

	"github.com/manifoldco/promptui"
)

var COMMAND_MAP = map[string]func() error{
	"exit": func() error { os.Exit(0); return nil },
	"test": test,
}

func main() {
	for true {
		basePrompt := promptui.Prompt{
			Label:    "Command",
			Validate: validateBase,
		}

		input, err := basePrompt.Run()
		if err != nil {
			log.Fatal(err)
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		err = COMMAND_MAP[input]()
		if err != nil {
			log.Printf("Error occurred during running of %s\n", input)
			log.Fatal(err)
		}
	}
}

var test = func() error {
	fmt.Printf("You are in the test function")
	return nil
}
