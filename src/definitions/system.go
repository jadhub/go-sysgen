package definitions

type (
	// System is the main node, obiously
	System struct {
		// SystemType may be Solitary or Multiple
		SystemType    string `valid:"alpha,lowercase,in(solitary|multiple)"`
		NumberOfStars int    `valid:"numeric"`
		// CompanionOrbits lists the number of Orbits
		CompanionOrbits OrbitData
		// Bodies in a System are Stars
		Bodies []Star
	}
	// OrbitData tracks Data for additional Orbits
	OrbitData struct {
		// Orbits is the number of Orbits
		Orbits int `valid:"numeric"`
		// Type can be either close or far
		Type string `valid:"alpha,lowercase,in(close|far|normal)"`
		// FarCompanionDistance holds a distance in AU for any far Orbits
		FarCompanionDistance int `valid:"numeric"`
		/*
			NOTE:
			Any star system with small massive stars, such as a white dwarf, within 2au of the primary will be a candidate for
			periodic novas. The more massive smaller star leaches stellar material from it‘s less dense companion. As the gas
			is compressed and heated on or near the surface of the smaller star nuclear fusion of the material will occur,
			blasting the shell of gas away in a violent nova explosion.
			In such a system the close orbit planets and perhaps even the outer planets may be striped of their atmospheres.
			Life as we know it on the surface of such worlds will be impossible.
			Very close stars may be physically touching, making for a very interesting display.
		*/
	}
	// Star is a singular Star in a System, which may have one or multiple Stars
	Star struct {
		// Primary tells if this the Primary Star of this system
		Primary bool
		// Class is the Classification, e.g. A,E,F
		Class string `valid:"alpha,uppercase,length(1|1)"`
		// DecimalClassification saves the Decimal Classification After Class and before Size
		DecimalClassification int `valid:"numeric"`
		// Size categories Size in Roman Numbers
		Size string `valid:"alpha,uppercase,length(1|3)"`
		// Dwarf marks this as a Dwarf Star Class
		Dwarf bool
		// DwarfType gives the Subclassification of a Dwarf Star
		DwarfType string `valid:"alpha,length(1|2)"`
		// SpectralClass of the Star
		SpectralClass SpectralClass
	}
	// SpectralClass describes Spectral Properties, etc.
	SpectralClass struct {
		// Class is a single uppercase Letter Description for the Spectral Class
		Class string `valid:"alpha,uppercase,length(1|1)"`
		// Spectrum is a Color Description, e.g. Yellow
		Spectrum string `valid:"alpha"`
		// SurfaceTemperature is a freeform Text, e.g. 3500 to 5000 k
		SurfaceTemperature string `valid:"alphanum"`
		// Mass is the Mass in Average Sun Sizes, e.g. 1, 0.3, etc.
		Mass float32 `valid:"float"`
		// Color is a RGB hex code
		Color string `valid:"hexadecimal"`
		// Luminosity Class Info
		Luminosity LuminosityClass
	}
	// LuminosityClass Classification
	LuminosityClass struct {
		// Type Classification
		Type string `valid:"alpha,uppercase,length(1|3)"`
		// Description is a freeform Text, e.g. Luminous giants
		Description string `valid:"alphanum"`
	}
)

func (star *Star) GetDecimalClassification() int {
	decimalClassDie := Dice{}

	// d Class White Dwarfs have no Decimal Classification
	if star.Dwarf == true {
		switch star.DwarfType {
		case
			"dA",
			"dF",
			"dG":
			return -1
		}
	}

	// 5 to 9 are not possible for K and H Class of size IV
	if star.Size == "IV" {
		switch star.Class {
		case
			"K",
			"H":
			decimalClassDie.Sides = 4
			return decimalClassDie.Roll()
		}
	}

	// 0 to 4 are not possible for B,A and F Class of size VI
	if star.Size == "VI" {
		switch star.Class {
		case
			"B",
			"A",
			"F":
			decimalClassDie.Sides = 5
			return 4 + decimalClassDie.Roll()
		}
	}

	decimalClassDie.Sides = 9
	return decimalClassDie.Roll()
}

// ApplyGenerator applies a Generator the the System
func (system *System) ApplyGenerator(generator Generator) {
	generator.Run(system)
}
