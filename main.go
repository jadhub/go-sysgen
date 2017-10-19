package main

import (
	"go-system-gen/src/definitions"

	"go-system-gen/src/generators"

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
		generators.CompanionOrbitGenerator{},
	)

	for _, currentGenerator := range generatorChain.Generators {
		generatorChain.System.ApplyGenerator(currentGenerator)
	}

	fmt.Printf("%+v\n", generatorChain.System)
	fmt.Printf("%s", generatorChain.System.SystemType)
}
