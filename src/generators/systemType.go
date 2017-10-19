package generators

import (
	"go-sysgen/src/definitions"
)

type (
	// SystemTypeGenerator is the base Generator to decide if the system is Singular, Binary, Ternary ore larger
	SystemTypeGenerator struct{}
)

// GetName returns the Generators Name‚
func (Generator SystemTypeGenerator) GetName() string {
	return "SystemTypeGenerator"
}

// Run runs the Generator
func (Generator SystemTypeGenerator) Run(system *definitions.System) {
	currentProb := Generator.GetProbability()

	result := Generator.GetResult(
		currentProb.Dice.Roll(),
		currentProb.Lookup,
	)

	system.SystemType = result.(string)
}

// GetResult fetches the result
func (Generator SystemTypeGenerator) GetResult(roll int, lookup map[int]interface{}) interface{} {
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

// GetProbability returns Probability Data
func (Generator SystemTypeGenerator) GetProbability() definitions.Probability {
	var result definitions.Probability

	result.Dice = definitions.Dice{
		Sides: 10,
	}

	result.Lookup = map[int]interface{}{
		5:  "solitary",
		10: "multiple",
	}

	return result
}
