import random

inspirationList = ["Hades", "Tem Tem", "Vampire Princess", "Monster Train", "Doctor Who", "Witcher", "TWD", "Crumbling Keep", "Hannibal", "Buffy", "Angel", "Numenera", "ToME", "SSB", "MTG", "D3", "PoE", "GW2", "DD", "D2", "UA", "GD", "Talisman",
                   "Armello", "ES", "NW", "Median", "DOS", "D&D", "D&D", "LoZ", "DS", "BoI", "Pillars", "NA", "FO", "MD", "HoF", "EH", "FEH", "TL2", "Tyranny", "STS", "Pokemon", "AQ", "WoW", "DF", "BB", "E7", "LE", "PF", "WH"]

if __name__ == "__main__":
    while True:
        dummy = input("Press enter to generate inspiration: ")
        print(random.choice(inspirationList))
