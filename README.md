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

### Manual value cheatsheet

Always based on dmg where possible

- Default/more: 30%
- Slight: 15%
- Substantial: 60%
- Massive: 120%
- Barely: 7%

### Point value cheatsheet

- Dominate: 8
- Stagger/Debil/exh/confused/uncon/double disadv condi: 5
- Single disadv condi/sluggish/silence/charm/root: 3
- Prone/taunt/similar/slow: 2
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

- Investigate replacing body armours with non-party class feature-esque uniques? Or maybe subclasses? Replace base types for different armour weights to try to match archetypes nicely.

- Replace rings: Soul gems, new slot is your soul, absorb gems to add thematic mods (eg. add wep dmg gem to give +2 wep dmg while bloodied). Using same gem either upgrades existing mod of that type or rerolls it (player's choice). new tags:
    - choice of ability score; temp hp; choice of type res; speed; choice of weapon class; 1h; 2h; shield; wep dmg; spell dmg; choice of dmg type; phys dmg; non-phys dmg; minion off; phys res; non-phys res; dot; debuff; buff; crit; minion def; aoe; throw; choice of range type; pers area; conc; hit die healing; max hit dice; hp; choice of save; ac; throw

- Remove amulets

- New loot option: attunement-requiring magic items. Any that fit into an equipment slot already in use are reflavoured to be on another slot. Can only attune to 1 at a time. Use existing wondrous weightings of rarities, but pushed 1 up (making the 2% Artifacts). Obviously, this means they will be pretty powerful, and must therefore be high on the loot table

- Replace hard doubled mods with min 3pts (or just use D4 system? min 3 -> min 2 -> min 1 for 1 affix, 2, and 3 respectively)

- Change books such that they are modifiers on generic types of powers, eg. your sphere AoE powers are now lines with length of their diameter

- Change crafting stones to be instant-use, but allow rollback (Shrines). Put them instead of the gold-replacements for commons on the wondrous/attunement rolls. Those slots give 1 shrine, actual crafting table slot gives 2.

- Ban homebrew classes? Still allow homebrew subclasses? Harder for a subclass to spiral out of defined balance constraints. Loads of subclass + class options nowadays

- How do people feel about the themepark-ness structure of our campaigns? As opposed to more sandboxy concepts where the players define their own goals and ambitions (Downside is potentially campaign length). Subquestion if happy, how about streamlining exploration elements (eg. food/water/resources, searching for loot, etc.) (If wanting to bring them back, I'd want them to be more involved, and not just extra words for everyone to say)

- Mage Slayer discussion - rework into "ranged slayer" (give new name), defensive changes to targeting you from 10+' away (except melee + touch?), remove conc, off changes to when creature within certain distance gets similar targeted you get accuracy bonus vs the offender until end of next turn OR shift up to sp toward offender
- Change Specialist - to accuracy (next campaign is fine)
- Future changes: all armour mastery feats lose prof, gain half feat (DEX/CON, DEX/CON, STR/CON respectively); medium adds negative DEX doesn't reduce AC and changes CON thing to +2 half dmg pseudo-AC; heavy reduce to 2 and change to all atks; add to light 1st 2 pts of mobile, remove mobile; remove half feat from inspiring leader

- Discuss party wipes --> single player death. Heroic sacrifice actions?

- Discuss no side-gaming

- Discuss Adrian's golden home with its perfectly balanced distances from all + we can get our dice blessed by Muffin

- New card set: sidekicks

