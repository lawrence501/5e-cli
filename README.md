# 5e-cli
Golang CLI helper for my D&amp;D 5e campaigns

## Loot Rules

### Sell values

- Mundane: 5gp
- Tome: 10gp
- Amulet: 1gp
- Relic: 60gp
- Enchanted Equipment: 35gp

### Relic level costs (same price - 150)

2. 20
3. 30
4. 50
5. 80
6. 120
7. 170
8. 230
9. 300
10. 400

Bentley: 1400
Adrian: 890
Dan: 280
Lawrence: 1400

### Relic generation

1. 1st base affix
2. 2nd base affix
3. 3rd base affix
4. Thematic
5. Name Thematic
6. Random from source
7. Random PoE
8. Random affix

### Dream relic generation

1. 1st base affix
2. 2nd base affix
3. 3rd base affix
4. Thematic
5. Class
6. Backstory
7. Race
8. Feats
9. Current relics
10. Glyph path
11. o5e/Pathfinder
12. Reroll with disadv

### Dream relic customisations

- Backstory
- Race
- Any one of your feats
- Any one of your current items
- Any one of your glyph paths

### Manual value cheatsheet

Always based on dmg where possible

- Default/more: 30%
- Slight: 15%
- Substantial: 60%
- Massive: 90%
- Barely: 7%

### Point value cheatsheet

- Dominate: 8
- Stagger/Debil/exh/confused/uncon/silence/double disadv condi: 5
- Single disadv condi/sluggish/charm/root/weak: 3
- Prone/taunt/similar: 2
- Condi lasts until save/longer: x2
- Limited feature conditional: -1
- General conditional: -2
- Party positioning conditional: -1
- Common location conditional: -1
- Uncommon location conditional: -3
- Incoming healing: 2
- Passive regen: 3
- Mark effects: x2
- Glyphic 1d4 proc: 3
- Wep/spell 1d4 proc: 5
- Ignore res/imm: +3
- Ignore type-res/imm: +4
- Ignore all res/imm: +5
- Ignore res and immunity: +10
- Equipment proficiency: +3
- Language proficiency: +2
- Armour proficiency requirement: -2
- Weapon proficiency requirement: -3
- 1/LR spellcast: 1/lvl
- 1/SR spellcast: 2/lvl
- 1/combat spellcast: 3/lvl
- Learn spell: 1/lvl
- Unlimited spellcast: 5/lvl
- Minion creation: 5/CR
- A choice becomes random: -3

## Condition definitions

### Confused

A confused creature cannot take reactions.

On its turn, a confused creature rolls a d8 to determine what it does:
1-4. Debilitated until the end of its turn.
5-6. Dazed until the end of its turn.
7+. Makes a melee weapon attack against a randomly-determined creature at the start of its turn as a free action.

### Dazed

Shares condition immunity with incapacitation.

A dazed creature has disadvantage with Offence and Defence.

### Debilitated

Shares condition immunity with stun.

A debilitated creature suffers the following effects:
- Disadvantage with damaging Offence.
- All damage dealt is halved.
- Cannot take bonus actions or reactions.
- Does not grant disadvantage to adjacent ranged attackers.

### Dominated

Shares condition immunity with charm.

A dominated creature is charmed. Additionally, the dominated creature gains a free action at the start of each of its turns that can be taken to both move up to its speed and make either a single weapon attack or cast a cantrip. This free action is completely controlled by the creature that dominated the victim.

### Frightened

As before, but instead of removing a creature's ability to move towards the source of their fear, they become debilitated until the end of their turn when doing so.

### Rattled

Shares condition immunity with frighten.

A rattled creature cannot benefit from expertise dice and cannot take reactions.

### Slowed

Shares condition immunity with exhaustion.

A slowed creature's speed is halved.

### Sluggish

Shares condition immunity with exhaustion.

A sluggish creature reduces the number of attacks it can make with its Attack action by 1 (or Multiattack, where relevant).

### Staggered

Shares condition immunity with paralysis.

A staggered creature suffers the following effects:
- Incoming attack damage is always critical.
- Disadvantage with damaging Defence.

### Taunted

Shares condition immunity with charmed.

A taunted creature must target at least their taunter with all offensive actions.

## Terminology

### Splash Damage

Splash damage is always optional. It deals its damage, up to the amount of the original hit, to all creatures of your choice within 5' of the target.

## Revisited Combat Maneuvers

### Cavalier Stance

While riding your mount, your AC increases by 1.

## Ideas for next campaign/things to discuss

- Remove amulets

- New loot option: attunement-requiring magic items. Any that fit into an equipment slot already in use are reflavoured to be on another slot. Can only attune to 1 at a time. Use existing wondrous weightings of rarities, but pushed 1 up (making the 2% Artifacts). Obviously, this means they will be pretty powerful, and must therefore be high on the loot table

- Replace hard doubled mods with min 3pts (or just use D4 system? min 3 -> min 2 -> min 1 for 1 affix, 2, and 3 respectively)

- Change books such that they are modifiers on generic types of powers, eg. your sphere AoE powers are now lines with length of their diameter

- Change crafting stones to be instant-use, but allow rollback (Shrines). Put them instead of the gold-replacements for commons on the wondrous/attunement rolls. Those slots give 1 shrine, actual crafting table slot gives 2.

- Change myth cards into amulets, with set bonuses on combining the same set

- Crystals: 
    - Drop as a creature type, each starts at CR 1
    - When a crystal is generated, its power is randomly chosen from the contained creature's powers (respecting cooldowns, etc)
    - Combine a crystal with another of the same type to increase its CR by 1
    - When a crystal is tiered up, you can choose to keep the existing creature (though it will gain no benefits from the tier up)
    - Smash together 2 crystals to generate a new crystal of a type that a player is using (if anyone has none, can be anything)
    - 10% chance each turn per player to "flare", using its power. Calc on combat start with script, don't tell players other than telling them when they flare
    - The crystals count as the wearer's minions for the sake of their powers, but whenever they reference themselves, they affect the player instead

- Replace res with barrier?
    - Values increased by 3, so dmg res goes from 1 per 2pts to 2 per 1pt + single pt resistances go to 4 per pt. Single type res goes to 5 per pt, polarity res is 3 per pt.

- Drop the Looking for Trouble journey activity, force all players to always choose at least 1 journey activity per day, and make the DC for them a static 18. This allows other journey activities to become possible as your character levels up, as you may not always be able to use your primary activity when the weather, etc. is bad for it.
    - Because of redistributing the negatives from weather, even a standardly-maxed activity should drop below 50% success (+5 ability, +5 prof at L13, -4 from weather results in +6 bonus, which is 45% success chance)
    - All players MUST take at least 1 skill prof for their highest ability on char creation
    - All journey activities score a sub fail on a roll of 2.

- New ideas from Crucible:
    - Interesting selling affixes in standard pool

- Relics are not char bound, they are levelled up with gold like old days
    - Can't downlevel, can only get resources back by completely selling (100% refund)
    - Max level 6 (ie. 5 levels from base), all level costs = 100gp
    - Fully reveal on find to max level
    - 3x choice on lvl
    - Start with 3x 1pt mods, rather than 2pts, to help non-relics compete
    - 6 thematic mods per relic

- Augment system - some sort of experience and gain on breakpoints? Have class feature-esque boosts in the pool, and maybe stuff that adds synergy between backstory pairs. Every 50xp?

- Replace/remove belts now that journey activities are essentially loot slots

- New dream mirror system
    - On 1st mirror, get a random relic and choose a class feature, adding its mods to your pool (can only choose each feature once).
    - On even mirrors, get a random mod from your pool added to you, with min point value equal to the number of affixes you now have
    - On non-1st odd mirrors, choose a class feature, adding its mods to your pool (can only choose each feature once), then a random mod on you is chosen and you can either upgrade or reroll it

- Document loot mechanics

- Future loot options suggestion (d100 roll)
1. Wondrous Items (low)
2. Tomes (high)
3. Rings (med)
4. Amulets (med)
5. Shrines (med)
6. Dream Mirrors (very high)
7. Glyphs (very high)
8. Relics (high)
9. Tarots (high)
10. Crystals (low)
11. Magic Items (low)
 
8% each, remaining goes to reroll w/ upgrade (11%) + choose w/ upgrade (1% max)
