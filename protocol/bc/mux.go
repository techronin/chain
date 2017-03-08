package bc

type Mux struct {
	body struct {
		Sources []valueSource // issuances, spends, and muxes
		Program Program
		ExtHash Hash
	}

	witness struct {
		Destinations []ValueDestination // outputs, retirements, and muxes
	}
}

func (Mux) Type() string         { return "mux1" }
func (m *Mux) Body() interface{} { return m.body }

func (Mux) Ordinal() int { return -1 }

func (m *Mux) Destinations() []ValueDestination {
	return m.witness.Destinations
}

func NewMux(sources []valueSource, program Program) *Mux {
	m := new(Mux)
	m.body.Sources = sources
	m.body.Program = program
	return m
}
