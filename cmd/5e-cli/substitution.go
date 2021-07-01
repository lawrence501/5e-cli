package main

import "math/rand"

var dmgTypeSub = func() string {
	return DAMAGE_TYPES[rand.Intn(len(DAMAGE_TYPES))]
}

var abilityScoreSub = func() string {
	return ABILITY_SCORES[rand.Intn(len(ABILITY_SCORES))]
}
