package main

var COLOUR_UPGRADE_DESCRIPTIONS map[string]string = map[string]string{
	"1":           "Trap: Reroll with -1 colour and disadv",
	"3":           "Mundane: Sells for +25gp",
	"4":           "Wondrous: Adv on rarity",
	"5":           "Essence: Adv on dmg type",
	"6":           "1E: +1 upgrade point",
	"8":           "Double value 1E: Adv on mod",
	"7":           "Book: Adv",
	"14":          "2E: +1 upgrade point",
	"12":          "Amulet: +1 starting tier",
	"10":          "Colour reroll: +1 colour",
	"11":          "1d20 + 150gp: +50% gold",
	"2":           "Blessing: +1 tier",
	"13":          "Ring: Adv on mod",
	"9":           "Myth Cards: Each card has adv rarity",
	"16":          "3E: +1 upgrade point",
	"15":          "Double value 2E: Adv on 1 mod",
	"17":          "Crafting stone: enter stone prefix",
	"18":          "Dream mirror: +1 starting lvl",
	"19":          "Glyph: Counts as +1 glyph",
	"20":          "Relic: +1 starting lvl",
	"fate":        "Fate stone: Adv on rerolls",
	"empowering":  "Empowering stone: +1 upgrade point",
	"enhancing":   "Enhancing stone: Adv on mod",
	"attuning":    "Attuning stone: +1 starting lvl",
	"learning":    "Learning stone: +1 relic upgrade choice",
	"cloning":     "Cloning stone: +1 reroll of final result",
	"negating":    "Negating stone: +1 reroll of final distribution",
	"sacrificial": "Sacrificial stone: +1 min loot roll",
	"chaotic":     "Chaotic stone: +1 upgrade point",
	"blunt":       "Blunt stone: another -1 die size, but +2x to the multi",
	"brittle":     "Brittle stone: another -1 AC, but +3x to the multi",
	"gambling":    "Gambling stone: adv reroll",
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
	"temp hp",
	"ab",
	"phys res",
	"non-phys res",
	"speed",
	"weapon class",
	"weapon hands",
	"dual-wielding",
	"holding shield",
	"wep dmg",
	"spell dmg",
	"non-phys dmg",
	"phys dmg",
	"dmg polarity",
	"minion dmg",
	"dot",
	"condi",
	"buff",
	"crit",
	"minion hp",
	"aoe",
	"proj",
	"pers area",
	"conc",
	"hit die healing",
	"max hit dice",
	"hp",
	"save",
	"wep ac",
	"spell ac",
	"melee",
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

var SKILLS []string = []string{
	"athletics",
	"acrobatics",
	"sleight of hand",
	"stealth",
	"arcana",
	"history",
	"investigation",
	"nature",
	"religion",
	"animal handling",
	"insight",
	"medicine",
	"perception",
	"survival",
	"deception",
	"intimidation",
	"performance",
	"persuasion",
}

var DAMAGE_POLARITIES []string = []string{
	"physical",
	"non-physical",
}

var PARTY_MEMBERS []string = []string{
	"Fillipo",
	"Eva",
	"Vinton",
	"Tymm",
}

var INSIGHTS map[string]int = map[string]int{
	"Fillipo":    4,
	"Eva":        8,
	"Vinton":     -1,
	"Tymm":       6,
	"Marguerita": 1,
}

var LIGHT_TYPES []string = []string{
	"sunlight",
	"darkness",
}

var CONDITIONS []string = []string{
	"blind",
	"charm",
	"deaf",
	"frighten",
	"grapple",
	"daze",
	"invisible",
	"stagger",
	"petrify",
	"poison",
	"prone",
	"restrain",
	"debilitate",
	"unconscious",
	"exhaust",
	"sluggish",
	"dominate",
	"rattle",
	"taunt",
}

var AOE_SHAPES []string = []string{
	"cone",
	"line",
	"sphere",
	"cube",
}

var SCHOOLS []string = []string{
	"conjuration",
	"necromancy",
	"evocation",
	"abjuration",
	"transmutation",
	"divination",
	"enchantment",
	"illusion",
}

var CREATURE_TYPES []string = []string{
	"aberration",
	"beast",
	"celestial",
	"construct",
	"dragon",
	"elemental",
	"fey",
	"fiend",
	"giant",
	"humanoid",
	"monstrosity",
	"ooze",
	"plant",
	"undead",
}

var HEALTH_STATUSES []string = []string{
	"healthy",
	"bloodied",
}

var WEAPON_TRAITS []string = []string{
	"rapid",
	"charge (+2 die size)",
	"compound",
	"concealable",
	"conduit",
	"deadly",
	"disarm",
	"dual-wielding",
	"fatal (+1 die size)",
	"flexible (random other)",
	"forceful",
	"grapple",
	"heavy",
	"mental",
	"non-lethal",
	"parrying",
	"propulsive",
	"loading",
	"responsive",
	"returning",
	"shove",
	"sweep",
	"targe (1)",
	"trip",
	"unarmed",
	"thrown (20'/60')",
	"finesse",
	"two-handed",
	"simple",
	"martial",
	"reach",
}

var SIZE_DIFFERENCES []string = []string{
	"larger",
	"smaller",
}

var ARMOUR_WEIGHTS []string = []string{
	"unarmoured",
	"light",
	"medium",
	"heavy",
}

var ENEMY_ARMOUR_FORMS []string = []string{
	"metal-armoured",
	"unarmoured",
}

var RANGE_TYPES []string = []string{
	"melee",
	"proj",
}

var MUTATION_TYPES []string = []string{
	"powerful",
	"beneficial",
	"distinctive",
	"harmful",
}

var RACES []string = []string{
	"dragonborn",
	"dwarf",
	"elf",
	"gnome",
	"half-elf",
	"halfling",
	"half-orc",
	"human",
	"tiefling",
}

var GENDERS []string = []string{
	"male",
	"female",
}

var XILOANS []string = []string{
	"cloud",
	"aroshi",
	"p'all",
	"mensis",
	"gawdian",
}

var CLASSES []string = []string{
	"adept",
	"bard",
	"berserker",
	"cleric",
	"druid",
	"fighter",
	"herald",
	"marshal",
	"ranger",
	"rogue",
	"sorcerer",
	"warlock",
	"wizard",
	"artificer",
	"savant",
}
