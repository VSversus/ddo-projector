package main

import "testing"

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
