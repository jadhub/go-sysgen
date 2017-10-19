package generators

import (
	"go-system-gen/src/definitions"
)

type (
	// CompanionOrbitGenerator decides how many other Stars are orbiting in the System
	CompanionOrbitGenerator struct{}
)

// GetName returns the Generators Name‚
func (Generator CompanionOrbitGenerator) GetName() string {
	return "CompanionOrbitGenerator"
}

// Run runs the Generator
func (Generator CompanionOrbitGenerator) Run(system *definitions.System) {
	currentProb := Generator.GetProbability()

	if system.SystemType == "solitary" {
		system.CompanionOrbits = 0
		return
	}

	result := Generator.GetResult(
		currentProb.Dice.Roll(),
		currentProb.Lookup,
	)

	system.CompanionOrbits = result.(int)
}

// GetResult fetches the result
func (Generator CompanionOrbitGenerator) GetResult(roll int, lookup map[int]interface{}) interface{} {
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
func (Generator CompanionOrbitGenerator) GetProbability() definitions.Probability {
	var result definitions.Probability

	result.Dice = definitions.Dice{
		Sides: 10,
	}

	result.Lookup = map[int]interface{}{
		1:  1,
		2:  2,
		3:  result.Dice.Roll() + 2,
		4:  result.Dice.Roll() + 4,
		5:  result.Dice.Roll() + 6,
		6:  result.Dice.Roll() + 8,
		7:  result.Dice.Roll() + 10,
		10: 1,
	}

	return result
}
