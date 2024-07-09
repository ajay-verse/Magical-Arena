package helpers

import (
	// Go Internal Packages
	"fmt"

	// Local Packages
	player "magicalarena/internal/player"
	x "magicalarena/internal/utils/constants"
)

func PrintWelcomeMessage() {
	fmt.Println(x.CYAN + "\n\t✨ Welcome to Magical Arena ✨\n" + x.RESET)
}

func PrintPlayerDetails(player *player.Player) {
	fmt.Printf(x.CYAN + "🧙 %s Details 🧙\n" + x.RESET, player.Name())
	fmt.Printf(x.GREEN+"🩸 Health: %d  "+x.RESET, player.Health())
	fmt.Printf(x.RED+"💪 Strength: %d  "+x.RESET, player.Strength())
	fmt.Printf(x.PURPLE+"🗡️  Attack: %d\n\n"+x.RESET, player.Attack())
}
