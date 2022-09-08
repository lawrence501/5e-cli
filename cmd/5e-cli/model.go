package main

type Generic struct {
	Name        string `json:"name"`
	Description string `json:"description"`
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

type Relics struct {
	Weapon []Relic `json:"weapon"`
	Armour []Relic `json:"armour"`
}

type Encounters struct {
	Hostile  []string `json:"hostile"`
	Positive []string `json:"positive"`
}

type Tome struct {
	Name        string `json:"name"`
	Target      string `json:"target"`
	Description string `json:"description"`
}

type Cards struct {
	Common   []Card `json:"common"`
	Uncommon []Card `json:"uncommon"`
	Rare     []Card `json:"rare"`
}

type Card struct {
	Name   string `json:"name"`
	Set    string `json:"set"`
	Rarity string
}

type Wondrous struct {
	Common    []Generic `json:"common"`
	Uncommon  []Generic `json:"uncommon"`
	Rare      []Generic `json:"rare"`
	VeryRare  []Generic `json:"very rare"`
	Legendary []Generic `json:"legendary"`
}

type Rings struct {
	Uncommon  []Ring `json:"uncommon"`
	Rare      []Ring `json:"rare"`
	VeryRare  []Ring `json:"very rare"`
	Legendary []Ring `json:"legendary"`
	Artifact  []Ring `json:"artifact"`
}

type Ring struct {
	Name    string   `json:"name"`
	Effects []string `json:"effects"`
}

type Mutation struct {
	Powerful    []Generic `json:"powerful"`
	Beneficial  []Generic `json:"beneficial"`
	Distinctive []Generic `json:"distinctive"`
	Harmful     []Generic `json:"harmful"`
}

type Enchant struct {
	Description string   `json:"description"`
	PointValue  string   `json:"pointValue"`
	Upgrade     string   `json:"upgrade"`
	Tags        []string `json:"tags"`
}

type AmuletSet struct {
	Name    string   `json:"name"`
	Amulets []Amulet `json:"amulets"`
}

type Amulet struct {
	Name   string `json:"name"`
	Effect string `json:"effect"`
}

type SimpleGeneric struct {
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

type Bodies struct {
	Unarmoured []Body `json:"unarmoured"`
	Light      []Body `json:"light"`
	Medium     []Body `json:"medium"`
	Heavy      []Body `json:"heavy"`
}

type Body struct {
	Name      string          `json:"name"`
	Mods      []string        `json:"mods"`
	Variables []BodyVariables `json:"variables"`
}

type BodyVariables struct {
	Mods      []string `json:"mods"`
	Variables []string `json:"variables"`
}
