package main

import (
	"fmt"
)

type LevelData struct {
	Level int
	Ranks []int
}

var heroicLevelsLowTR = []LevelData{
	{Level: 1, Ranks: []int{0, 800, 1600, 2400, 3200}},
	{Level: 2, Ranks: []int{4000, 6400, 8800, 11200, 13600}},
	{Level: 3, Ranks: []int{16000, 20800, 25600, 30400, 35200}},
	{Level: 4, Ranks: []int{40000, 46400, 52800, 59200, 65600}},
	{Level: 5, Ranks: []int{72000, 80000, 88000, 96000, 104000}},
	{Level: 6, Ranks: []int{112000, 121600, 131200, 140800, 150400}},
	{Level: 7, Ranks: []int{160000, 173000, 186000, 199000, 212000}},
	{Level: 8, Ranks: []int{225000, 241000, 257000, 273000, 289000}},
	{Level: 9, Ranks: []int{305000, 324000, 343000, 362000, 381000}},
	{Level: 10, Ranks: []int{400000, 422000, 444000, 466000, 488000}},
	{Level: 11, Ranks: []int{510000, 534000, 558000, 582000, 606000}},
	{Level: 12, Ranks: []int{630000, 656000, 682000, 708000, 734000}},
	{Level: 13, Ranks: []int{760000, 788000, 816000, 844000, 872000}},
	{Level: 14, Ranks: []int{900000, 930000, 960000, 990000, 1020000}},
	{Level: 15, Ranks: []int{1050000, 1082000, 1114000, 1146000, 1178000}},
	{Level: 16, Ranks: []int{1210000, 1243000, 1276000, 1309000, 1342000}},
	{Level: 17, Ranks: []int{1375000, 1409000, 1443000, 1477000, 1511000}},
	{Level: 18, Ranks: []int{1545000, 1580000, 1615000, 1650000, 1685000}},
	{Level: 19, Ranks: []int{1720000, 1756000, 1792000, 1828000, 1864000}},
	{Level: 20, Ranks: []int{1900000, 0, 0, 0, 0}},
}

var heroicLevelsMidTR = []LevelData{
	{Level: 1, Ranks: []int{0, 1200, 2400, 3600, 4800}},
	{Level: 2, Ranks: []int{6000, 9600, 13200, 16800, 20400}},
	{Level: 3, Ranks: []int{24000, 31200, 38400, 45600, 52800}},
	{Level: 4, Ranks: []int{60000, 69600, 79200, 88800, 98400}},
	{Level: 5, Ranks: []int{108000, 120000, 132000, 144000, 156000}},
	{Level: 6, Ranks: []int{168000, 182400, 196800, 211200, 225600}},
	{Level: 7, Ranks: []int{240000, 259500, 279000, 298500, 318000}},
	{Level: 8, Ranks: []int{337500, 361500, 385500, 409500, 433500}},
	{Level: 9, Ranks: []int{457500, 486000, 514500, 543000, 571500}},
	{Level: 10, Ranks: []int{600000, 633000, 666000, 699000, 732000}},
	{Level: 11, Ranks: []int{765000, 801000, 837000, 873000, 909000}},
	{Level: 12, Ranks: []int{945000, 984000, 1023000, 1062000, 1101000}},
	{Level: 13, Ranks: []int{1140000, 1182000, 1224000, 1266000, 1308000}},
	{Level: 14, Ranks: []int{1350000, 1395000, 1440000, 1485000, 1530000}},
	{Level: 15, Ranks: []int{1575000, 1623000, 1671000, 1719000, 1767000}},
	{Level: 16, Ranks: []int{1815000, 1864500, 1914000, 1963500, 2013000}},
	{Level: 17, Ranks: []int{2062500, 2113500, 2164500, 2215500, 2266500}},
	{Level: 18, Ranks: []int{2317500, 2370000, 2422500, 2475000, 2527500}},
	{Level: 19, Ranks: []int{2580000, 2634000, 2688000, 2742000, 2796000}},
	{Level: 20, Ranks: []int{2850000, 0, 0, 0, 0}}, // heroic XP cap
}

var heroicLevelsHighTR = []LevelData{
	{Level: 1, Ranks: []int{0, 1600, 3200, 4800, 6400}},
	{Level: 2, Ranks: []int{8000, 12800, 17600, 22400, 27200}},
	{Level: 3, Ranks: []int{32000, 41600, 51200, 60800, 70400}},
	{Level: 4, Ranks: []int{80000, 92800, 105600, 118400, 131200}},
	{Level: 5, Ranks: []int{144000, 160000, 176000, 192000, 208000}},
	{Level: 6, Ranks: []int{224000, 243200, 262400, 281600, 300800}},
	{Level: 7, Ranks: []int{320000, 346000, 372000, 398000, 424000}},
	{Level: 8, Ranks: []int{450000, 482000, 514000, 546000, 578000}},
	{Level: 9, Ranks: []int{610000, 648000, 686000, 724000, 762000}},
	{Level: 10, Ranks: []int{800000, 844000, 888000, 932000, 976000}},
	{Level: 11, Ranks: []int{1020000, 1068000, 1116000, 1164000, 1212000}},
	{Level: 12, Ranks: []int{1260000, 1312000, 1364000, 1416000, 1468000}},
	{Level: 13, Ranks: []int{1520000, 1576000, 1632000, 1688000, 1744000}},
	{Level: 14, Ranks: []int{1800000, 1860000, 1920000, 1980000, 2040000}},
	{Level: 15, Ranks: []int{2100000, 2164000, 2228000, 2292000, 2356000}},
	{Level: 16, Ranks: []int{2420000, 2486000, 2552000, 2618000, 2684000}},
	{Level: 17, Ranks: []int{2750000, 2818000, 2886000, 2954000, 3022000}},
	{Level: 18, Ranks: []int{3090000, 3160000, 3230000, 3300000, 3370000}},
	{Level: 19, Ranks: []int{3440000, 3512000, 3584000, 3656000, 3728000}},
	{Level: 20, Ranks: []int{3800000, 0, 0, 0, 0}}, // heroic XP cap
}

var epicLevels = []LevelData{
	{Level: 20, Ranks: []int{0, 120000, 240000, 360000, 480000}},
	{Level: 21, Ranks: []int{600000, 730000, 860000, 990000, 1120000}},
	{Level: 22, Ranks: []int{1250000, 1390000, 1530000, 1670000, 1810000}},
	{Level: 23, Ranks: []int{1950000, 2100000, 2250000, 2400000, 2550000}},
	{Level: 24, Ranks: []int{2700000, 2860000, 3020000, 3180000, 3340000}},
	{Level: 25, Ranks: []int{3500000, 3670000, 3840000, 4010000, 4180000}},
	{Level: 26, Ranks: []int{4350000, 4530000, 4710000, 4890000, 5070000}},
	{Level: 27, Ranks: []int{5250000, 5440000, 5630000, 5820000, 6010000}},
	{Level: 28, Ranks: []int{6200000, 6400000, 6600000, 6800000, 7000000}},
	{Level: 29, Ranks: []int{7200000, 7410000, 7620000, 7830000, 8040000}},
	{Level: 30, Ranks: []int{8250000, 0, 0, 0, 0}}, // epic XP cap
}

var legendaryLevels = []LevelData{
	{Level: 30, Ranks: []int{0, 320000, 640000, 960000, 1280000}},
	{Level: 31, Ranks: []int{1600000, 2000000, 2400000, 2800000, 3200000}},
	{Level: 32, Ranks: []int{3600000, 4020000, 4440000, 4860000, 5280000}},
	{Level: 33, Ranks: []int{5700000, 6140000, 6580000, 7020000, 7460000}},
	{Level: 34, Ranks: []int{7900000, 8360000, 8820000, 9280000, 9740000}},
	{Level: 35, Ranks: []int{10200000, 10680000, 11160000, 11640000, 12120000}},
	{Level: 36, Ranks: []int{12600000, 0, 0, 0, 0}}, // legendary XP cap
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

func getLevelFromExperiencePoints(levels []LevelData, experiencePoints int) (int, int, int) {
	var characterLevel int
	var characterRank int
	var experiencesToNextLevel int

	// Determine character level based on experience points
	for i := len(levels) - 1; i >= 0; i-- {
		if levels[i].Ranks[0] <= experiencePoints {
			characterLevel = levels[i].Level
			// Determine experiences needed for next level
			if i < len(levels)-1 {
				experiencesToNextLevel = levels[i+1].Ranks[0] - experiencePoints
			} else {
				experiencesToNextLevel = 0 // Max level reached, no more XP needed for next level
				characterRank = 0
				break
			}
			// Determine rank based on experience points
			for j := 4; j >= 0; j-- {
				if levels[i].Ranks[j] <= experiencePoints {
					characterRank = j
					break
				}
			}
			break
		}
	}

	return characterLevel, characterRank, experiencesToNextLevel
}

func printCharacterStatus(level int, rank int, experiencesToNextLevel int) {
	fmt.Println("Your character level is:", level)
	fmt.Println("Your character rank is:", rank)
	if experiencesToNextLevel > 0 {
		fmt.Println("Experience points needed for next level:", experiencesToNextLevel)
	} else {
		fmt.Println("You have reached the maximum level in heroic, epic or legendary tier. No more experience points are needed.")
	}
}

func printResult(levelData []LevelData, levelTypeText string, experiencePoints int) {
	fmt.Println("You have chosen", levelTypeText, " and", experiencePoints, "experience points.")
	characterLevel, characterRank, experiencesToNextLevel := getLevelFromExperiencePoints(levelData, experiencePoints)
	printCharacterStatus(characterLevel, characterRank, experiencesToNextLevel)
}

func main() {
	fmt.Println("Welcome to the XP Calculator!")
	fmt.Println()

	// Get level type from user input (1 = heroic, 2 = epic, 3 = legendary)
	levelType := getTypeofLevel()

	// Get reincarnations status from user input (only for heroic levels)
	var reincarnationsStatus int
	if levelType == 1 {
		reincarnationsStatus = getReincarnationsStatus()
	}

	// Get experience points from user input
	experiencePoints := getExperiencePoints()

	// Determine character level, rank, and experiences needed for next level based on user input
	switch levelType {
	case 1:
		switch reincarnationsStatus {
		case 1:
			printResult(heroicLevelsLowTR, "Heroic Level with 0-3 true reincarnations", experiencePoints)
		case 2:
			printResult(heroicLevelsMidTR, "Heroic Level with 4-6 true reincarnations", experiencePoints)
		case 3:
			printResult(heroicLevelsHighTR, "Heroic Level with 7 or more true reincarnations", experiencePoints)
		default:
			fmt.Println("Invalid reincarnations status")
		}
	case 2:
		printResult(epicLevels, "Epic Level", experiencePoints)
	case 3:
		printResult(legendaryLevels, "Legendary Level", experiencePoints)
	default:
		fmt.Println("Invalid level type")
	}
}
