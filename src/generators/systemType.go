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
	currentProb := Generator.getSystemProbability()

	result := definitions.GetProbabilityResult(
		currentProb.RollAll(),
		currentProb.Lookup,
	)

	system.SystemType = result.(string)

	if system.SystemType == "multiple" {
		currentProb := Generator.getStarNumberProbability()
		result := definitions.GetProbabilityResult(
			currentProb.RollAll(),
			currentProb.Lookup,
		)
		system.NumberOfStars = result.(int)
	} else {
		system.NumberOfStars = 1
	}
}

func (Generator SystemTypeGenerator) getSystemProbability() definitions.Probability {
	var result definitions.Probability

	result.Dice = append(result.Dice, definitions.Dice{
		Sides: 10,
	})

	result.Lookup = map[int]interface{}{
		5:  "solitary",
		10: "multiple",
	}

	return result
}

func (Generator SystemTypeGenerator) getStarNumberProbability() definitions.Probability {
	var result definitions.Probability

	result.Dice = append(result.Dice, definitions.Dice{
		Sides: 20,
	})

	result.Lookup = map[int]interface{}{
		10: 1,
		16: 2,
		20: 3,
	}

	return result
}
