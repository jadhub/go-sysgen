package main

import (
	"go-sysgen/src/definitions"

	"go-sysgen/src/generators"

	"fmt"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func main() {
	generatorChain := definitions.GeneratorChain{
		System: new(definitions.System),
	}

	generatorChain.Generators = append(generatorChain.Generators,
		generators.SystemTypeGenerator{},
		generators.PrimaryStarGenerator{},
		generators.CompanionOrbitGenerator{},
	)

	for _, currentGenerator := range generatorChain.Generators {
		generatorChain.System.ApplyGenerator(currentGenerator)
	}

	fmt.Printf("%+v\n", generatorChain.System)
}
