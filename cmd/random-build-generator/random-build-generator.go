package main

import (
	"fmt"
	"math/rand/v2"
)

type Race struct {
	Name       string
	FreeToPlay bool
	Iconic     bool
}

type Class struct {
	Name       string
	FreeToPlay bool
	Archetype  string
}

var racesOptions = []Race{
	{Name: "Human", FreeToPlay: true, Iconic: false},
	{Name: "Elf", FreeToPlay: true, Iconic: false},
	{Name: "Dwarf", FreeToPlay: true, Iconic: false},
	{Name: "Halfling", FreeToPlay: true, Iconic: false},
	{Name: "Gnome", FreeToPlay: true, Iconic: false},
	{Name: "Half-Orc", FreeToPlay: true, Iconic: false},
	{Name: "Half-Elf", FreeToPlay: true, Iconic: false},
	{Name: "Tiefling", FreeToPlay: true, Iconic: false},
	{Name: "Dragonborn", FreeToPlay: true, Iconic: false},
	{Name: "Wood Elf", FreeToPlay: true, Iconic: false},
	{Name: "Drow", FreeToPlay: true, Iconic: false},
	{Name: "Aasimar", FreeToPlay: false, Iconic: false},
	{Name: "Shifter", FreeToPlay: false, Iconic: false},
	{Name: "Dhampir", FreeToPlay: false, Iconic: false},
	{Name: "Duergar", FreeToPlay: false, Iconic: false},
	{Name: "Eladrin", FreeToPlay: false, Iconic: false},
	{Name: "Tabaxi", FreeToPlay: false, Iconic: false},
	{Name: "Warforged", FreeToPlay: false, Iconic: false},
	{Name: "Purple Dragon Knight", FreeToPlay: false, Iconic: true},
	{Name: "Bladeforged", FreeToPlay: false, Iconic: true},
	{Name: "Deep Gnome", FreeToPlay: false, Iconic: true},
	{Name: "Chaosmancer", FreeToPlay: false, Iconic: true},
	{Name: "Dark Bargainer", FreeToPlay: false, Iconic: true},
	{Name: "Mindcleaver", FreeToPlay: false, Iconic: true},
	{Name: "Morninglord", FreeToPlay: false, Iconic: true},
	{Name: "Razorclaw", FreeToPlay: false, Iconic: true},
	{Name: "Scoundrel", FreeToPlay: false, Iconic: true},
	{Name: "Scurge", FreeToPlay: false, Iconic: true},
	{Name: "Shadar-kai", FreeToPlay: false, Iconic: true},
	{Name: "Trailblazer", FreeToPlay: false, Iconic: true},
}

var classesOptions = []Class{
	{Name: "Fighter", FreeToPlay: true, Archetype: ""},
	{Name: "Wizard", FreeToPlay: true, Archetype: ""},
	{Name: "Rogue", FreeToPlay: true, Archetype: ""},
	{Name: "Cleric", FreeToPlay: true, Archetype: ""},
	{Name: "Paladin", FreeToPlay: true, Archetype: ""},
	{Name: "Ranger", FreeToPlay: true, Archetype: ""},
	{Name: "Bard", FreeToPlay: true, Archetype: ""},
	{Name: "Monk", FreeToPlay: true, Archetype: ""},
	{Name: "Druid", FreeToPlay: true, Archetype: ""},
	{Name: "Barbarian", FreeToPlay: true, Archetype: ""},
	{Name: "Sorcerer", FreeToPlay: true, Archetype: ""},
	{Name: "Warlock", FreeToPlay: true, Archetype: ""},
	{Name: "Artificer", FreeToPlay: false, Archetype: ""},
	{Name: "Alchemist", FreeToPlay: false, Archetype: ""},
	{Name: "Favored Soul", FreeToPlay: false, Archetype: ""},
	{Name: "Stormsinger", FreeToPlay: true, Archetype: "Bard"},
	{Name: "Dark Apostate", FreeToPlay: true, Archetype: "Cleric"},
	{Name: "Sacred Fist", FreeToPlay: true, Archetype: "Paladin"},
	{Name: "Blightcaster", FreeToPlay: true, Archetype: "Druid"},
	{Name: "Dark Hunter", FreeToPlay: true, Archetype: "Ranger"},
	{Name: "Acolyte of the Skin", FreeToPlay: true, Archetype: "Warlock"},
	{Name: "Dragon Lord", FreeToPlay: false, Archetype: "Fighter"},
	{Name: "Dragon Disciple", FreeToPlay: false, Archetype: "Monk"},
	{Name: "Arcane Trickster", FreeToPlay: false, Archetype: "Rogue"},
	{Name: "Wild Mage", FreeToPlay: false, Archetype: "Sorcerer"},
}

func getFreeToPlayRacesOnly(items []Race) []Race {
	out := make([]Race, 0, len(items))
	for _, item := range items {
		if item.FreeToPlay {
			out = append(out, item)
		}
	}
	return out
}

func removeIconicRaces(items []Race) []Race {
	out := make([]Race, 0, len(items))
	for _, item := range items {
		if !item.Iconic {
			out = append(out, item)
		}
	}
	return out
}

func filterClassesByFreeToPlay(items []Class, wantFree bool) []Class {
	out := make([]Class, 0, len(items))
	for _, item := range items {
		if item.FreeToPlay == wantFree {
			out = append(out, item)
		}
	}
	return out
}

func getRandomRace(freeToPlay bool, iconicIncluded bool) string {
	pool := racesOptions

	if freeToPlay {
		pool = getFreeToPlayRacesOnly(pool)
	}
	if !iconicIncluded {
		pool = removeIconicRaces(pool)
	}

	return pool[rand.IntN(len(pool))].Name
}

func getRandomClass(freeToPlay bool) (string, string) {
	if !freeToPlay {
		randomIndex := rand.IntN(len(classesOptions))
		return classesOptions[randomIndex].Name, classesOptions[randomIndex].Archetype
	}
	filteredClasses := filterClassesByFreeToPlay(classesOptions, freeToPlay)
	randomIndex := rand.IntN(len(filteredClasses))
	return filteredClasses[randomIndex].Name, filteredClasses[randomIndex].Archetype
}

func canAddClass(selectedClasses []string, candidateClassName string, candidateArchetype string) bool {
	for _, selected := range selectedClasses {
		if selected == candidateClassName {
			return false
		}

		if candidateArchetype != "" && selected == candidateArchetype {
			return false
		}

		for _, classOption := range classesOptions {
			if classOption.Name == selected && classOption.Archetype == candidateClassName {
				return false
			}
		}
	}

	return true
}

func getLevels(numberOfClasses int) [3]int {
	switch numberOfClasses {
	case 1:
		return [3]int{20, 0, 0}
	case 2:
		total := 20
		first := rand.IntN(19) + 1
		second := total - first
		return [3]int{first, second, 0}
	case 3:
		total := 20
		first := rand.IntN(19) + 1
		second := rand.IntN(total-first-1) + 1
		third := total - first - second
		return [3]int{first, second, third}
	default:
		println("Invalid number of classes")
		return [3]int{0, 0, 0}
	}
}

func generateClasses(freeToPlayRaces bool, numClasses int) []string {
	var generatedClasses []string

	for i := 0; i < numClasses; i++ {
		for {
			randomClassName, randomClassArchetype := getRandomClass(freeToPlayRaces)
			// Ensure the randomly selected class is not already in the generated classes list
			if canAddClass(generatedClasses, randomClassName, randomClassArchetype) {
				generatedClasses = append(generatedClasses, randomClassName)
				break
			}
		}
	}
	return generatedClasses
}

func askForFreeToPlay() bool {
	fmt.Println("")
	fmt.Println("Free to Play only?")
	fmt.Println("1. Only Free to Play")
	fmt.Println("2. Include Premium Races")
	var freeToPlayChoice int
	_, err := fmt.Scanln(&freeToPlayChoice)
	if err != nil {
		fmt.Println("--------------------------------------")
		fmt.Println("Invalid input. Please enter a whole number, for example: 1")
		fmt.Println("--------------------------------------")
		fmt.Println()
		var discard string
		fmt.Scanln(&discard)
		return askForFreeToPlay()
	}
	switch freeToPlayChoice {
	case 1:
		return true
	case 2:
		return false
	default:
		fmt.Println("--------------------------------------")
		fmt.Println("Invalid input. Please enter 1 or 2")
		fmt.Println("--------------------------------------")
		fmt.Println()
		return askForFreeToPlay()
	}
}

func askForIconicRace() bool {
	fmt.Println("")
	fmt.Println("Include Iconic Races?")
	fmt.Println("1. Yes, include Iconic Races")
	fmt.Println("2. No, do not include Iconic Races")
	var iconicRacesChoice int
	_, err := fmt.Scanln(&iconicRacesChoice)
	if err != nil {
		fmt.Println("--------------------------------------")
		fmt.Println("Invalid input. Please enter a whole number, for example: 1")
		fmt.Println("--------------------------------------")
		fmt.Println()
		var discard string
		fmt.Scanln(&discard)
		return askForIconicRace()
	}
	switch iconicRacesChoice {
	case 1:
		return true
	case 2:
		return false
	default:
		fmt.Println("--------------------------------------")
		fmt.Println("Invalid input. Please enter 1 or 2")
		fmt.Println("--------------------------------------")
		fmt.Println()
		return askForIconicRace()
	}
}

func askForMulticlass() int {
	fmt.Println("")
	fmt.Println("Do you want to go pure or multiclass? Select how many classes you want to use:")
	fmt.Println("1. Pure")
	fmt.Println("2. Two classes")
	fmt.Println("3. Three classes")
	fmt.Println("4. Randomize how many classes")
	var multiclassChoice int
	_, err := fmt.Scanln(&multiclassChoice)
	if err != nil {
		fmt.Println("--------------------------------------")
		fmt.Println("Invalid input. Please enter a whole number, for example: 1")
		fmt.Println("--------------------------------------")
		fmt.Println()
		var discard string
		fmt.Scanln(&discard)
		return askForMulticlass()
	}
	switch multiclassChoice {
	case 1, 2, 3:
		return multiclassChoice
	case 4:
		multiclassChoice = rand.IntN(3) + 1 // 1..3
		return multiclassChoice
	default:
		fmt.Println("--------------------------------------")
		fmt.Println("Invalid input. Please enter 1, 2, 3, or 4")
		fmt.Println("--------------------------------------")
		fmt.Println()
		return askForMulticlass()
	}
}

func main() {
	fmt.Println()
	fmt.Println("WELCOME TO THE RANDOM BUILD GENERATOR!")
	fmt.Println()

	freeToPlayOnly := askForFreeToPlay()
	iconicRacesIncluded := askForIconicRace()
	multiclassChoice := askForMulticlass()
	generatedRace := getRandomRace(freeToPlayOnly, iconicRacesIncluded)

	// Generate the output
	fmt.Println()
	fmt.Println("GENERATED BUILD:")
	fmt.Println("Race:", generatedRace)
	switch multiclassChoice {
	case 1:
		generatedClass := generateClasses(freeToPlayOnly, 1)
		fmt.Printf("Class: %s [20]\n", generatedClass[0])
	case 2:
		// Implement logic for two classes
		generatedClasses := generateClasses(freeToPlayOnly, 2)
		levels := getLevels(2) // Example call to getLevels function for two classes
		fmt.Printf("Classes: %s [%d] and %s [%d]\n", generatedClasses[0], levels[0], generatedClasses[1], levels[1])
	case 3:
		// Implement logic for three classes
		generatedClasses := generateClasses(freeToPlayOnly, 3)
		levels := getLevels(3) // Example call to getLevels function for three classes
		fmt.Printf("Classes: %s [%d], %s [%d] and %s [%d]\n", generatedClasses[0], levels[0], generatedClasses[1], levels[1], generatedClasses[2], levels[2])
	}
}
