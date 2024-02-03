package main

type Generic struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Chaos struct {
	Trigger []string `json:"trigger"`
	Target  []string `json:"target"`
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
	Hostile  []HostileEncounter `json:"hostile"`
	Positive []string           `json:"positive"`
}

type HostileEncounter struct {
	Name string   `json:"name"`
	ID   int      `json:"id"`
	Tags []string `json:"tags"`
}

type Tome struct {
	Name        string `json:"name"`
	Target      string `json:"target"`
	Description string `json:"description"`
}

type Affix struct {
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

type GlyphPath struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tiers       []string `json:"tiers"`
}

type Relic struct {
	Name            string  `json:"name"`
	Type            string  `json:"type"`
	StartingAffixes []Affix `json:"startingAffixes"`
}

type Weathers struct {
	Common []string  `json:"common"`
	Exotic []Generic `json:"exotic"`
}

type RingBase struct {
	Base     string   `json:"base"`
	Affixes  []string `json:"affixes"`
	Capstone string   `json:"capstone"`
}
