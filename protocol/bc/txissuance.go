package bc

type Issuance struct {
	body struct {
		Anchor  Hash
		Value   AssetAmount
		Data    Hash
		ExtHash Hash
	}
	ordinal int

	witness struct {
		Destination         ValueDestination
		InitialBlockID      Hash
		AssetDefinitionHash Hash
		IssuanceProgram     Program
		Arguments           [][]byte
	}

	// Anchor is a pointer to the manifested entry corresponding to
	// body.Anchor.
	Anchor Entry // *nonce or *spend
}

func (Issuance) Type() string           { return "issuance1" }
func (iss *Issuance) Body() interface{} { return iss.body }

func (iss Issuance) Ordinal() int { return iss.ordinal }

func (iss *Issuance) AnchorID() Hash {
	return iss.body.Anchor
}

func (iss *Issuance) Data() Hash {
	return iss.body.Data
}

func (iss *Issuance) AssetID() AssetID {
	return iss.body.Value.AssetID
}

func (iss *Issuance) Amount() uint64 {
	return iss.body.Value.Amount
}

func (iss *Issuance) Destination() ValueDestination {
	return iss.witness.Destination
}

func (iss *Issuance) InitialBlockID() Hash {
	return iss.witness.InitialBlockID
}

func (iss *Issuance) IssuanceProgram() Program {
	return iss.witness.IssuanceProgram
}

func (iss *Issuance) Arguments() [][]byte {
	return iss.witness.Arguments
}

func (iss *Issuance) SetDestination(id Hash, value AssetAmount, pos uint64, e Entry) {
	iss.witness.Destination = ValueDestination{
		Ref:      id,
		Value:    value,
		Position: pos,
		Entry:    e,
	}
}

func (iss *Issuance) SetInitialBlockID(hash Hash) {
	iss.witness.InitialBlockID = hash
}

func (iss *Issuance) SetAssetDefinitionHash(hash Hash) {
	iss.witness.AssetDefinitionHash = hash
}

func (iss *Issuance) SetIssuanceProgram(prog Program) {
	iss.witness.IssuanceProgram = prog
}

func (iss *Issuance) SetArguments(args [][]byte) {
	iss.witness.Arguments = args
}

func NewIssuance(anchor Entry, value AssetAmount, data Hash, ordinal int) *Issuance {
	iss := new(Issuance)
	iss.body.Anchor = EntryID(anchor)
	iss.Anchor = anchor
	iss.body.Value = value
	iss.body.Data = data
	iss.ordinal = ordinal
	return iss
}
