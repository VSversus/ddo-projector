## XP Calculator for DDO

A CLI tool to calculate character level, rank, and XP needed for next level in **Dungeons & Dragons Online** (DDO).

This calculator helps you quickly determine your character progression based on experience points. It accounts for:
- **Heroic levels** (1-20) with 3 difficulty tiers based on true reincarnations
- **Epic levels** (20-30)
- **Legendary levels** (30-36)

### Usage

```bash
go run xp-calculator.go
```

## Random Build Generator

A CLI tool that generates a random DDO character build. It can choose a race, filter races by free-to-play status and iconic status, and generate a pure or multiclass build with compatible classes.

### Usage

Run the generator from the repository root:

```bash
go run ./cmd/random-build-generator
```

The generator asks whether to include premium races, whether to include iconic races, and whether to create a pure, two-class, or three-class build. For multiclass builds, it avoids duplicate classes, incompatible archetypes, and classes without a shared alignment.