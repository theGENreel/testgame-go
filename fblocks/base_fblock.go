package fblocks

type IFBlock interface {
	Symbol() string
}

type BaseFBlock struct {
	interactable bool
	symbol       string
}

func (fblock BaseFBlock) Symbol() string {
	return fblock.symbol
}
