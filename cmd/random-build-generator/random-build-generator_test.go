package main

import (
	"slices"
	"testing"
)

type testCasesRaces struct {
	name     string
	input    []Race
	expected []Race
}

func TestGetFreeToPlayRacesOnly(t *testing.T) {

	testCases := []testCasesRaces{
		{
			name: "no race returned if none are free to play",
			input: []Race{
				{Name: "Aasimar", FreeToPlay: false},
				{Name: "Shifter", FreeToPlay: false},
				{Name: "Warforged", FreeToPlay: false},
			},
			expected: []Race{},
		},
		{
			name: "only free to play races are returned",
			input: []Race{
				{Name: "Human", FreeToPlay: true},
				{Name: "Aasimar", FreeToPlay: false},
				{Name: "Elf", FreeToPlay: true},
			},
			expected: []Race{
				{Name: "Human", FreeToPlay: true},
				{Name: "Elf", FreeToPlay: true},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := getFreeToPlayRacesOnly(tc.input)

			if len(result) != len(tc.expected) {
				t.Fatalf("expected %d races, got %d", len(tc.expected), len(result))
			}

			for index, race := range result {
				if race != tc.expected[index] {
					t.Errorf("expected race %v at index %d, got %v", tc.expected[index], index, race)
				}
			}
		})
	}
}

func TestRemoveIconicRaces(t *testing.T) {

	testCases := []testCasesRaces{
		{
			name: "no race returned if none are iconic",
			input: []Race{
				{Name: "Deep Gnome", Iconic: true},
				{Name: "Shadar-kai", Iconic: true},
				{Name: "Purple Dragon Knight", Iconic: true},
			},
			expected: []Race{},
		},
		{
			name: "only non-iconic races are returned",
			input: []Race{
				{Name: "Human", Iconic: false},
				{Name: "Aasimar", Iconic: false},
				{Name: "Mindcleaver", Iconic: true},
			},
			expected: []Race{
				{Name: "Human", Iconic: false},
				{Name: "Aasimar", Iconic: false},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := removeIconicRaces(tc.input)

			if len(result) != len(tc.expected) {
				t.Fatalf("expected %d races, got %d", len(tc.expected), len(result))
			}

			for index, race := range result {
				if race != tc.expected[index] {
					t.Errorf("expected race %v at index %d, got %v", tc.expected[index], index, race)
				}
			}
		})
	}
}

func TestFilterClassesByFreeToPlay(t *testing.T) {

	type testCasesClasses struct {
		name     string
		input    []Class
		expected []Class
	}

	testCases := []testCasesClasses{
		{
			name: "no class returned if none are free to play",
			input: []Class{
				{Name: "Alchemist", FreeToPlay: false},
				{Name: "Wild Mage", FreeToPlay: false},
			},
			expected: []Class{},
		},
		{
			name: "only free to play classes are returned",
			input: []Class{
				{Name: "Fighter", FreeToPlay: true},
				{Name: "Alchemist", FreeToPlay: false},
				{Name: "Rogue", FreeToPlay: true},
			},
			expected: []Class{
				{Name: "Fighter", FreeToPlay: true},
				{Name: "Rogue", FreeToPlay: true},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := filterClassesByFreeToPlay(tc.input, true)

			if len(result) != len(tc.expected) {
				t.Fatalf("expected %d classes, got %d", len(tc.expected), len(result))
			}

			for index, class := range result {
				if class != tc.expected[index] {
					t.Errorf("expected class %v at index %d, got %v", tc.expected[index], index, class)
				}
			}
		})
	}
}

func TestIsAlignmentCompatible(t *testing.T) {

	type testCasesAlignment struct {
		name               string
		candidateAlignment CompatibleAlignment
		selectedAlignments []CompatibleAlignment
		expected           bool
	}

	testCases := []testCasesAlignment{
		{
			name:               "candidate alignment compatible",
			candidateAlignment: CompatibleAlignment{LawfulGood: true},
			selectedAlignments: []CompatibleAlignment{{LawfulGood: true}},
			expected:           true,
		},
		{
			name:               "candidate alignment incompatible",
			candidateAlignment: CompatibleAlignment{ChaoticGood: true},
			selectedAlignments: []CompatibleAlignment{{LawfulGood: true}},
			expected:           false,
		},
		{
			name:               "candidate alignment compatible with just one selected alignment",
			candidateAlignment: CompatibleAlignment{ChaoticGood: true},
			selectedAlignments: []CompatibleAlignment{{ChaoticGood: true}, {LawfulGood: true}},
			expected:           false,
		},
		{
			name: "candidate alignment compatible with everything",
			candidateAlignment: CompatibleAlignment{
				LawfulGood:     true,
				LawfulNeutral:  true,
				NeutralGood:    true,
				TrueNeutral:    true,
				ChaoticGood:    true,
				ChaoticNeutral: true,
			},
			selectedAlignments: []CompatibleAlignment{{ChaoticGood: true}, {ChaoticGood: true}},
			expected:           true,
		},
		{
			name: "candidate compatible with everything but selected conflicting with each other",
			candidateAlignment: CompatibleAlignment{
				LawfulGood:     true,
				LawfulNeutral:  true,
				NeutralGood:    true,
				TrueNeutral:    true,
				ChaoticGood:    true,
				ChaoticNeutral: true,
			},
			selectedAlignments: []CompatibleAlignment{{LawfulGood: true}, {LawfulNeutral: true}},
			expected:           false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isAlignmentCompatible(tc.candidateAlignment, tc.selectedAlignments)

			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestGetAvailableRaces(t *testing.T) {

	type testCasesGetRandomRace struct {
		name               string
		freeToPlay         bool
		iconicIncluded     bool
		expectedOutputPool []Race
	}

	var inputPool = []Race{
		{Name: "Race1", FreeToPlay: true, Iconic: false},
		{Name: "Race2", FreeToPlay: false, Iconic: false},
		{Name: "Race3", FreeToPlay: true, Iconic: true},
		{Name: "Race4", FreeToPlay: false, Iconic: true},
	}

	testCases := []testCasesGetRandomRace{
		{
			name:           "freeToPlay=false, iconicIncluded=true",
			freeToPlay:     false,
			iconicIncluded: true,
			expectedOutputPool: []Race{
				{Name: "Race1", FreeToPlay: true, Iconic: false},
				{Name: "Race2", FreeToPlay: false, Iconic: false},
				{Name: "Race3", FreeToPlay: true, Iconic: true},
				{Name: "Race4", FreeToPlay: false, Iconic: true}},
		},
		{
			name:           "freeToPlay=true, iconicIncluded=true",
			freeToPlay:     true,
			iconicIncluded: true,
			expectedOutputPool: []Race{
				{Name: "Race1", FreeToPlay: true, Iconic: false},
				{Name: "Race3", FreeToPlay: true, Iconic: true},
			},
		},
		{
			name:           "freeToPlay=false, iconicIncluded=false",
			freeToPlay:     false,
			iconicIncluded: false,
			expectedOutputPool: []Race{
				{Name: "Race1", FreeToPlay: true, Iconic: false},
				{Name: "Race2", FreeToPlay: false, Iconic: false},
			},
		},
		{
			name:           "freeToPlay=true, iconicIncluded=false",
			freeToPlay:     true,
			iconicIncluded: false,
			expectedOutputPool: []Race{
				{Name: "Race1", FreeToPlay: true, Iconic: false},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := getAvailableRaces(inputPool, tc.freeToPlay, tc.iconicIncluded)
			if !slices.Equal(result, tc.expectedOutputPool) {
				t.Errorf("Test %s failed: expected %v, got %v", tc.name, tc.expectedOutputPool, result)
			}
		})
	}
}

func TestGetRandomRace(t *testing.T) {

	racePool := []Race{
		{Name: "Human", FreeToPlay: true, Iconic: false},
	}

	result := getRandomRace(racePool, true, false)

	if result != "Human" {
		t.Errorf("expected Human, got %s", result)
	}
}

func TestGetRandomClass(t *testing.T) {

	classPool := []Class{
		{Name: "Fighter"},
	}

	result := getRandomClass(classPool)

	if result != "Fighter" {
		t.Errorf("expected Fighter, got %s", result)
	}
}

func TestCanAddClass(t *testing.T) {

	type testCasesCanAddClass struct {
		name            string
		selectedClasses []string
		candidateClass  string
		expectedOutput  bool
	}

	testCases := []testCasesCanAddClass{
		{
			name:            "can add class when selected empty",
			selectedClasses: []string{},
			candidateClass:  "Wizard",
			expectedOutput:  true,
		},
		{
			name:            "can add class when different from selected",
			selectedClasses: []string{"Fighter", "Ranger"},
			candidateClass:  "Wizard",
			expectedOutput:  true,
		},
		{
			name:            "cannot add class when it is already selected",
			selectedClasses: []string{"Wizard", "Fighter"},
			candidateClass:  "Wizard",
			expectedOutput:  false,
		},
		{
			name:            "cannot add archetype when original class already selected",
			selectedClasses: []string{"Ranger", "Fighter"},
			candidateClass:  "Dark Hunter",
			expectedOutput:  false,
		},
		{
			name:            "cannot add class when archetype already selected",
			selectedClasses: []string{"Dark Hunter", "Warlock"},
			candidateClass:  "Ranger",
			expectedOutput:  false,
		},
		{
			name:            "cannot add archetype when it is already selected",
			selectedClasses: []string{"Dark Hunter"},
			candidateClass:  "Dark Hunter",
			expectedOutput:  false,
		},
		{
			name:            "can add archetype when a different archetype is already selected",
			selectedClasses: []string{"Dark Hunter", "Artificer"},
			candidateClass:  "Blightcaster",
			expectedOutput:  true,
		},
		{
			name:            "candidate shares at least one valid alignment with all selected classes",
			selectedClasses: []string{"Fighter", "Rogue"},
			candidateClass:  "Paladin",
			expectedOutput:  true,
		},
		{
			name:            "candidate does not share a valid alignment with one of the selected classes",
			selectedClasses: []string{"Fighter", "Barbarian"},
			candidateClass:  "Paladin",
			expectedOutput:  false,
		},
		{
			name:            "candidate does not share a valid alignment with each of the selected classes",
			selectedClasses: []string{"Druid", "Barbarian"},
			candidateClass:  "Paladin",
			expectedOutput:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := canAddClass(tc.selectedClasses, tc.candidateClass)

			if result != tc.expectedOutput {
				t.Errorf("expected %v, got %v", tc.expectedOutput, result)
			}
		})
	}
}
