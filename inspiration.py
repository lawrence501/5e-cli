import random

inspirationList = ["Sorcery", "Marvel", "Diablo Immortal", "FFXIV", "Lost Ark", "Fairytales", "John Cleaver", "LevelUp", "No Sleep", "New World", "Hades", "Tem Tem", "Vampire Princess", "Doctor Who", "Witcher", "TWD", "Hannibal", "Buffy", "Angel", "Numenera", "ToME", "MTG", "D3", "PoE", "GW2", "DD", "D2", "UA", "GD", "Talisman",
                   "ES", "Neverwinter", "Median", "DOS", "D&D", "D&D", "LoZ", "DS", "BoI", "Pillars", "Nier", "FO", "MD", "HoF", "EH", "FEH", "TL2", "Tyranny", "STS", "Pokemon", "AQ", "WoW", "DF", "BB", "E7", "LE", "PF", "WH"]

if __name__ == "__main__":
    while True:
        dummy = input("Press enter to generate inspiration: ")
        print(random.choice(inspirationList))
