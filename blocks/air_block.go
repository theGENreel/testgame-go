package blocks

type AirBlock struct {
	BaseBlock
}

func NewAirBlock() *AirBlock {
	return &AirBlock{BaseBlock{symbol: " "}}
}
