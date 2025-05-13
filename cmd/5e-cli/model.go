package main

type Generic struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Chaos struct {
	Trigger []string `json:"trigger"`
	Target  []string `json:"target"`
}

type Mutation struct {
	Name        string `json:"name"`
	Target      string `json:"target"`
	Description string `json:"description"`
}

type Affix struct {
	Description string   `json:"description"`
	PointValue  string   `json:"pointValue"`
	Upgrade     string   `json:"upgrade"`
	Affinities  []string `json:"affinities"`
}

type Amulet struct {
	Name     string  `json:"name"`
	Base     Affix   `json:"base"`
	Upgrades []Affix `json:"upgrades"`
}

type GlyphPath struct {
	Theme string   `json:"theme"`
	Tiers []string `json:"tiers"`
}

type Relic struct {
	Name            string  `json:"name"`
	StartingAffixes []Affix `json:"startingAffixes"`
}

type Crystal struct {
	Name     string   `json:"name"`
	Passives []string `json:"passives"`
	Powers   []string `json:"powers"`
}

type SimulationWeights struct {
	Dungeon int `json:"dungeon"`
	Hunt    int `json:"hunt"`
	Journey int `json:"journey"`
	Puzzle  int `json:"puzzle"`
}

type SimulationAffixes struct {
	Positive []SimulationAffix `json:"positive"`
	Negative []SimulationAffix `json:"negative"`
}

type SimulationAffix struct {
	Description     string   `json:"description"`
	Display         string   `json:"display"`
	SimulationTypes []string `json:"simulationTypes"`
}

type Craft struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Tier        int    `json:"tier"`
	Upgrade     string `json:"upgrade"`
	Rank        int
}

type Fumble struct {
	Trigger string `json:"trigger"`
	Effect  string `json:"effect"`
}
