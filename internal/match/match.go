package match

import (
	player "magicalarena/internal/player"
)

type Match struct {
	attacker *player.Player
	defender *player.Player
}

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
