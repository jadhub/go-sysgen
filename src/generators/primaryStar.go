package generators

import (
	"go-sysgen/src/definitions"
)

type (
	// PrimaryStarGenerator generates the Primary Star
	PrimaryStarGenerator struct{}
)

// GetName returns the Generators Name‚
func (Generator PrimaryStarGenerator) GetName() string {
	return "PrimaryStarGenerator"
}

// Run runs the Generator
func (Generator PrimaryStarGenerator) Run(system *definitions.System) {
	var resultStar definitions.Star
	var result interface{}

	resultStar.Primary = true
	resultStar.Dwarf = false

	classProb := Generator.getStarClassProbability()
	result = definitions.GetProbabilityResult(
		classProb.RollAll(),
		classProb.Lookup,
	)
	resultStar.Class = result.(string)

	sizeProb := Generator.getStarSizeProbability()
	result = definitions.GetProbabilityResult(
		sizeProb.RollAll(),
		sizeProb.Lookup,
	)
	resultStar.Size = result.(string)

	resultStar.DecimalClassification = resultStar.GetDecimalClassification()

	resultStar.DwarfType = "-"

	system.Bodies = append(system.Bodies, resultStar)
}

func (Generator PrimaryStarGenerator) getStarClassProbability() definitions.Probability {
	var result definitions.Probability

	result.Dice = append(result.Dice, definitions.Dice{
		Sides: 10,
	}, definitions.Dice{
		Sides: 10,
	})

	result.Lookup = map[int]interface{}{
		3:  "A",
		7:  "F",
		12: "G",
		17: "K",
		20: "M",
	}

	return result
}

func (Generator PrimaryStarGenerator) getStarSizeProbability() definitions.Probability {
	var result definitions.Probability

	result.Dice = append(result.Dice, definitions.Dice{
		Sides: 10,
	}, definitions.Dice{
		Sides: 10,
	})

	result.Lookup = map[int]interface{}{
		2:  "II",
		4:  "III",
		8:  "IV",
		18: "V",
		20: "VI",
	}

	return result
}
