package main

import "testing"

func TestGetLevelFromExperiencePoints_HeroicLowTR_Level1Rank0(t *testing.T) {
	experiencePoints := 0

	level, rank, xpToNextLevel := getLevelFromExperiencePoints(heroicLevelsLowTR, experiencePoints)

	if level != 1 {
		t.Fatalf("expected level 1, got %d", level)
	}

	if rank != 0 {
		t.Fatalf("expected rank 0, got %d", rank)
	}

	if xpToNextLevel != 4000 {
		t.Fatalf("expected 4000 XP to next level, got %d", xpToNextLevel)
	}
}

func TestGetLevelFromExperiencePoints_HeroicLowTR_Level1Rank2(t *testing.T) {
	experiencePoints := 1600

	level, rank, xpToNextLevel := getLevelFromExperiencePoints(heroicLevelsLowTR, experiencePoints)

	if level != 1 {
		t.Fatalf("expected level 1, got %d", level)
	}

	if rank != 2 {
		t.Fatalf("expected rank 2, got %d", rank)
	}

	if xpToNextLevel != 2400 {
		t.Fatalf("expected 2400 XP to next level, got %d", xpToNextLevel)
	}
}

func TestGetLevelFromExperiencePoints_HeroicMidTR_Level10Rank1(t *testing.T) {
	experiencePoints := 650000

	level, rank, xpToNextLevel := getLevelFromExperiencePoints(heroicLevelsMidTR, experiencePoints)

	if level != 10 {
		t.Fatalf("expected level 10, got %d", level)
	}

	if rank != 1 {
		t.Fatalf("expected rank 1, got %d", rank)
	}

	if xpToNextLevel != 115000 {
		t.Fatalf("expected 115000 XP to next level, got %d", xpToNextLevel)
	}
}

func TestGetLevelFromExperiencePoints_HeroicHighTR_Level20Rank0(t *testing.T) {
	experiencePoints := 3800000

	level, rank, xpToNextLevel := getLevelFromExperiencePoints(heroicLevelsHighTR, experiencePoints)

	if level != 20 {
		t.Fatalf("expected level 20, got %d", level)
	}

	if rank != 0 {
		t.Fatalf("expected rank 0, got %d", rank)
	}

	if xpToNextLevel != 0 {
		t.Fatalf("expected 0 XP to next level, got %d", xpToNextLevel)
	}
}

func TestGetLevelFromExperiencePoints_EpicTR_Level29Rank4(t *testing.T) {
	experiencePoints := 8249999

	level, rank, xpToNextLevel := getLevelFromExperiencePoints(epicLevels, experiencePoints)

	if level != 29 {
		t.Fatalf("expected level 29, got %d", level)
	}

	if rank != 4 {
		t.Fatalf("expected rank 4, got %d", rank)
	}

	if xpToNextLevel != 1 {
		t.Fatalf("expected 1 XP to next level, got %d", xpToNextLevel)
	}
}

func TestGetLevelFromExperiencePoints_MaxLevel(t *testing.T) {
	experiencePoints := 12600000

	level, rank, xpToNextLevel := getLevelFromExperiencePoints(legendaryLevels, experiencePoints)

	if level != 36 {
		t.Fatalf("expected level 36, got %d", level)
	}

	if rank != 0 {
		t.Fatalf("expected rank 0, got %d", rank)
	}

	if xpToNextLevel != 0 {
		t.Fatalf("expected 0 XP to next level, got %d", xpToNextLevel)
	}
}
