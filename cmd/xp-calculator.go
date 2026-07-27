package main

import (
	"fmt"
)

func getCharacterStatus() int {
	for {
		var characterStatus int
		// Get character status from user input
		fmt.Println("WHAT IS YOUR CHARACTER STATUS?")
		fmt.Println("1 - Adventurer/Champion (no true reincarnation)")
		fmt.Println("2 - Hero (one true reincarnation)")
		fmt.Println("3 - Legend (two or more true reincarnations)")
		_, err := fmt.Scanln(&characterStatus)
		if err != nil {
			fmt.Println("--------------------------------------")
			fmt.Println("Invalid input. Please enter a whole number, for example: 1")
			fmt.Println("--------------------------------------")
			fmt.Println()
			var discard string
			fmt.Scanln(&discard)
			continue
		}
		if characterStatus == 1 || characterStatus == 2 || characterStatus == 3 {
			fmt.Println("Valid choice:", characterStatus)
			fmt.Println()
			return characterStatus
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

func main() {
	fmt.Println("Welcome to the XP Calculator!")
	fmt.Println()

	characterStatus := getCharacterStatus()
	_ = characterStatus

	experiencePoints := getExperiencePoints()
	_ = experiencePoints
}
