package entities

import (
	"github.com/rthornton128/goncurses"
	"testestest/blocks"
	"testestest/game_map"
)

type EntityPlayer struct {
	BaseEntity
}

func NewEntityPlayer(camera *game_map.Camera) *EntityPlayer {
	return &EntityPlayer{BaseEntity{gameMap: camera.GameMap(), x: 1, y: 1, symbol: "@", insideBlock: blocks.NewAirBlock()}}
}

func (entity EntityPlayer) Control(key goncurses.Key) {
	switch key {
	case goncurses.Key(int('w')):
		entity.MoveUp()
	case goncurses.Key(int('s')):
		entity.MoveDown()
	case goncurses.Key(int('d')):
		entity.MoveRight()
	case goncurses.Key(int('a')):
		entity.MoveLeft()
	}
}
