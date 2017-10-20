package generators

import (
	"fmt"
	"go-sysgen/src/definitions"
	"strconv"
	"strings"
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
	currentProb := Generator.getCompanionProbability()

	var resultOrbitData definitions.OrbitData

	// Solitary Systems have no Companion Stars
	if system.SystemType == "solitary" {
		// Empty Object
		system.CompanionOrbits = resultOrbitData
		return
	}

	result := definitions.GetProbabilityResult(
		currentProb.RollAll(),
		currentProb.Lookup,
	)

	switch v := result.(type) {
	case int:
		// Integer Number of Orbits
		resultOrbitData.Orbits = result.(int)
		resultOrbitData.FarCompanionDistance = 0
		resultOrbitData.Type = "normal"
	case string:
		// This is either a close or a far orbit
		splitResult := strings.Split(result.(string), "|")
		numberOrbits, err := strconv.Atoi(splitResult[0])

		if err != nil {
			resultOrbitData.Orbits = 0
		}
		resultOrbitData.Orbits = numberOrbits
		resultOrbitData.Type = splitResult[1]

		die := definitions.Dice{
			Sides: 10,
		}

		switch typeResult := splitResult[1]; typeResult {
		case "close":
			resultOrbitData.FarCompanionDistance = 0
			break
		case "far":
			resultOrbitData.FarCompanionDistance = die.Roll() * 1000
			break
		default:
			resultOrbitData.FarCompanionDistance = -1
			break
		}
	default:
		fmt.Printf("unexpected type %T", v)
	}

	primaryStar := system.Bodies[0]

	// Modify Orbits according to 4a
	switch primaryStar.Size {
	case
		"Ia",
		"Ib":
		resultOrbitData.Orbits = resultOrbitData.Orbits + 9
		break
	case "II":
		resultOrbitData.Orbits = resultOrbitData.Orbits + 8
		break
	case "III":
		resultOrbitData.Orbits = resultOrbitData.Orbits + 6
		break
	}

	switch primaryStar.Class {
	case
		"M":
		resultOrbitData.Orbits = resultOrbitData.Orbits - 6
		break
	case "K":
		resultOrbitData.Orbits = resultOrbitData.Orbits - 3
		break
	}

	system.CompanionOrbits = resultOrbitData
}

func (Generator CompanionOrbitGenerator) getCompanionProbability() definitions.Probability {
	var result definitions.Probability

	die := definitions.Dice{
		Sides: 10,
	}

	result.Dice = append(result.Dice, die)

	result.Lookup = map[int]interface{}{
		1:  "1|close",
		2:  2,
		3:  die.Roll() + 2,
		4:  die.Roll() + 4,
		5:  die.Roll() + 6,
		6:  die.Roll() + 8,
		7:  die.Roll() + 10,
		10: "1|far",
	}

	return result
}
