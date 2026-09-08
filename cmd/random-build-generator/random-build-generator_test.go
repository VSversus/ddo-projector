package main

import "testing"

func TestGetFreeToPlayRacesOnly(t *testing.T) {
	testCases := []struct {
		name     string
		input    []Race
		expected []Race
	}{
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
