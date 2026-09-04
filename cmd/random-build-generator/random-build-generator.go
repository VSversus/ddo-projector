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
	Alignment  Alignment
}

type Alignment struct {
	LawfulGood     bool
	LawfulNeutral  bool
	NeutralGood    bool
	TrueNeutral    bool
	ChaoticGood    bool
	ChaoticNeutral bool
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
	{Name: "Fighter", FreeToPlay: true, Archetype: "", Alignment: Alignment{LawfulGood: true, LawfulNeutral: true, NeutralGood: true, TrueNeutral: true, ChaoticGood: true, ChaoticNeutral: true}},
	{Name: "Wizard", FreeToPlay: true, Archetype: "", Alignment: Alignment{LawfulGood: true, LawfulNeutral: true, NeutralGood: true, TrueNeutral: true, ChaoticGood: true, ChaoticNeutral: true}},
	{Name: "Rogue", FreeToPlay: true, Archetype: "", Alignment: Alignment{LawfulGood: true, LawfulNeutral: true, NeutralGood: true, TrueNeutral: true, ChaoticGood: true, ChaoticNeutral: true}},
	{Name: "Cleric", FreeToPlay: true, Archetype: "", Alignment: Alignment{LawfulGood: true, LawfulNeutral: true, NeutralGood: true, TrueNeutral: true, ChaoticGood: true, ChaoticNeutral: true}},
	{Name: "Paladin", FreeToPlay: true, Archetype: "", Alignment: Alignment{LawfulGood: true, LawfulNeutral: false, NeutralGood: false, TrueNeutral: false, ChaoticGood: false, ChaoticNeutral: false}},
	{Name: "Ranger", FreeToPlay: true, Archetype: "", Alignment: Alignment{LawfulGood: true, LawfulNeutral: true, NeutralGood: true, TrueNeutral: true, ChaoticGood: true, ChaoticNeutral: true}},
	{Name: "Bard", FreeToPlay: true, Archetype: "", Alignment: Alignment{LawfulGood: false, LawfulNeutral: false, NeutralGood: true, TrueNeutral: true, ChaoticGood: true, ChaoticNeutral: true}},
	{Name: "Monk", FreeToPlay: true, Archetype: "", Alignment: Alignment{LawfulGood: true, LawfulNeutral: true, NeutralGood: false, TrueNeutral: false, ChaoticGood: false, ChaoticNeutral: false}},
	{Name: "Druid", FreeToPlay: true, Archetype: "", Alignment: Alignment{LawfulGood: false, LawfulNeutral: true, NeutralGood: true, TrueNeutral: true, ChaoticGood: false, ChaoticNeutral: true}},
	{Name: "Barbarian", FreeToPlay: true, Archetype: "", Alignment: Alignment{LawfulGood: false, LawfulNeutral: false, NeutralGood: true, TrueNeutral: true, ChaoticGood: true, ChaoticNeutral: true}},
	{Name: "Sorcerer", FreeToPlay: true, Archetype: "", Alignment: Alignment{LawfulGood: true, LawfulNeutral: true, NeutralGood: true, TrueNeutral: true, ChaoticGood: true, ChaoticNeutral: true}},
	{Name: "Warlock", FreeToPlay: true, Archetype: "", Alignment: Alignment{LawfulGood: true, LawfulNeutral: true, NeutralGood: true, TrueNeutral: true, ChaoticGood: true, ChaoticNeutral: true}},
	{Name: "Artificer", FreeToPlay: false, Archetype: "", Alignment: Alignment{LawfulGood: true, LawfulNeutral: true, NeutralGood: true, TrueNeutral: true, ChaoticGood: true, ChaoticNeutral: true}},
	{Name: "Alchemist", FreeToPlay: false, Archetype: "", Alignment: Alignment{LawfulGood: true, LawfulNeutral: true, NeutralGood: true, TrueNeutral: true, ChaoticGood: true, ChaoticNeutral: true}},
	{Name: "Favored Soul", FreeToPlay: false, Archetype: "", Alignment: Alignment{LawfulGood: true, LawfulNeutral: true, NeutralGood: true, TrueNeutral: true, ChaoticGood: true, ChaoticNeutral: true}},
	{Name: "Stormsinger", FreeToPlay: true, Archetype: "Bard", Alignment: Alignment{LawfulGood: true, LawfulNeutral: true, NeutralGood: true, TrueNeutral: true, ChaoticGood: true, ChaoticNeutral: true}},
	{Name: "Dark Apostate", FreeToPlay: true, Archetype: "Cleric", Alignment: Alignment{LawfulGood: true, LawfulNeutral: true, NeutralGood: true, TrueNeutral: true, ChaoticGood: true, ChaoticNeutral: true}},
	{Name: "Sacred Fist", FreeToPlay: true, Archetype: "Paladin", Alignment: Alignment{LawfulGood: true, LawfulNeutral: true, NeutralGood: true, TrueNeutral: true, ChaoticGood: true, ChaoticNeutral: true}},
	{Name: "Blightcaster", FreeToPlay: true, Archetype: "Druid", Alignment: Alignment{LawfulGood: true, LawfulNeutral: true, NeutralGood: true, TrueNeutral: true, ChaoticGood: true, ChaoticNeutral: true}},
	{Name: "Dark Hunter", FreeToPlay: true, Archetype: "Ranger", Alignment: Alignment{LawfulGood: true, LawfulNeutral: true, NeutralGood: true, TrueNeutral: true, ChaoticGood: true, ChaoticNeutral: true}},
	{Name: "Acolyte of the Skin", FreeToPlay: true, Archetype: "Warlock", Alignment: Alignment{LawfulGood: false, LawfulNeutral: true, NeutralGood: false, TrueNeutral: false, ChaoticGood: false, ChaoticNeutral: true}},
	{Name: "Dragon Lord", FreeToPlay: false, Archetype: "Fighter", Alignment: Alignment{LawfulGood: true, LawfulNeutral: true, NeutralGood: true, TrueNeutral: true, ChaoticGood: true, ChaoticNeutral: true}},
	{Name: "Dragon Disciple", FreeToPlay: false, Archetype: "Monk", Alignment: Alignment{LawfulGood: true, LawfulNeutral: true, NeutralGood: true, TrueNeutral: true, ChaoticGood: true, ChaoticNeutral: true}},
	{Name: "Arcane Trickster", FreeToPlay: false, Archetype: "Rogue", Alignment: Alignment{LawfulGood: true, LawfulNeutral: true, NeutralGood: true, TrueNeutral: true, ChaoticGood: true, ChaoticNeutral: true}},
	{Name: "Wild Mage", FreeToPlay: false, Archetype: "Sorcerer", Alignment: Alignment{LawfulGood: true, LawfulNeutral: true, NeutralGood: true, TrueNeutral: true, ChaoticGood: true, ChaoticNeutral: true}},
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

func isAlignmentCompatible(candidateAlignment Alignment, selectedAlignments []Alignment) bool {
	compatible := candidateAlignment

	for _, selected := range selectedAlignments {
		compatible.LawfulGood = compatible.LawfulGood && selected.LawfulGood
		compatible.LawfulNeutral = compatible.LawfulNeutral && selected.LawfulNeutral
		compatible.NeutralGood = compatible.NeutralGood && selected.NeutralGood
		compatible.TrueNeutral = compatible.TrueNeutral && selected.TrueNeutral
		compatible.ChaoticGood = compatible.ChaoticGood && selected.ChaoticGood
		compatible.ChaoticNeutral = compatible.ChaoticNeutral && selected.ChaoticNeutral
	}

	return compatible.LawfulGood ||
		compatible.LawfulNeutral ||
		compatible.NeutralGood ||
		compatible.TrueNeutral ||
		compatible.ChaoticGood ||
		compatible.ChaoticNeutral
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

func getRandomClass(availableClasses []Class) string {
	randomIndex := rand.IntN(len(availableClasses))
	return availableClasses[randomIndex].Name
}

func canAddClass(selectedClasses []string, candidateClassName string) bool {
	// get archetype and alignments of candidate class
	candidateAlignments := Alignment{}
	candidateArchetype := ""
	for _, class := range classesOptions {
		if class.Name == candidateClassName {
			candidateAlignments = class.Alignment
			candidateArchetype = class.Archetype
			break
		}
	}

	// get all alignments of the already selected classes
	selectedAlignments := []Alignment{}
	for _, selected := range selectedClasses {
		for _, class := range classesOptions {
			if class.Name == selected {
				selectedAlignments = append(selectedAlignments, class.Alignment)
				break
			}
		}
	}

	for _, selected := range selectedClasses {
		// Check if the selected class is the same as the candidate class (covers also the case when the selected archetype is the same as candidate archetype)
		if selected == candidateClassName {
			return false
		}
		// Check if the selected class is the same as the candidate archetype
		if candidateArchetype != "" && selected == candidateArchetype {
			return false
		}
		// Check if the selected archetype is the same as the candidate class
		for _, class := range classesOptions {
			if class.Name == selected && class.Archetype == candidateClassName {
				return false
			}
		}
		// check if the candidate allowed alignments is compatible with the already selected classes
		// selected class must have at least one alignment with true in common with the candidate class
		if !isAlignmentCompatible(candidateAlignments, selectedAlignments) {
			return false
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

func removeClass(classes []Class, className string) []Class {
	remaining := make([]Class, 0, len(classes))

	for _, class := range classes {
		if class.Name != className {
			remaining = append(remaining, class)
		}
	}

	return remaining
}

func generateClasses(freeToPlayOnly bool, numClasses int) []string {
	var generatedClasses []string
	var availableClasses []Class = classesOptions

	if freeToPlayOnly {
		availableClasses = filterClassesByFreeToPlay(availableClasses, true)
	}

	for i := 0; i < numClasses; i++ {
		for len(generatedClasses) < numClasses && len(availableClasses) > 0 {
			randomClassName := getRandomClass(availableClasses)
			// Ensure the randomly selected class is not already in the generated classes list
			if canAddClass(generatedClasses, randomClassName) {
				generatedClasses = append(generatedClasses, randomClassName)
				break
			} else {
				availableClasses = removeClass(availableClasses, randomClassName)
			}
		}
	}
	return generatedClasses
}

func askForFreeToPlay() bool {
	fmt.Println("")
	fmt.Println("Free to Play only?")
	fmt.Println("1. Include Premium Races")
	fmt.Println("2. Only Free to Play")
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
		return false
	case 2:
		return true
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
