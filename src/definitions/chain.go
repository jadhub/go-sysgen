package definitions

type (
	// GeneratorChain holds all Generators which will be applied
	GeneratorChain struct {
		// System is the System Object the Chain Works on
		System *System
		// Generators has all Generator Classes
		Generators []Generator
	}
	// Generator is a singular Generator
	Generator interface {
		// GetName returns the Name of the Generator
		GetName() string
		// Run runs the Generators Logic and returns a value
		Run(system *System)
	}
	// Probability structures Dice Rolling and Results
	Probability struct {
		Dice   []Dice
		Lookup map[int]interface{}
	}
)

// GetProbabilityResult looks up a given Dice Roll Result in a Lookup Map and return its value
func GetProbabilityResult(roll int, lookup map[int]interface{}) interface{} {
	var result interface{}
	lastIndex := 0

	for index, value := range lookup {

		if roll <= index && roll >= lastIndex {
			result = value
		}
		lastIndex = index + 1
	}

	return result
}

// RollAll rolls all added Dice for a Probabilty and returns the combined result
func (Probability *Probability) RollAll() int {
	var result int
	for _, Die := range Probability.Dice {
		result = result + Die.Roll()
	}

	return result
}
