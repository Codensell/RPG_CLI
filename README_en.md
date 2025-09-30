## RPG CLI — Turn-Based Auto-Battler (Go + ncurses)

#### A mini terminal game: create your character, choose a class, and battles proceed automatically. Your role is to make decisions between fights — upgrading and selecting weapons. The goal is to win 5 battles in a row.

#### Features

Terminal UI using ncurses: player/enemy frames, battle log in the center.

Turn-based auto-combat with a 2-second timer between actions.

Initiative system, hit/miss mechanic, and detailed damage pipeline.

#### Classes and Effects

Warrior: Action Surge (extra attack), Shield, STR +1 at level 3.

Barbarian: Rage, Stone Skin, STA +1 at level 3.

Rogue: Sneak Attack, AGI +1 at level 2, Poison at level 3.

Enemy traits (immunities, vulnerabilities, perks).

Loot system with weapon drops and modal for comparison/replacement.

Progression: multiclassing up to total level 3, healing after victory.

Series of 5 wins → victory screen.

Unit test suite covering key effects and calculations.

#### How to Play

Build and Run

Requirements:
Go 1.20+

Библиотека ncurses:

macOS: brew install ncurses

Ubuntu/Debian: sudo apt-get install libncurses5-dev libncursesw5-dev

Arch: sudo pacman -S ncurses

Windows: WSL/WSL2 с Linux-дистрибутивом

go run ./cmd/game

#### How Combat Works

Initiative: The one with higher AGI goes first; player wins ties.

Hit/Miss: Random roll in [1 .. AGI_att + AGI_def]; ≤ AGI_def means miss.

Damage pipeline:

Base: weapon + STR.

Offensive buffs: Action Surge, Rage, Sneak Attack, Dragon Breath.

Partial defenses: immunities/vulnerabilities affecting weapon component only
(e.g., immunity to slashing nullifies slashing weapon damage; vulnerability to blunt doubles it).

Global defenses: Shield (−3 if STR_def > STR_att), Stone Skin (−STA_def).

Bottom cap: final damage can't go below 0.

Poison: applies on hit if perk is active; ticks on victim's turns: +1 on turn 2, +2 on turn 3 and beyond (after global defenses).

Timing: One attack attempt every ~2 seconds (UI controls timing; calculations handled in domain layer).

#### Progression & Loot

Level-Up: Increases MaxHP by (class HP per level) + STA. Attribute bonuses apply first, then HP is calculated.

Perks: Unlocked by class level.

Loot: After victory, prompt to replace current weapon with enemy drop (with modal comparison).
Drop table is defined in internal/loot and easily customizable (via LootTable interface).

#### Testing

From the project root:

go test ./internal/domain -v

Covers:

Initiative tie-breaker.

Action Surge (bonus not blocked by weapon immunity).

Vulnerability to blunt (affects weapon part only).

Rage per turn (+2/+2/+2/−1).

Dragon Breath every 3rd turn.

Shield / Stone Skin as global reductions.

Damage floor at zero.

Correct poison ticking.

#### Extensibility

SOLID Principles:

SRP: domain handles pure logic, ui/term handles rendering only, loot holds game data.

OCP: New enemies/effects/weapons are added via isolated data/extensions, no rewriting needed.

DIP: Loot uses an interface — alternative implementations (e.g. weighted drops) can be plugged in without changing core logic.

New effects are best added as flags/parameters in Perks + handled in the damage pipeline.
UI remains unaware of these effects and doesn't require modification.