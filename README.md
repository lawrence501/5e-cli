# 5e-cli
Golang CLI helper for my D&amp;D 5e campaigns

## Loot Rules

### Myth Dust

Upon finding duplicate Myth Cards, they are turned into an amount of Myth Dust based on their rarity:
- Common: 1
- Uncommon: 3
- Rare: 5

During a long rest, a creature may pray to Aroshi to turn their Myth Dust into specific cards at a cost based on the new card's rarity:
- Common: 3
- Uncommon: 9
- Rare: 15

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

### Point value cheatsheet

- Stagger/Debil/exh/confused/double disadv condi: 5
- Single disadv condi/sluggish/silence/charm: 3
- Prone/taunt/similar/root: 2
- Slow: 1
- Condi lasts until save/longer: x2
- Limited feature conditional: -1
- General conditional: -2
- Party positioning conditional: -1
- Common location conditional: -1
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
- Armour proficiency requirement: -2
- Weapon proficiency requirement: -3
- 1/LR spellcast: 1/lvl
- Learn spell: 1/lvl
- Unlimited spellcast: 5/lvl

## Condition definitions

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

A dominated creature is charmed. Additionally, the creature that dominated it controls the dominated creature during the dominated creature's turn. It can take move actions as usual and can use its standard action to take any Attack action (or Multiattack, where relevant).

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

- Investigate replacing body armours with non-party class feature-esque uniques? Or maybe subclasses? Replace base types for different armour weights to try to match archetypes nicely.

- Replace rings: 2x slots, start as very basic single stat, single point values, put gems into them to add thematic mods (eg. add wep dmg gem to give +2 wep dmg while bloodied). Using same gem either upgrades existing of that type or rerolls it

- Replace hard doubled mods with min 3pts

- Ban homebrew classes? Still allow homebrew subclasses? Harder for a subclass to spiral out of defined balance constraints. Loads of subclass + class options nowadays

- How do people feel about the themepark-ness structure of our campaigns? As opposed to more sandboxy concepts where the players define their own goals and ambitions (Downside is potentially campaign length). Subquestion if happy, how about streamlining exploration elements (eg. food/water/resources, searching for loot, etc.) (If wanting to bring them back, I'd want them to be more involved, and not just extra words for everyone to say)

- How do people feel about current character customisation through loot? Currently lots of different elements involved, and periods of loot rolls + town stuff often takes over an hour due to many decision points

- Mage Slayer discussion - rework into "ranged slayer" (give new name), defensive changes to targeting you from 10+' away (except melee + touch?), remove conc, off changes to when creature within certain distance gets similar targeted you get accuracy bonus vs the offender until end of next turn OR shift up to sp toward offender
- Change Specialist - to accuracy (next campaign is fine)