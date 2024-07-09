package helpers

import (
	// Go Internal Packages
	"fmt"

	// Local Packages
	player "magicalarena/internal/player"
	x "magicalarena/internal/utils/constants"
)

// PrintWelcomeMessage prints the welcome message at the start of the game.
func PrintWelcomeMessage() {
	fmt.Println(x.CYAN + "\n\t✨ Welcome to Magical Arena ✨\n" + x.RESET)
}

// PrintRoundName prints the name of the current round.
func PrintRoundName(round int) {
	fmt.Printf(x.RED+"\n    🏹 Attack: %d 🏹\n\n"+x.RESET, round)
}

// PrintPlayerDetails prints the details of a player.
func PrintPlayerDetails(player *player.Player) {
	fmt.Printf(x.YELLOW+"🧙 %s Details 🧙\n"+x.RESET, player.Name())
	fmt.Printf(x.GREEN+"🩸 Health: %d  "+x.RESET, player.Health())
	fmt.Printf(x.WHITE+"💪 Strength: %d  "+x.RESET, player.Strength())
	fmt.Printf(x.PURPLE+"🏹 Attack: %d\n\n"+x.RESET, player.Attack())
}

// PrintDiceRollDetails prints the details of a dice roll.
func PrintDiceRollDetails(attacker string, attackerRoll int, defender string, defenderRoll int) {
	fmt.Printf(x.BLUE+"%s 🎲 %d vs %s 🎲 %d\n"+x.RESET, attacker, attackerRoll, defender, defenderRoll)
}

// PrintAttackDetails prints the details of an attack.
func PrintAttackDetails(attackDamage int, defenceProtection int) {
	fmt.Printf(x.BLUE+"Attack🏹 %d vs Defence🛡️ %d \n"+x.RESET, attackDamage, defenceProtection)
}

// PrintHealthDetails prints the details of health after an attack.
func PrintHealthDetails(attackerHealth int, defenderHealth int) {
	fmt.Printf(x.BLUE+"Health🔋 %d vs Health🔋 %d \n\n"+x.RESET, attackerHealth, defenderHealth)
}
