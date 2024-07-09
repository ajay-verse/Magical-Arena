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
	fmt.Printf(x.YELLOW+"🧙 %s Details 🧙\n"+x.RESET, player.Name())
	fmt.Printf(x.GREEN+"🩸 Health: %d  "+x.RESET, player.Health())
	fmt.Printf(x.RED+"💪 Strength: %d  "+x.RESET, player.Strength())
	fmt.Printf(x.PURPLE+"🗡️ Attack: %d\n\n"+x.RESET, player.Attack())
}

func PrintDiceRollDetails(attacker string, attackerRoll int, defender string, defenderRoll int) {
	fmt.Printf(x.BLUE+"🎲 %s %d vs %s %d 🎲\n"+x.RESET, attacker, attackerRoll, defender, defenderRoll)
}
