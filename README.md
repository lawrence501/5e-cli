# 5e-cli
Golang CLI helper for my D&amp;D 5e campaigns

## Loot Rules

### Relic generation

1. 1st base affix
2. 2nd base affix
3. 3rd base affix
4. Thematic
5. Name Thematic
6. Random from source
7. Random PoE
8. Random affix

### Point value cheatsheet

- Para: 10
- Stun: 8
- Stagger/Debil/double disadv condi: 5
- Single disadv condi/sluggish: 3
- Prone/similar: 2
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

### Dominated

Shares condition immunity with charm.

A dominated creature is charmed. Additionally, the creature that dominated it controls the dominated creature during the dominated creature's turn. It can take move actions as usual and can use its standard action to take any Attack action (or Multiattack, where relevant).

### Rattled

Shares condition immunity with debilitation.

A rattled creature cannot benefit from expertise dice and cannot take reactions.

### Sluggish

Shares condition immunity with exhaustion.

A sluggish creature reduces the number of attacks it can make with its Attack action by 1 (or Multiattack, where relevant).

### Staggered

Shares condition immunity with paralysis.

A staggered creature suffers the following effects:
- Incoming attack damage is always critical.
- All incoming attacks have advantage.
- All speeds are set to 0 feet.

## Terminology

### Splash Damage

Splash damage is always optional. It deals its damage, up to the amount of the original hit, to all creatures of your choice within 5' of the target.

## Revisited Combat Maneuvers

### Cavalier Stance

While riding your mount, your AC increases by 1.