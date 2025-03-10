package main

var COLOUR_UPGRADE_DESCRIPTIONS map[string]string = map[string]string{
	"tarot":     "Tarot: +1 card draw",
	"amulet":    "Amulet: +1 option when upgrading",
	"equipment": "Equipment: +1 upgrade point randomly allocated",
	"mirror":    "Dream Mirror: +1 upgrade point on either new mod or upgrade/reroll",
	"glyph":     "Glyph: Grants +1 tier",
	"crystal":   "Crystal: +5% flare chance",
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

var ABILITIES []string = []string{
	"Strength",
	"Dexterity",
	"Intelligence",
	"Wisdom",
	"Charisma",
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
	"axe",
	"spear",
	"caster",
	"dart",
	"bow",
	"sling",
	"flail",
	"polearm",
	"sword",
	"hammer",
	"targe",
	"pick",
	"trap",
}

var SKILLS []string = []string{
	"athletics",
	"brawling",
	"finesse",
	"stealth",
	"magiscience",
	"history",
	"investigation",
	"nature",
	"insight",
	"medicine",
	"perception",
	"survival",
	"deception",
	"intimidation",
	"persuasion",
}

var DAMAGE_POLARITIES []string = []string{
	"physical",
	"non-physical",
}

var PARTY_MEMBERS []string = []string{
	"Dekel",
	"Bentley",
	"Ede",
}

var FLARE_CHANCES map[string]int = map[string]int{
	"Dekel":   10,
	"Bentley": 10,
	"Ede":     10,
}

var CITIES []string = []string{
	"The Hub",
}

var INSIGHTS map[string]int = map[string]int{
	"Dekel":    0,
	"Bentley":  0,
	"Ede":      0,
	"Sidekick": 0,
}

var LIGHT_TYPES []string = []string{
	"sunlight",
	"darkness",
}

var CONDITIONS []string = []string{
	"blinded",
	"charmed",
	"confused",
	"dazed",
	"deafened",
	"debilitated",
	"dominated",
	"doomed",
	"fatigue",
	"fixated",
	"frightened",
	"grappled",
	"incapacitated",
	"poisoned",
	"prone",
	"rattled",
	"restrained",
	"slowed",
	"sluggish",
	"staggered",
	"strife",
	"taunted",
	"unconscious",
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
	"angled (smaller/larger)",
	"charge",
	"deadly",
	"delayed",
	"deployed",
	"explosive",
	"fatal",
	"forceful",
	"mounted",
	"non-lethal",
	"offhand (your main hand weapon's damage die increases by 2 sizes)",
	"opener",
	"parrying",
	"prepared",
	"rapid",
	"reach",
	"responsive",
	"restraining",
	"returning",
	"sweep",
	"targe",
	"thrown (20'/60')",
	"triggered (attack the triggering creature with this weapon)",
	"two-handed",
	"unarmed",
}

var SIZE_DIFFERENCES []string = []string{
	"larger",
	"smaller",
}

var INITIATIVE_DIFFERENCES []string = []string{
	"ahead of",
	"behind",
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
	"ranged",
	"area",
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
	"helmet",
	"body",
	"gloves",
	"boots",
}

var FEATS []string = []string{
	"alert",
	"athlete",
	"blinktouched",
	"brawler",
	"charger",
	"crippler",
	"decayer",
	"defensive duelist",
	"dual wielder",
	"dungeon delver",
	"durable",
	"eldritch adept (req spell)",
	"fighting initiate",
	"grappler",
	"great weapon master",
	"healer",
	"heavy armour master",
	"inspiring leader",
	"keen mind",
	"light armour master",
	"magic initiate",
	"martial scholar",
	"medium armour master",
	"metamagic adept (req spell)",
	"mounted combatant",
	"reflective",
	"resilient",
	"sentinel",
	"sharpshooter",
	"shield master",
	"skilled",
	"skulker",
	"socialite",
	"specialist",
	"spell touched",
	"summoner",
	"survivor",
	"tactician",
	"telekinetic",
	"war caster (req spell)",
	"warlord",
}

var SIMPLE_WEAPONS []string = []string{
	"club",
	"dagger",
	"gauntlets",
	"greatclub",
	"handaxe",
	"mace",
	"quarterstaff",
	"sickle",
	"spear",
	"arcane staff",
	"blowgun",
	"light crossbow",
	"sling",
}

var MARTIAL_WEAPONS []string = []string{
	"battleaxe",
	"boomerang",
	"bullwhip",
	"dueling dagger",
	"flail",
	"glaive",
	"greataxe",
	"greatsword",
	"halberd",
	"javelin",
	"lance",
	"light hammer",
	"longsword",
	"maul",
	"punching dagger",
	"rapier",
	"ring blade",
	"sabre",
	"scimitar",
	"scythe",
	"shortsword",
	"starknife",
	"tower shield",
	"trident",
	"warhammer",
	"warpick",
	"whip",
	"composite bow",
	"dart",
	"hand crossbow",
	"heavy crossbow",
	"longbow",
	"net",
	"shortbow",
	"shrapnel trap",
}

var LANGUAGES []string = []string{
	"common",
	"dwarvish",
	"elvish",
	"giant",
	"gnomish",
	"goblin",
	"halfling",
	"orc",
	"abyssal",
	"celestial",
	"draconic",
	"deep speech",
	"infernal",
	"primordial",
	"sylvan",
	"undercommon",
}

var PLANES []string = []string{
	"material plane",
	"ethereal plane",
	"astral plane",
	"realm of the dead",
	"elemental plane of air",
	"elemental plane of earth",
	"elemental plane of fire",
	"elemental plane of water",
	"nine hells",
	"seven heavens",
	"twisting canopy of the eternal forest",
	"brilliant board of the cosmic game",
	"warping stairways of the colossal tower",
	"infinite rooms of the black cathedral",
	"grand halls of the shimmering library",
	"caverns of flesh",
	"blank canvas",
	"eternal slumber",
	"cycling swamp",
}

var SPELL_LISTS []string = []string{
	"bard",
	"cleric",
	"druid",
	"herald",
	"sorcerer",
	"warlock",
	"wizard",
	"artificer",
}

var MARTIAL_TRADITIONS []string = []string{
	"adamant mountain",
	"biting zephyr",
	"mirror's glint",
	"mist and shade",
	"rapid current",
	"razor's edge",
	"sanguine knot",
	"spirited steed",
	"tempered iron",
	"tooth and claw",
	"unending wheel",
	"arcane knight",
	"beast unity",
	"comedic jabs",
	"eldritch blackguard",
	"gallant heart",
}

var AFFINITIES []string = []string{
	"control",
	"accuracy",
	"mobility",
	"resource",
	"damage",
	"survivability",
	"wealth",
	"utility",
}

var INACTIVE_DMS []string = []string{
	"Dekel",
	"Bentley",
}

var SIMULATION_TYPES []string = []string{
	"dungeon",
	"hunt",
	"journey",
	"puzzle",
}

var CHANCE_RESULTS []string = []string{
	"+1 affix pair",
	"+2 affix pairs",
	"+3 affix pairs",
	"unique",
}

var CORRUPTION_RESULTS []string = []string{
	"do nothing",
	"reroll to [mundane, 1 affix pair, 2 affix pairs, 3 affix pairs, unique]",
	"add 1 affix pair",
	"give each player a chaotic modifier, without telling them until they trigger",
}
