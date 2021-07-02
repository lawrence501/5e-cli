package main

import (
	"math/rand"
	"regexp"
	"strings"
)

var SUBSTITUTION_MAP = map[string][]string{
	"$dmgType":      DAMAGE_TYPES,
	"$abilityScore": ABILITY_SCORES,
	"$hitForm":      HIT_FORMS,
	"$physType":     PHYS_TYPES,
	"$willAbility":  WILL_ABILITIES,
	"$weaponHands":  WEAPON_HANDS,
	"$weaponClass":  WEAPON_CLASSES,
}

func processMod(modString string) string {
	matcher := regexp.MustCompile(`\$\S+`)
	subs := matcher.FindAllString(modString, -1)
	ret := modString
	for _, s := range subs {
		sub := SUBSTITUTION_MAP[s][rand.Intn(len(SUBSTITUTION_MAP[s]))]
		ret = strings.Replace(ret, s, sub, 1)
	}
	return ret
}
