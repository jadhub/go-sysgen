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
		// GetResult fetches the result
		GetResult(roll int, lookup map[int]interface{}) interface{}
		// GetProbability returns Probability Data
		GetProbability() Probability
	}
	// Probability structures Dice Rolling and Results
	Probability struct {
		Dice   Dice
		Lookup map[int]interface{}
	}
)
