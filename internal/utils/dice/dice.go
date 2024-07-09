package dice

import "math/rand"

// Roll a die with a specified number of sides
func Roll(sides int) int {
	return rand.Intn(sides) + 1
}

// Roll a 6-sided die
func RollD6() int {
	return Roll(6)
}

// Roll multiple dice of a specified type
func RollMultiple(sides, count int) []int {
	rolls := make([]int, count)
	for i := 0; i < count; i++ {
		rolls[i] = Roll(sides)
	}
	return rolls
}
