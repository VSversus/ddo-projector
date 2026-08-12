package main

import (
	"testing"
)

type testCase struct {
	name             string
	levelData        []LevelData
	experiencePoints int
	expectedLevel    int
	expectedRank     int
	expectedXPToNext int
}

func TestGetLevelFromExperiencePoints_TableDriven(t *testing.T) {

	testCases := []testCase{
		{
			name:             "Heroic Low TR - Level 1 Rank 0",
			levelData:        heroicLevelsLowTR,
			experiencePoints: 0,
			expectedLevel:    1,
			expectedRank:     0,
			expectedXPToNext: 4000,
		},
		{
			name:             "Heroic Low TR - Level 1 Rank 2",
			levelData:        heroicLevelsLowTR,
			experiencePoints: 1600,
			expectedLevel:    1,
			expectedRank:     2,
			expectedXPToNext: 2400,
		},
		{
			name:             "Heroic Mid TR - Level 10 Rank 1",
			levelData:        heroicLevelsMidTR,
			experiencePoints: 650000,
			expectedLevel:    10,
			expectedRank:     1,
			expectedXPToNext: 115000,
		},
		{
			name:             "Heroic High TR - Level 20 Rank 0",
			levelData:        heroicLevelsHighTR,
			experiencePoints: 3800000,
			expectedLevel:    20,
			expectedRank:     0,
			expectedXPToNext: 0,
		},
		{
			name:             "Epic TR - Level 29 Rank 4",
			levelData:        epicLevels,
			experiencePoints: 8249999,
			expectedLevel:    29,
			expectedRank:     4,
			expectedXPToNext: 1,
		},
		{
			name:             "Max xp - Level 36 Rank 0",
			levelData:        legendaryLevels,
			experiencePoints: 12600000,
			expectedLevel:    36,
			expectedRank:     0,
			expectedXPToNext: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			level, rank, xpToNextLevel := getLevelFromExperiencePoints(tc.levelData, tc.experiencePoints)

			if level != tc.expectedLevel {
				t.Fatalf("expected level %d, got %d", tc.expectedLevel, level)
			}

			if rank != tc.expectedRank {
				t.Fatalf("expected rank %d, got %d", tc.expectedRank, rank)
			}

			if xpToNextLevel != tc.expectedXPToNext {
				t.Fatalf("expected %d XP to next level, got %d", tc.expectedXPToNext, xpToNextLevel)
			}
		})
	}
}
