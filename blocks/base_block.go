package blocks

type IBlock interface {
	Symbol() string
	Opaque() bool
}

type BaseBlock struct {
	opaque       bool
	ticking      bool
	interactable bool
	symbol       string
}

func (block BaseBlock) Symbol() string {
	return block.symbol
}

func (block BaseBlock) Opaque() bool {
	return block.opaque
}
