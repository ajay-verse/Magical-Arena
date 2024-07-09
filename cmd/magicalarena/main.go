package main

import (
	match "magicalarena/internal/match"
	player "magicalarena/internal/player"
)

func main() {
	println("Welcome to Magical Arena")

	player1 := player.NewPlayer("Bhairava", 50, 5, 10)
	player2 := player.NewPlayer("Supreme Yaskin", 100, 10, 5)

	match1 := match.CreateNewMatch(player1, player2)
	res := match1.Fight()
	println(res, "Wins")
}
