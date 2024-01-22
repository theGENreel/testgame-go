package fblocks

type StoneFBlock struct {
	BaseFBlock
}

func NewStoneFBlock() *StoneFBlock {
	return &StoneFBlock{BaseFBlock{symbol: "#", interactable: false}}
}
