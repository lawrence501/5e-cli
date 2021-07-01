package main

var COLOUR_UPGRADE_DESCRIPTIONS map[string]string = map[string]string{
	"1":  "Trap: Reroll with -1 colour and disadv",
	"2":  "1d20gp: +50% gold",
	"3":  "Mundane: Sells for +25gp",
	"4":  "Wondrous: Adv on rarity",
	"5":  "1E: +1 upgrade point",
	"6":  "1d20 + 50gp: +50% gold",
	"7":  "Essence: Adv on dmg type",
	"8":  "2E: +1 upgrade point",
	"11": "Amulet: +1 starting tier",
	"10": "Colour reroll: +1 colour",
	"9":  "1d20 + 150gp: +50% gold",
	"12": "Body armour: Adv on mod",
	"13": "Ring: Adv on mod",
	"14": "Double value 1E: Adv on mod",
	"15": "3E: +1 upgrade point",
	"16": "Belt: Adv on mod",
	"17": "Crafting stone: enter stone prefix",
	"18": "Dream mirror: +1 starting lvl",
	"19": "Glyph: Counts as +1 glyph",
	"20": "Relic: +1 starting lvl",
}

var DAMAGE_TYPES []string = []string{
	"acid",
	"bludgeoning",
	"cold",
	"fire",
	"force",
	"lightning",
	"necrotic",
	"piercing",
	"poison",
	"psychic",
	"radiant",
	"slashing",
	"thunder",
}

type Generic struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Traps struct {
	Standard []Generic `json:"standard"`
	Crit     []Generic `json:"crit"`
}

type MundanesAll struct {
	Standard Mundanes `json:"standard"`
	Crit     Mundanes `json:"crit"`
}

type Mundanes struct {
	Weapon []Mundane `json:"weapon"`
	Armour []Mundane `json:"armour"`
}

type Mundane struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

type Wondrous struct {
	Common    []Generic `json:"common"`
	Uncommon  []Generic `json:"uncommon"`
	Rare      []Generic `json:"rare"`
	VeryRare  []Generic `json:"very rare"`
	Legendary []Generic `json:"legendary"`
}

type Enchant struct {
	Description string   `json:"description"`
	PointValue  string   `json:"pointValue"`
	Upgrade     string   `json:"upgrade"`
	Tags        []string `json:"tags"`
}
