# Magical Arena

## Problem Statement
Design a Magical Arena. Every Player is defined by a “health” attribute, “strength” attribute and an “attack” attribute - all positive integers. The player dies if his health attribute touches 0. 

After Reading the Problem Statement throughly I decided to do this project using Go

## Overview
I came up with this folder structure in order to complete the project
```
Folder Structure:
Magical Arena/
├── cmd/
│   └── magicalarena/
│       └── main.go
├── internal/
│   ├── player/
│   │   ├── player.go
│   │   └── player_test.go
│   ├── match/
│   │   ├── match.go
│   │   └── match_test.go
│   ├── dice/
│       ├── dice.go
├── go.mod
├── Problem Statement.txt
└── README.md
```

## Features
- Create and manage player characters with attributes such as name, health, strength, and attack.
- Implement a turn-based combat system where players take turns attacking each other.
- Calculate damage and defense based on player attributes and random dice rolls.
- Display colorful messages and ASCII art to enhance the game's visual appeal.

## Getting Started
To run the Magical Arena simulation, follow these steps:

1. Clone or download the project repository to your local machine.
2. Open the project in your preferred IDE such as VS Code or Goland.


## Running the Application
1. Move to the directory Magical Arena
```bash
cd ~/pathtorepository/Magical Arena
```
2. Command to run main.go
```bash
go run ./cmd/magicalarena/main.go
```
3. Command to run test files
```bash
go test ./...
```

## Magical Arena - Player Package

This package contains the implementation of the `Player` struct and related methods for the Magical Arena game.

### Player Struct

The `Player` struct represents a character in the game with the following attributes:

- `name`: A string representing the name of the player.
- `health`: An integer representing the health points of the player.
- `strength`: An integer representing the strength of the player.
- `attack`: An integer representing the attack power of the player.

### Methods

The package provides the following methods for the `Player` struct:

1. `NewPlayer(name string, health int, strength int, attack int) (*Player, error)`: Creates a new `Player` instance with the given name, health, strength, and attack. Returns an error if the name is empty or any attribute is negative.
2. `Name() string`: Returns the name of the player.
3. `Health() int`: Returns the health points of the player.
4. `Attack() int`: Returns the attack power of the player.
5. `Strength() int`: Returns the strength of the player.
6. `SetName(name string) error`: Sets the name of the player. Returns an error if the name is empty.
7. `SetHealth(health int) error`: Sets the health points of the player. Returns an error if the health is negative.
8. `SetAttack(attack int) error`: Sets the attack power of the player. Returns an error if the attack is negative.
9. `SetStrength(strength int) error`: Sets the strength of the player. Returns an error if the strength is negative.

### Example Usage

```go
package main

import (
    // Go Internal Packages
	"fmt"

	// Local Packages
	player "magicalarena/internal/player"
)

func main() {
    // Create a new player
    player1, err := player.NewPlayer("Player 1", 100, 50, 20)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Get player attributes
    fmt.Println("Player Name:", player1.Name())
    fmt.Println("Player Health:", player1.Health())
    fmt.Println("Player Strength:", player1.Strength())
    fmt.Println("Player Attack:", player1.Attack())

    // Set player attributes
    err = player1.SetHealth(80)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Updated Player Health:", player1.Health())
}
```

## Magical Arena - Match Package
This package contains the implementation of the `Match` struct and related methods for simulating a single match between two players in the Magical Arena game.

### Match Struct

The `Match` struct represents a single match between two players. It contains the following fields:

- `attacker`: A pointer to the `Player` struct representing the attacker in the match.
- `defender`: A pointer to the `Player` struct representing the defender in the match.

### Methods

The package provides the following methods for the `Match` struct:

1. `CreateNewMatch(player1 *player.Player, player2 *player.Player) *Match`: Creates a new match between two players. The player with lower health starts as the attacker.
2. `Fight() string`: Simulates a turn-based combat between the attacker and defender. The function continues until one of the players' health drops to zero. It returns the name of the winner.

### Example Usage

```go
package main

import (
    // Go Internal Packages
	"fmt"

	// Local Packages
    match "magicalarena/internal/match"
	player "magicalarena/internal/player"
)

func main() {
    // Create two players
    player1 := player.NewPlayer("Player 1", 100, 50, 20)
    player2 := player.NewPlayer("Player 2", 80, 60, 15)

    // Create a new match
    match1 := match.CreateNewMatch(player1, player2)

    // Simulate the match
    winner := match1.Fight()
    fmt.Println("Winner:", winner)
}
```
