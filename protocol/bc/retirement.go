package bc

type Retirement struct {
	body struct {
		Source  valueSource
		Data    Hash
		ExtHash Hash
	}
	ordinal int

	// Source contains (a pointer to) the manifested entry corresponding
	// to body.Source.
	Source Entry // *issuance, *spend, or *mux
}

func (Retirement) Type() string         { return "retirement1" }
func (r *Retirement) Body() interface{} { return r.body }

func (r Retirement) Ordinal() int { return r.ordinal }

func (r *Retirement) AssetID() AssetID {
	return r.body.Source.Value.AssetID
}

func (r *Retirement) Amount() uint64 {
	return r.body.Source.Value.Amount
}

func (r *Retirement) Data() Hash {
	return r.body.Data
}

func NewRetirement(source valueSource, data Hash, ordinal int) *Retirement {
	r := new(Retirement)
	r.body.Source = source
	r.body.Data = data
	r.ordinal = ordinal
	return r
}
