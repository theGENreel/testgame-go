package game_map

import (
	"fmt"
	"github.com/rthornton128/goncurses"
)

type Camera struct {
	window     *goncurses.Window
	width      int
	height     int
	gameMap    *GameMap
	gameWindow *goncurses.Window
	infoWindow *goncurses.Window
	x          int
	y          int
	Exiting    bool
	Following  IPositioned
}

func NewCamera(window *goncurses.Window, gameMap *GameMap, width int, height int, gameWindow *goncurses.Window, infoWindow *goncurses.Window) *Camera {
	return &Camera{window: window, width: width, height: height, gameMap: gameMap, gameWindow: gameWindow, infoWindow: infoWindow, x: 0, y: 0, Exiting: false, Following: nil}
}

func (camera Camera) GameMap() *GameMap {
	return camera.gameMap
}

//func (camera Camera) SetFollowing(following IPositioned) {
//	camera.following = following
//}

func (camera Camera) draw(window *goncurses.Window) {
	camera.update()
	for line := camera.y; line < camera.height+camera.y; line++ {
		for cell := camera.x; cell < camera.width+camera.x; cell++ {
			if camera.gameMap.BodyLayer[line][cell].Symbol() == " " {
				window.Print(camera.gameMap.FloorLayer[line][cell].Symbol())
			} else {
				window.Print(camera.gameMap.BodyLayer[line][cell].Symbol())
			}
		}
		window.Print("\n")
	}
	window.Refresh()
}

func (camera Camera) drawInfo(window *goncurses.Window) {
	if camera.Following != nil {
		window.MovePrint(0, 0, fmt.Sprintf("%d:%d\n", camera.Following.Y(), camera.Following.X()))
	} else {
		window.MovePrint(0, 0, "NIL")
	}
}

func (camera Camera) update() {
	if camera.Following != nil {
		if camera.Following.X()-camera.x < int(float64(camera.width)*0.2) && camera.x > 0 {
			camera.x -= 1
		}
		if camera.Following.X()-camera.x > int(float64(camera.width)*0.8) && camera.x+camera.width < camera.gameMap.width {
			camera.x += 1
		}
		if camera.Following.Y()-camera.y < int(float64(camera.height)*0.2) && camera.y > 0 {
			camera.y -= 1
		}
		if camera.Following.Y()-camera.y > int(float64(camera.height)*0.8) && camera.y+camera.height < camera.gameMap.height {
			camera.y += 1
		}
	}
}

func (camera Camera) Tick() {
	key := camera.window.GetChar()
	if !false {
		camera.draw(camera.gameWindow)
		camera.drawInfo(camera.infoWindow)
		if camera.Following != nil {
			camera.Following.Control(key)
		}
	}
}
