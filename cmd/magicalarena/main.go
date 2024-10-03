package main

import (
	// Go Internal Packages
	"fmt"

	// Local Packages
	match "magicalarena/services/match"
	player "magicalarena/services/player"
)

func main() {
	player1, err1 := player.NewPlayer("Iron Man", 100, 5, 10)
	if err1 != nil {
		fmt.Println(err1)
		return
	}

	player2, err2 := player.NewPlayer("Spider Man", 50, 10, 5)
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	match1 := match.CreateNewMatch(player1, player2)
	winner := match1.Fight()
	fmt.Printf("🎉 Hurray! %s Won 🎉\n", winner)
}
