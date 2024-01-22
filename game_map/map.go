package game_map

import (
	"github.com/rthornton128/goncurses"
	"testestest/blocks"
	"testestest/fblocks"
)

type IPositioned interface {
	X() int
	Y() int
	Symbol() string
	Control(key goncurses.Key)
}

type ITicking interface {
	Tick()
}

type GameMap struct {
	width          int
	height         int
	FloorLayer     [][]fblocks.IFBlock
	BodyLayer      [][]blocks.IBlock
	tickingObjects []ITicking
}

func (gameMap GameMap) Width() int {
	return gameMap.width
}

func (gameMap GameMap) Height() int {
	return gameMap.height
}

func NewGameMap(width int, height int) *GameMap {
	gameMap := GameMap{width: width, height: height}
	for i := 0; i < height; i++ {
		var floorLine []fblocks.IFBlock
		var bodyLine []blocks.IBlock
		for j := 0; j < width; j++ {
			floorLine = append(floorLine, fblocks.NewStoneFBlock())
			bodyLine = append(bodyLine, blocks.NewAirBlock())
		}
		gameMap.FloorLayer = append(gameMap.FloorLayer, floorLine)
		gameMap.BodyLayer = append(gameMap.BodyLayer, bodyLine)
	}

	return &gameMap
}
