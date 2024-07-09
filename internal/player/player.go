package player

// Player represents a character in Magical Arena Game with attributes given.
type Player struct {
	name     string // Name of the player.
	health   int    // Health points of the player.
	strength int    // Strength of the player.
	attack   int    // Attack power of the player.
}

// NewPlayer creates a new Player instance with the given name, health, strength, and attack.
func NewPlayer(name string, health int, strength int, attack int) *Player {
	return &Player{
		name:     name,
		health:   health,
		strength: strength,
		attack:   attack,
	}
}
