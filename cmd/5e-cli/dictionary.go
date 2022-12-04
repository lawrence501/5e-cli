package main

var COLOUR_UPGRADE_DESCRIPTIONS map[string]string = map[string]string{
	"wondrous":  "Wondrous: +1 adv on rarity",
	"tarot":     "Tarot: +1 card draw",
	"gem":       "Soul gem: +1 adv on tag",
	"amulet":    "Amulet: +1 adv on set",
	"equipment": "Equipment: +1 upgrade point",
	"ring":      "Ring: +1 adv on rarity",
	"shrine":    "Shrine: +1 reroll of proposed outcome",
	"body":      "Body: +1 adv",
	"tome":      "Tome: +1 adv",
	"mirror":    "Dream Mirror: +1 customisation option",
	"glyph":     "Glyph: Counts as +1 glyph",
	"belt":      "Belt: +1 adv on base activity",
	"crystal":   "Crystal: +1 adv on creature type",
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

var GEM_TAGS []string = []string{
	"ability",
	"temp hp",
	"type res",
	"speed",
	"weapon class",
	"1h",
	"2h",
	"shield",
	"wep dmg",
	"spell dmg",
	"type dmg",
	"phys dmg",
	"non-phys dmg",
	"minion off",
	"minion def",
	"phys res",
	"non-phys res",
	"dot",
	"debuff",
	"buff",
	"crit",
	"aoe",
	"melee",
	"proj",
	"pers area",
	"conc",
	"hp",
	"save",
	"ac",
}

var ABILITIES []string = []string{
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
	"Fillipo":    5,
	"Eva":        9,
	"Vinton":     -1,
	"Tymm":       7,
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
	"fatigue",
	"strife",
	"sluggish",
	"dominate",
	"rattle",
	"taunt",
	"confuse",
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
	"fatal (+2 die size)",
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

var TAROT_CARDS []string = []string{
	"fool",
	"magician",
	"high priestess",
	"empress",
	"emperor",
	"hierophant",
	"lovers",
	"chariot",
	"justice",
	"hermit",
	"wheel of fortune",
	"strength",
	"hanged man",
	"death",
	"temperance",
	"devil",
	"tower",
	"stars",
	"moon",
	"sun",
	"judgement",
	"world",
	"wands",
	"cups",
	"swords",
	"pentacles",
}

var EQUIP_SLOTS []string = []string{
	"onhand",
	"offhand",
	"helmet",
	"gloves",
	"pants",
	"boots",
}

var JOURNEY_ACTIVITIES []string = []string{
	"befriend",
	"busk",
	"chronicle",
	"entertain",
	"march",
	"gather",
	"gossip",
	"harvest",
	"pray",
	"rob",
	"scout",
	"shelter",
}
