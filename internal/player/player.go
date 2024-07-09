package player

import "errors"

// Player represents a character in Magical Arena Game with attributes given.
type Player struct {
	name     string // Name of the player.
	health   int    // Health points of the player.
	strength int    // Strength of the player.
	attack   int    // Attack power of the player.
}

// NewPlayer creates a new Player instance with the given name, health, strength, and attack.
func NewPlayer(name string, health int, strength int, attack int) (*Player, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}
	if health < 0  || strength < 0 || attack < 0 {
		return nil, errors.New("atrributes cannot be negative")
	}
	return &Player{
		name:     name,
		health:   health,
		strength: strength,
		attack:   attack,
	}, nil
}

// Name returns the name of the player.
func (p *Player) Name() string {
	return p.name
}

// Health returns the health points of the player.
func (p *Player) Health() int {
	return p.health
}

// Attack returns the attack power of the player.
func (p *Player) Attack() int {
	return p.attack
}

// Strength returns the strength of the player.
func (p *Player) Strength() int {
	return p.strength
}

// SetName sets the name of the player.
func (p *Player) SetName(name string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}
	p.name = name
	return nil
}

// SetHealth sets the health points of the player.
func (p *Player) SetHealth(health int) error {
	if health < 0 {
		return errors.New("health cannot be negative")
	}
	p.health = health
	return nil
}

// SetAttack sets the attack power of the player.
func (p *Player) SetAttack(attack int) error {
	if attack < 0 {
		return errors.New("attack cannot be negative")
	}
	p.attack = attack
	return nil
}

// SetStrength sets the strength of the player.
func (p *Player) SetStrength(strength int) error {
	if strength < 0 {
		return errors.New("strength cannot be negative")
	}
	p.strength = strength
	return nil
}
