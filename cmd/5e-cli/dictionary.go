package main

var COLOUR_UPGRADE_DESCRIPTIONS map[string]string = map[string]string{
	"1":          "Trap: Reroll with -1 colour and disadv",
	"3":          "1d20gp: +50% gold",
	"4":          "Mundane: Sells for +25gp",
	"5":          "Wondrous: Adv on rarity",
	"6":          "1E: +1 upgrade point",
	"7":          "1d20 + 50gp: +50% gold",
	"8":          "Essence: Adv on dmg type",
	"9":          "2E: +1 upgrade point",
	"12":         "Amulet: +1 starting tier",
	"10":         "Colour reroll: +1 colour",
	"11":         "1d20 + 150gp: +50% gold",
	"2":          "Blessing: +1 tier",
	"13":         "Ring: Adv on mod",
	"14":         "Double value 1E: Adv on mod",
	"15":         "3E: +1 upgrade point",
	"16":         "Belt: Adv on mod",
	"17":         "Crafting stone: enter stone prefix",
	"18":         "Dream mirror: +1 starting lvl",
	"19":         "Glyph: Counts as +1 glyph",
	"20":         "Relic: +1 starting lvl",
	"fate":       "Fate stone: Adv on rerolls",
	"empowering": "Empowering stone: +1 upgrade point",
	"enhancing":  "Enhancing stone: Adv on mod",
	"attuning":   "Attuning stone: +1 starting lvl",
	"learning":   "Learning stone: +1 relic upgrade choice",
}

var PHYS_TYPES []string = []string{
	"bludgeoning",
	"piercing",
	"slashing",
}

var NON_PHYS_TYPES []string = []string{
	"acid",
	"cold",
	"fire",
	"force",
	"lightning",
	"necrotic",
	"poison",
	"psychic",
	"radiant",
	"thunder",
}

var DAMAGE_TYPES []string = append(PHYS_TYPES, NON_PHYS_TYPES...)

var RING_TAGS []string = []string{
	"ability score",
	"temp HP",
	"AB",
	"phys res",
	"speed",
}

var ABILITY_SCORES []string = []string{
	"strength",
	"dexterity",
	"constitution",
	"intelligence",
	"wisdom",
	"charisma",
}

var HIT_FORMS []string = []string{
	"weapon",
	"spell",
}

var WILL_ABILITIES []string = []string{
	"intelligence",
	"wisdom",
	"charisma",
}

var WEAPON_HANDS []string = []string{
	"one-handed",
	"two-handed",
}

var WEAPON_CLASSES []string = []string{
	"club",
	"knife",
	"brawling",
	"spear",
	"caster",
	"dart",
	"bow",
	"sling",
	"sword",
	"axe",
	"flail",
	"polearm",
	"pick",
	"hammer",
	"shield",
}
