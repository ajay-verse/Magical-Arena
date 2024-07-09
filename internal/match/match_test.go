package match

import (
	// Go Internal Packages
	"testing"

	// Local Packages
	player "magicalarena/internal/player"
	
)

func TestCreateNewMatch(t *testing.T) {
	player1, _ := player.NewPlayer("Player 1", 100, 5, 10)
	player2, _ := player.NewPlayer("Player 1", 50, 5, 10)

	match := CreateNewMatch(player1, player2)

	if match.attacker != player2 {
		t.Errorf("Expected attacker to be Player 2, got %s", match.attacker.Name())
	}
	if match.defender != player1 {
		t.Errorf("Expected defender to be Player 1, got %s", match.defender.Name())
	}
}

func TestFight(t *testing.T) {
	player1, _ := player.NewPlayer("Player 1", 100, 5, 10)
	player2, _ := player.NewPlayer("Player 2", 50, 10, 5)

	match := CreateNewMatch(player1, player2)
	winner := match.Fight()

	if winner != "Player 1" && winner != "Player 2" {
		t.Errorf("Expected winner to be either Player 1 or Player 2, got %s", winner)
	}
}
