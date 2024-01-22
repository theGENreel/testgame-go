package entities

import (
	"testestest/blocks"
	"testestest/game_map"
)

type BaseEntity struct {
	gameMap     *game_map.GameMap
	symbol      string
	x           int
	y           int
	side        string
	insideBlock blocks.IBlock
	tickCounter int
	speed       int
}

func (entity *BaseEntity) X() int {
	return entity.x
}

func (entity *BaseEntity) Y() int {
	return entity.y
}

func (entity *BaseEntity) Symbol() string {
	return entity.symbol
}

func (entity *BaseEntity) Opaque() bool {
	return false
}

func (entity *BaseEntity) Tick() {

}

func (entity *BaseEntity) MoveUp() {
	entity.side = "u"
	if entity.y > 0 {
		if !entity.gameMap.BodyLayer[entity.Y()-1][entity.X()].Opaque() {
			entity.gameMap.BodyLayer[entity.Y()][entity.X()] = entity.insideBlock
			entity.insideBlock = entity.gameMap.BodyLayer[entity.Y()-1][entity.X()]
			entity.gameMap.BodyLayer[entity.Y()-1][entity.X()] = entity
			entity.y -= 1
		}
	}
}

func (entity *BaseEntity) MoveDown() {
	entity.side = "d"
	if entity.y < entity.gameMap.Height()-1 {
		if !entity.gameMap.BodyLayer[entity.Y()+1][entity.X()].Opaque() {
			entityPointer := entity.gameMap.BodyLayer[entity.Y()][entity.X()]
			entity.gameMap.BodyLayer[entity.Y()][entity.X()] = entity.insideBlock
			entity.insideBlock = entity.gameMap.BodyLayer[entity.Y()+1][entity.X()]
			entity.y = entity.y + 1
			entity.gameMap.BodyLayer[entity.Y()][entity.X()] = entityPointer
		}
	}
	return
}

func (entity *BaseEntity) MoveLeft() {

}

func (entity *BaseEntity) MoveRight() {
	entity.side = "r"
	if entity.x < entity.gameMap.Width()-1 {
		if !entity.gameMap.BodyLayer[entity.Y()][entity.X()+1].Opaque() {
			entity.gameMap.BodyLayer[entity.Y()][entity.X()] = entity.insideBlock
			entity.insideBlock = entity.gameMap.BodyLayer[entity.Y()][entity.X()+1]
			entity.gameMap.BodyLayer[entity.Y()][entity.X()+1] = entity
			entity.x += 1
		}
	}
}
