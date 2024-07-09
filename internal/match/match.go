package match

import (
	player "magicalarena/internal/player"
	dice "magicalarena/internal/utils/dice"
	helpers "magicalarena/internal/utils/helpers"
)

// Match represents a single match between two players.
type Match struct {
	attacker *player.Player
	defender *player.Player
}

// CreateNewMatch creates a new match between two players.
// The player with lower health starts as the attacker.
func CreateNewMatch(player1 *player.Player, player2 *player.Player) *Match {
	var attacker, defender *player.Player
	if player1.Health() < player2.Health() {
		attacker = player1
		defender = player2
	} else {
		attacker = player2
		defender = player1
	}
	return &Match{
		attacker: attacker,
		defender: defender,
	}
}

// Fight simulates a turn-based combat between the attacker and defender.
// The function continues until one of the players' health drops to zero.
// It returns the name of the winner.
func (m *Match) Fight() string {
	var round int = 1
	for m.attacker.Health() > 0 && m.defender.Health() > 0 {
		helpers.PrintRoundName(round)

		attackRoll := dice.RollD6()
		defendRoll := dice.RollD6()
		helpers.PrintDiceRollDetails(m.attacker.Name(), attackRoll, m.defender.Name(), defendRoll)

		attackDamage := attackRoll * m.attacker.Attack()
		defendDamage := defendRoll * m.defender.Strength()
		helpers.PrintAttackDetails(attackDamage, defendDamage)

		netDamage := attackDamage - defendDamage
		if netDamage < 0 {
			netDamage = 0
		}

		m.defender.SetHealth(m.defender.Health() - netDamage)
		helpers.PrintHealthDetails(m.attacker.Health(), m.defender.Health())
		// Switch roles for next turn
		m.attacker, m.defender = m.defender, m.attacker
		round++
	}

	if m.attacker.Health() <= 0 {
		return m.defender.Name()
	} else {
		return m.attacker.Name()
	}
}
