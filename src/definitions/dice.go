package definitions

import (
	"math/rand"
	"time"
)

type (
	// Dice construct
	Dice struct {
		// Sides property
		Sides int
	}
)

// Roll a preconfigured Die and returns the result
func (die *Dice) Roll() int {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	return random.Intn(die.Sides)
}
