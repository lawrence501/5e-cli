import random

inspirationList = ["Mewgenics", "Kingdom Hearts", "Seekers", "Draw Steel", "Aethermancer", "Legends of Runeterra", "Slormancer", "Soulstone Survivors", "Supernatural", "Magicraft", "4e", "PoE2", "Good Wife", "Slice & Dice", "Brotato", "Death Must Die", "Bing", "BG3", "HSR", "Dimension 20", "Underworlds", "Solasta", "Magnus Archives", "D4", "Torchlight Infinite", "LoL", "DotA 2", "Sorcery", "MCU", "Diablo Immortal", "FFXIV", "Lost Ark", "Fairytales", "John Cleaver", "LevelUp", "No Sleep", "New World", "Hades", "Tem Tem", "Vampire Princess", "Doctor Who", "TWD", "Hannibal", "Buffyverse", "Cypher", "ToME", "D3", "PoE", "GW2", "DD", "D2", "UA", "GD", "Talisman",
                   "Elder Scrolls", "Neverwinter", "Median", "Divinity OS 2", "D&D", "D&D", "Zelda", "Dark Souls", "Binding of Isaac", "Pillars of Eternity", "Nier", "Fallout", "Monsters' Den", "Hand of Fate", "Fire Emblem Heroes", "Torchlight 2", "Tyranny", "Slay the Spire", "Pokemon", "AdventureQuest", "WoW", "DragonFable", "Bloodborne", "Epic7", "Last Epoch", "Pathfinder", "Age of Sigmar"]

if __name__ == "__main__":
    while True:
        dummy = input("Press enter to generate inspiration: ")
        print(random.choice(inspirationList))
