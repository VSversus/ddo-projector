package main

// todo create unit tests for getCharacterStatus and getExperiencePoints functions

import (
	"fmt"
)

type LevelData struct {
	Level int
	Rank0 int
	Rank1 int
	Rank2 int
	Rank3 int
	Rank4 int
}

var heroicLevelsLowTR = []LevelData{
	{Level: 1, Rank0: 0, Rank1: 800, Rank2: 1600, Rank3: 2400, Rank4: 3200},
	{Level: 2, Rank0: 4000, Rank1: 6400, Rank2: 8800, Rank3: 11200, Rank4: 13600},
	{Level: 3, Rank0: 16000, Rank1: 20800, Rank2: 25600, Rank3: 30400, Rank4: 35200},
	{Level: 4, Rank0: 40000, Rank1: 46400, Rank2: 52800, Rank3: 59200, Rank4: 65600},
	{Level: 5, Rank0: 72000, Rank1: 80000, Rank2: 88000, Rank3: 96000, Rank4: 104000},
	{Level: 6, Rank0: 112000, Rank1: 121600, Rank2: 131200, Rank3: 140800, Rank4: 150400},
	{Level: 7, Rank0: 160000, Rank1: 173000, Rank2: 186000, Rank3: 199000, Rank4: 212000},
	{Level: 8, Rank0: 225000, Rank1: 241000, Rank2: 257000, Rank3: 273000, Rank4: 289000},
	{Level: 9, Rank0: 305000, Rank1: 324000, Rank2: 343000, Rank3: 362000, Rank4: 381000},
	{Level: 10, Rank0: 400000, Rank1: 422000, Rank2: 444000, Rank3: 466000, Rank4: 488000},
	{Level: 11, Rank0: 510000, Rank1: 534000, Rank2: 558000, Rank3: 582000, Rank4: 606000},
	{Level: 12, Rank0: 630000, Rank1: 656000, Rank2: 682000, Rank3: 708000, Rank4: 734000},
	{Level: 13, Rank0: 760000, Rank1: 788000, Rank2: 816000, Rank3: 844000, Rank4: 872000},
	{Level: 14, Rank0: 900000, Rank1: 930000, Rank2: 960000, Rank3: 990000, Rank4: 1020000},
	{Level: 15, Rank0: 1050000, Rank1: 1082000, Rank2: 1114000, Rank3: 1146000, Rank4: 1178000},
	{Level: 16, Rank0: 1210000, Rank1: 1243000, Rank2: 1276000, Rank3: 1309000, Rank4: 1342000},
	{Level: 17, Rank0: 1375000, Rank1: 1409000, Rank2: 1443000, Rank3: 1477000, Rank4: 1511000},
	{Level: 18, Rank0: 1545000, Rank1: 1580000, Rank2: 1615000, Rank3: 1650000, Rank4: 1685000},
	{Level: 19, Rank0: 1720000, Rank1: 1756000, Rank2: 1792000, Rank3: 1828000, Rank4: 1864000},
	{Level: 20, Rank0: 1900000, Rank1: 0, Rank2: 0, Rank3: 0, Rank4: 0},
}

var heroicLevelsMidTR = []LevelData{
	{Level: 1, Rank0: 0, Rank1: 1200, Rank2: 2400, Rank3: 3600, Rank4: 4800},
	{Level: 2, Rank0: 6000, Rank1: 9600, Rank2: 13200, Rank3: 16800, Rank4: 20400},
	{Level: 3, Rank0: 24000, Rank1: 31200, Rank2: 38400, Rank3: 45600, Rank4: 52800},
	{Level: 4, Rank0: 60000, Rank1: 69600, Rank2: 79200, Rank3: 88800, Rank4: 98400},
	{Level: 5, Rank0: 108000, Rank1: 120000, Rank2: 132000, Rank3: 144000, Rank4: 156000},
	{Level: 6, Rank0: 168000, Rank1: 182400, Rank2: 196800, Rank3: 211200, Rank4: 225600},
	{Level: 7, Rank0: 240000, Rank1: 259500, Rank2: 279000, Rank3: 298500, Rank4: 318000},
	{Level: 8, Rank0: 337500, Rank1: 361500, Rank2: 385500, Rank3: 409500, Rank4: 433500},
	{Level: 9, Rank0: 457500, Rank1: 486000, Rank2: 514500, Rank3: 543000, Rank4: 571500},
	{Level: 10, Rank0: 600000, Rank1: 633000, Rank2: 666000, Rank3: 699000, Rank4: 732000},
	{Level: 11, Rank0: 765000, Rank1: 801000, Rank2: 837000, Rank3: 873000, Rank4: 909000},
	{Level: 12, Rank0: 945000, Rank1: 984000, Rank2: 1023000, Rank3: 1062000, Rank4: 1101000},
	{Level: 13, Rank0: 1140000, Rank1: 1182000, Rank2: 1224000, Rank3: 1266000, Rank4: 1308000},
	{Level: 14, Rank0: 1350000, Rank1: 1395000, Rank2: 1440000, Rank3: 1485000, Rank4: 1530000},
	{Level: 15, Rank0: 1575000, Rank1: 1623000, Rank2: 1671000, Rank3: 1719000, Rank4: 1767000},
	{Level: 16, Rank0: 1815000, Rank1: 1864500, Rank2: 1914000, Rank3: 1963500, Rank4: 2013000},
	{Level: 17, Rank0: 2062500, Rank1: 2113500, Rank2: 2164500, Rank3: 2215500, Rank4: 2266500},
	{Level: 18, Rank0: 2317500, Rank1: 2370000, Rank2: 2422500, Rank3: 2475000, Rank4: 2527500},
	{Level: 19, Rank0: 2580000, Rank1: 2634000, Rank2: 2688000, Rank3: 2742000, Rank4: 2796000},
	{Level: 20, Rank0: 2850000, Rank1: 0, Rank2: 0, Rank3: 0, Rank4: 0}, // heroic XP cap
}

var heroicLevelsHighTR = []LevelData{
	{Level: 1, Rank0: 0, Rank1: 1600, Rank2: 3200, Rank3: 4800, Rank4: 6400},
	{Level: 2, Rank0: 8000, Rank1: 12800, Rank2: 17600, Rank3: 22400, Rank4: 27200},
	{Level: 3, Rank0: 32000, Rank1: 41600, Rank2: 51200, Rank3: 60800, Rank4: 70400},
	{Level: 4, Rank0: 80000, Rank1: 92800, Rank2: 105600, Rank3: 118400, Rank4: 131200},
	{Level: 5, Rank0: 144000, Rank1: 160000, Rank2: 176000, Rank3: 192000, Rank4: 208000},
	{Level: 6, Rank0: 224000, Rank1: 243200, Rank2: 262400, Rank3: 281600, Rank4: 300800},
	{Level: 7, Rank0: 320000, Rank1: 346000, Rank2: 372000, Rank3: 398000, Rank4: 424000},
	{Level: 8, Rank0: 450000, Rank1: 482000, Rank2: 514000, Rank3: 546000, Rank4: 578000},
	{Level: 9, Rank0: 610000, Rank1: 648000, Rank2: 686000, Rank3: 724000, Rank4: 762000},
	{Level: 10, Rank0: 800000, Rank1: 844000, Rank2: 888000, Rank3: 932000, Rank4: 976000},
	{Level: 11, Rank0: 1020000, Rank1: 1068000, Rank2: 1116000, Rank3: 1164000, Rank4: 1212000},
	{Level: 12, Rank0: 1260000, Rank1: 1312000, Rank2: 1364000, Rank3: 1416000, Rank4: 1468000},
	{Level: 13, Rank0: 1520000, Rank1: 1576000, Rank2: 1632000, Rank3: 1688000, Rank4: 1744000},
	{Level: 14, Rank0: 1800000, Rank1: 1860000, Rank2: 1920000, Rank3: 1980000, Rank4: 2040000},
	{Level: 15, Rank0: 2100000, Rank1: 2164000, Rank2: 2228000, Rank3: 2292000, Rank4: 2356000},
	{Level: 16, Rank0: 2420000, Rank1: 2486000, Rank2: 2552000, Rank3: 2618000, Rank4: 2684000},
	{Level: 17, Rank0: 2750000, Rank1: 2818000, Rank2: 2886000, Rank3: 2954000, Rank4: 3022000},
	{Level: 18, Rank0: 3090000, Rank1: 3160000, Rank2: 3230000, Rank3: 3300000, Rank4: 3370000},
	{Level: 19, Rank0: 3440000, Rank1: 3512000, Rank2: 3584000, Rank3: 3656000, Rank4: 3728000},
	{Level: 20, Rank0: 3800000, Rank1: 0, Rank2: 0, Rank3: 0, Rank4: 0}, // heroic XP cap
}

var epicLevels = []LevelData{
	{Level: 20, Rank0: 0, Rank1: 120000, Rank2: 240000, Rank3: 360000, Rank4: 480000},
	{Level: 21, Rank0: 600000, Rank1: 730000, Rank2: 860000, Rank3: 990000, Rank4: 1120000},
	{Level: 22, Rank0: 1250000, Rank1: 1390000, Rank2: 1530000, Rank3: 1670000, Rank4: 1810000},
	{Level: 23, Rank0: 1950000, Rank1: 2100000, Rank2: 2250000, Rank3: 2400000, Rank4: 2550000},
	{Level: 24, Rank0: 2700000, Rank1: 2860000, Rank2: 3020000, Rank3: 3180000, Rank4: 3340000},
	{Level: 25, Rank0: 3500000, Rank1: 3670000, Rank2: 3840000, Rank3: 4010000, Rank4: 4180000},
	{Level: 26, Rank0: 4350000, Rank1: 4530000, Rank2: 4710000, Rank3: 4890000, Rank4: 5070000},
	{Level: 27, Rank0: 5250000, Rank1: 5440000, Rank2: 5630000, Rank3: 5820000, Rank4: 6010000},
	{Level: 28, Rank0: 6200000, Rank1: 6400000, Rank2: 6600000, Rank3: 6800000, Rank4: 7000000},
	{Level: 29, Rank0: 7200000, Rank1: 7410000, Rank2: 7620000, Rank3: 7830000, Rank4: 8040000},
	{Level: 30, Rank0: 8250000, Rank1: 0, Rank2: 0, Rank3: 0, Rank4: 0}, // epic XP cap
}

var legendaryLevels = []LevelData{
	{Level: 30, Rank0: 0, Rank1: 320000, Rank2: 640000, Rank3: 960000, Rank4: 1280000},
	{Level: 31, Rank0: 1600000, Rank1: 2000000, Rank2: 2400000, Rank3: 2800000, Rank4: 3200000},
	{Level: 32, Rank0: 3600000, Rank1: 4020000, Rank2: 4440000, Rank3: 4860000, Rank4: 5280000},
	{Level: 33, Rank0: 5700000, Rank1: 6140000, Rank2: 6580000, Rank3: 7020000, Rank4: 7460000},
	{Level: 34, Rank0: 7900000, Rank1: 8360000, Rank2: 8820000, Rank3: 9280000, Rank4: 9740000},
	{Level: 35, Rank0: 10200000, Rank1: 10680000, Rank2: 11160000, Rank3: 11640000, Rank4: 12120000},
	{Level: 36, Rank0: 12600000, Rank1: 0, Rank2: 0, Rank3: 0, Rank4: 0}, // legendary XP cap
}

func getTypeofLevel() int {
	for {
		var levelType int
		// Get level type from user input
		fmt.Println("WHAT TYPE OF LEVEL IS YOUR CHARACTER?")
		fmt.Println("Only heroic levels are affected by number of true reincarnations.")
		fmt.Println("Please choose one of the following options:")
		fmt.Println("1 - Heroic Level (1-20)")
		fmt.Println("2 - Epic Level (20-30)")
		fmt.Println("3 - Legendary Level (30-36)")
		_, err := fmt.Scanln(&levelType)
		if err != nil {
			fmt.Println("--------------------------------------")
			fmt.Println("Invalid input. Please enter a whole number, for example: 1")
			fmt.Println("--------------------------------------")
			fmt.Println()
			var discard string
			fmt.Scanln(&discard)
			continue
		}
		if levelType == 1 || levelType == 2 || levelType == 3 {
			fmt.Println("Valid choice:", levelType)
			fmt.Println()
			return levelType
		}

		fmt.Println("--------------------------------------")
		fmt.Println("Invalid choice. Use only 1, 2, or 3.")
		fmt.Println("--------------------------------------")
		fmt.Println()
	}
}

func getReincarnationsStatus() int {
	for {
		var reincarnationsStatus int
		// Get character status from user input
		fmt.Println("HOW MANY TRUE REINCARNATIONS DOES YOUR CHARACTER HAVE?")
		fmt.Println("This affects the XP calculation. Please choose one of the following options:")
		fmt.Println("1 - 0-3 true reincarnations")
		fmt.Println("2 - 4-6 true reincarnations")
		fmt.Println("3 - 7 or more true reincarnations")
		_, err := fmt.Scanln(&reincarnationsStatus)
		if err != nil {
			fmt.Println("--------------------------------------")
			fmt.Println("Invalid input. Please enter a whole number, for example: 1")
			fmt.Println("--------------------------------------")
			fmt.Println()
			var discard string
			fmt.Scanln(&discard)
			continue
		}
		if reincarnationsStatus == 1 || reincarnationsStatus == 2 || reincarnationsStatus == 3 {
			fmt.Println("Valid choice:", reincarnationsStatus)
			fmt.Println()
			return reincarnationsStatus
		}

		fmt.Println("--------------------------------------")
		fmt.Println("Invalid choice. Use only 1, 2, or 3.")
		fmt.Println("--------------------------------------")
		fmt.Println()
	}
}

func getExperiencePoints() int {
	for {
		var experiencePoints int
		// Get experience points from user input
		fmt.Println("ENTER YOUR EXPERIENCE POINTS:")
		_, err := fmt.Scanln(&experiencePoints)
		if err != nil {
			fmt.Println("--------------------------------------")
			fmt.Println("Invalid input. Please enter a whole number, for example: 123456")
			fmt.Println("--------------------------------------")
			fmt.Println()
			var discard string
			fmt.Scanln(&discard)
			continue
		} else if experiencePoints < 0 {
			fmt.Println("--------------------------------------")
			fmt.Println("Invalid input. Please enter a positive whole number, for example: 123456")
			fmt.Println("--------------------------------------")
			fmt.Println()
			continue
		}

		fmt.Println("Experience points accepted:", experiencePoints)
		fmt.Println()
		return experiencePoints
	}
}

func getLevelFromExperiencePoints(levelType, experiencePoints int) (int, int, int) {
	var characterLevel int
	_ = characterLevel
	var rank int
	_ = rank
	var experiencesToNextLevel int
	_ = experiencesToNextLevel

	switch levelType {
	// heroic low TR
	case 1:
		for i := len(heroicLevelsLowTR) - 1; i >= 0; i-- {
			if heroicLevelsLowTR[i].Rank0 <= experiencePoints {
				characterLevel = heroicLevelsLowTR[i].Level
				// todo implement rank and experiencesToNextLevel calculation
				break
			}
		}
	// heroic mid TR
	case 2:
	// heroic high TR
	case 3:
	// epic
	case 4:
	// legendary
	case 5:
	default:
		fmt.Println("Invalid level type")
	}
	return characterLevel, rank, experiencesToNextLevel
}

func main() {
	fmt.Println("Welcome to the XP Calculator!")
	fmt.Println()

	// Get level type from user input (1 = heroic, 2 = epic, 3 = legendary)
	levelType := getTypeofLevel()

	var reincarnationsStatus int
	if levelType == 1 {
		reincarnationsStatus = getReincarnationsStatus()
	}

	experiencePoints := getExperiencePoints()

	switch levelType {
	case 1:
		switch reincarnationsStatus {
		case 1:
			fmt.Println("You have chosen Heroic Level with 0-3 true reincarnations and", experiencePoints, "experience points.")
			characterLevel, _, _ := getLevelFromExperiencePoints(levelType, experiencePoints)
			fmt.Println("Your character level is:", characterLevel)
		case 2:
			fmt.Println("You have chosen Heroic Level with 4-6 true reincarnations and", experiencePoints, "experience points.")
		case 3:
			fmt.Println("You have chosen Heroic Level with 7 or more true reincarnations and", experiencePoints, "experience points.")
		default:
			fmt.Println("Invalid reincarnations status")
		}
	case 2:
		fmt.Println("You have chosen Epic Level with", experiencePoints, "experience points.")
	case 3:
		fmt.Println("You have chosen Legendary Level with", experiencePoints, "experience points.")
	default:
		fmt.Println("Invalid level type")
	}

	// todo implement the logic to calculate the level based on the inputs
}
