package main

import (
	"fmt"
	"math/rand/v2"
)

var originalClasses = []string{
	"Fighter",
	"Wizard",
	"Rogue",
	"Cleric",
	"Paladin",
	"Ranger",
	"Bard",
	"Monk",
	"Druid",
	"Barbarian",
	"Sorcerer",
	"Warlock",
	"Artificer",
	"Alchemist",
	"Favored Soul",
}

var archetypeClasses = []string{
	"Stormsinger",
	"Dark Apostate",
	"Sacred Fist",
	"Blightcaster",
	"Dark Hunter",
	"Acolyte of the Skin",
	"Dragon Lord",
	"Dragon Disciple",
	"Arcane Trickster",
	"Wild Mage",
}

var races = []string{
	"Human",
	"Elf",
	"Dwarf",
	"Halfling",
	"Gnome",
	"Half-Orc",
	"Half-Elf",
	"Tiefling",
	"Dragonborn",
	"Aasimar",
	"Shifter",
	"Wood Elf",
	"Drow",
	"Dhampir",
	"Duergar",
	"Eladrin",
	"Tabaxi",
	"Warforged",
}

func randomRace() string {
	return races[rand.IntN(len(races))]
}

func randomOriginalClass() string {
	return originalClasses[rand.IntN(len(originalClasses))]
}

func randomArchetypeClass() string {
	return archetypeClasses[rand.IntN(len(archetypeClasses))]
}

func containsClass(classes []string, wanted string) bool {
	for _, class := range classes {
		if class == wanted {
			return true
		}
	}

	return false
}

func getLevels() int {
	return rand.IntN(20) + 1
}

func generateClasses(numClasses int) (string, []string) {
	var generatedRace string
	var generatedClasses []string

	generatedRace = randomRace()

	for i := 0; i < numClasses; i++ {
		values := []int{1, 2}
		// Randomly choose between an original class and an archetype class
		choice := values[rand.IntN(len(values))]

		if choice == 1 {
			for {
				randomClass := randomOriginalClass()
				// Ensure the randomly selected original class is not already in the generated classes list
				if !containsClass(generatedClasses, randomClass) {
					generatedClasses = append(generatedClasses, randomClass)
					break
				}
			}
		} else if choice == 2 {
			for {
				randomClass := randomArchetypeClass()
				// Ensure the randomly selected archetype class is not already in the generated classes list
				if !containsClass(generatedClasses, randomClass) {
					generatedClasses = append(generatedClasses, randomClass)
					break
				}
			}
		}
	}
	return generatedRace, generatedClasses
}

func main() {
	for {
		fmt.Println("Welcome to the Random Build Generator!")
		fmt.Println()
		fmt.Println("Do you want to go pure or multiclass? Select how many classes you want to use:")
		fmt.Println("1. Pure")
		fmt.Println("2. Two classes")
		fmt.Println("3. Three classes")
		fmt.Println("4. Randomize how many classes")
		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("--------------------------------------")
			fmt.Println("Invalid input. Please enter a whole number, for example: 1")
			fmt.Println("--------------------------------------")
			fmt.Println()
			var discard string
			fmt.Scanln(&discard)
			continue
		}
		if choice == 1 || choice == 2 || choice == 3 || choice == 4 {
			fmt.Println("Valid choice:", choice)
			fmt.Println()
			switch choice {
			case 1:
				generatedRace, generatedClass := generateClasses(1)
				fmt.Println("GENERATED BUILD:")
				fmt.Println("Race:", generatedRace)
				fmt.Println("Class:", generatedClass[0], "20")
			case 2:
				// Implement logic for two classes
				generatedRace, generatedClasses := generateClasses(2)
				fmt.Println("GENERATED BUILD:")
				fmt.Println("Race:", generatedRace)
				fmt.Println("Classes:", generatedClasses[0], "and", generatedClasses[1])
			case 3:
				// Implement logic for three classes
				generatedRace, generatedClasses := generateClasses(3)
				fmt.Println("GENERATED BUILD:")
				fmt.Println("Race:", generatedRace)
				fmt.Println("Classes:", generatedClasses)
			case 4:
				// Implement logic for random number of classes
			}
			break
		} else {
			fmt.Println("--------------------------------------")
			fmt.Println("Invalid choice. Use only 1, 2, 3, or 4.")
			fmt.Println("--------------------------------------")
			fmt.Println()
		}
	}

}
