package match

import (
	// Go Internal Packages
	"testing"

	// Local Packages
	player "magicalarena/internal/player"
)

// TestCreateNewMatch tests the CreateNewMatch function.
// It creates two players with different initial health points and strengths.
// Then it creates a new match with these players and checks if the attacker and defender are set correctly.
func TestCreateNewMatch(t *testing.T) {
	player1, _ := player.NewPlayer("Player 1", 100, 5, 10)
	player2, _ := player.NewPlayer("Player 2", 50, 5, 10)

	match := CreateNewMatch(player1, player2)

	if match.attacker != player2 {
		t.Errorf("Expected attacker to be Player 2, got %s", match.attacker.Name())
	}
	if match.defender != player1 {
		t.Errorf("Expected defender to be Player 1, got %s", match.defender.Name())
	}
}

// TestFight tests the Fight method of the Match struct.
// It creates two players with different initial health points and strengths.
// Then it creates a new match with these players and simulates a fight.
// The function checks if the winner is either "Player 1" or "Player 2".
// If the winner is neither, it fails the test with an error message.
func TestFight(t *testing.T) {
	player1, _ := player.NewPlayer("Player 1", 100, 5, 10)
	player2, _ := player.NewPlayer("Player 2", 50, 10, 5)

	match := CreateNewMatch(player1, player2)
	winner := match.Fight()

	if winner != "Player 1" && winner != "Player 2" {
		t.Errorf("Expected winner to be either Player 1 or Player 2, got %s", winner)
	}
}
