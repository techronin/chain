package bc

type valueSource struct {
	Ref      Hash
	Value    AssetAmount
	Position uint64

	// The Entry corresponding to Ref, if available
	Entry
}

type ValueDestination struct {
	Ref      Hash
	Position uint64

	// The Entry corresponding to Ref, if available
	Entry
}
