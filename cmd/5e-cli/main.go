package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// COMMAND_MAP maps a command name (as passed on the CLI, spaces allowed) to the
// generator that produces its view-model.
var COMMAND_MAP = map[string]CommandFunc{
	"affix":            affix,
	"target affix":     targetAffix,
	"colour":           colour,
	"glyph":            glyph,
	"mutate":           mutate,
	"insight":          insight,
	"dmg up":           dmgUpgrade,
	"chaos":            chaos,
	"combat":           combat,
	"travel":           travel,
	"dream":            dream,
	"mission":          mission,
	"crystal":          crystal,
	"crystal power":    crystalPower,
	"amulet":           amulet,
	"prize":            prize,
	"tarot":            tarot,
	"relic":            relic,
	"loot":             loot,
	"loot-roll":        lootRoll,
	"enchant":          enchantedItem,
	"shrine":           shrine,
	"carnival":         carnival,
	"npc":              npc,
	"skill":            skill,
	"dmg type":         dmgType,
	"creature type":    creatureType,
	"ability":          ability,
	"condi":            condition,
	"dmg polarity":     dmgPolarity,
	"party member":     partyMember,
	"xiloan":           xiloan,
	"weapon class":     weaponClass,
	"phys type":        physType,
	"non-phys type":    nonPhysType,
	"class":            class,
	"feat":             feat,
	"simple wep":       simpleWeapon,
	"martial wep":      martialWeapon,
	"language":         language,
	"plane":            plane,
	"affinity":         affinity,
	"trait":            weaponTrait,
	"loot result":      lootResult,
	"journey activity": journeyActivity,
	"augment":          augment,
}

func main() {
	flag.StringVar(&dataDir, "data", defaultDataDir(), "path to the JSON data directory")
	flag.Usage = usage
	flag.Parse()

	cmdName := strings.Join(flag.Args(), " ")
	if cmdName == "" {
		usage()
		os.Exit(2)
	}

	fn, ok := COMMAND_MAP[cmdName]
	if !ok {
		fmt.Fprintf(os.Stderr, "unknown command: %q\n\n", cmdName)
		usage()
		os.Exit(2)
	}

	req, err := readRequest()
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not read request from stdin: %v\n", err)
		os.Exit(1)
	}

	vm, err := fn(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", cmdName, err)
		os.Exit(1)
	}

	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(vm); err != nil {
		fmt.Fprintf(os.Stderr, "could not encode result: %v\n", err)
		os.Exit(1)
	}
}

// readRequest parses the JSON request context from stdin. An empty or absent
// stdin (the common case for input-free generators) yields a zero Request.
func readRequest() (Request, error) {
	var req Request
	info, err := os.Stdin.Stat()
	if err != nil {
		return req, nil
	}
	// Nothing piped in (interactive terminal): skip reading to avoid blocking.
	if info.Mode()&os.ModeCharDevice != 0 {
		return req, nil
	}
	dec := json.NewDecoder(os.Stdin)
	if err := dec.Decode(&req); err != nil {
		// Treat an empty stream as "no inputs" rather than an error.
		if err.Error() == "EOF" {
			return Request{}, nil
		}
		return req, err
	}
	return req, nil
}

// defaultDataDir prefers a `data` directory next to the executable (so the
// shipped binary finds its data regardless of the caller's working directory),
// falling back to a relative `data` path.
func defaultDataDir() string {
	if exe, err := os.Executable(); err == nil {
		d := filepath.Join(filepath.Dir(exe), DATA_DIR)
		if info, err := os.Stat(d); err == nil && info.IsDir() {
			return d
		}
	}
	return DATA_DIR
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: 5e-cli [-data DIR] <command>\n\n")
	fmt.Fprintf(os.Stderr, "Reads an optional JSON request {\"inputs\":{...}} on stdin and writes a\n")
	fmt.Fprintf(os.Stderr, "JSON view-model on stdout.\n\nCommands:\n")
	names := make([]string, 0, len(COMMAND_MAP))
	for name := range COMMAND_MAP {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Fprintf(os.Stderr, "  %s\n", name)
	}
}
