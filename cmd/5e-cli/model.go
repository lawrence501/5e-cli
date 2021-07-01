package main

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

type Amulet struct {
	Name    string   `json:"name"`
	Mods    []string `json:"mods"`
	Upgrade string   `json:"upgrade"`
}

type BasicRing struct {
	Description string `json:"description"`
}

type ThematicRing struct {
	Name string   `json:"name"`
	Mods []string `json:"mods"`
	Tags []string `json:"tags"`
}

type GlyphPath struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tiers       []string `json:"tiers"`
}

type Relic struct {
	Name         string    `json:"name"`
	Type         string    `json:"type"`
	StartingMods []Enchant `json:"startingMods"`
}
