package definitions

type (
	// System is the main node, obiously
	System struct {
		// SystemType may be Solitary or Multiple
		SystemType string `valid:"alpha,lowercase,in(solitary|multiple)"`
		// CompanionOrbits lists the number of Orbits
		CompanionOrbits int `valid:"numeric"`
		// Bodies in a System are Stars
		Bodies []Star
	}
	// Star is a singular Star in a System, which may have one or multiple Stars
	Star struct {
		// Class is the Classification, e.g. A,E,F
		Class string `valid:"alpha,uppercase,length(1|1)"`
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

// ApplyGenerator applies a Generator the the System
func (system *System) ApplyGenerator(generator Generator) {
	generator.Run(system)
}
