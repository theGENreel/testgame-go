package main

import (
	"github.com/rthornton128/goncurses"
	"log"
	"testestest/entities"
	"testestest/game_map"
	"time"
)

func game(gameMap *game_map.GameMap, camera *game_map.Camera, player *entities.EntityPlayer, gameWindow *goncurses.Window, infoWindow *goncurses.Window) {
	ticker := time.NewTicker(50 * time.Millisecond)

	for range ticker.C {
		gameWindow.Clear()
		infoWindow.Clear()
		camera.Tick()
	}
	return
}

func newGame(window *goncurses.Window) {

	windowHeight, windowWidth := window.MaxYX()
	windowSeparator := int(float64(windowWidth) * 0.85)
	gameMap := game_map.NewGameMap(250, 250)
	gameWindow, _ := goncurses.NewWindow(windowHeight, windowSeparator, 0, 0)
	infoWindow, _ := goncurses.NewWindow(windowHeight, windowWidth-windowSeparator, 0, windowSeparator+1)
	camera := game_map.NewCamera(window, gameMap, windowWidth-(windowWidth-windowSeparator)-1, windowHeight-1, gameWindow, infoWindow)
	player := entities.NewEntityPlayer(camera)
	gameMap.BodyLayer[1][1] = player
	camera.Following = player
	game(gameMap, camera, player, gameWindow, infoWindow)
}

func mainMenu(window *goncurses.Window) {
	buttons := [2]string{"New", "Quit"}
	selection := 0
	for {
		for idx, button := range buttons {
			if idx == selection {
				window.AttrSet(goncurses.A_REVERSE)
			} else {
				window.AttrSet(goncurses.A_NORMAL)
			}
			window.MovePrint(1+idx, 1, button)
		}
		ch := window.GetChar()
		switch goncurses.KeyString(ch) {
		case "w":
			if selection > 0 {
				selection -= 1
			}
		case "s":
			if selection < len(buttons)-1 {
				selection += 1
			}
		case "enter":
			if selection == 0 {
				newGame(window)
			} else if selection == 1 {
				return
			}
		}
		window.Refresh()
	}
}

func main() {
	window, err := goncurses.Init()
	if err != nil {
		log.Fatal("init:", err)
	}
	defer goncurses.End()
	window.Keypad(true)
	window.Timeout(0)
	goncurses.Cursor(0)
	goncurses.Echo(false)
	goncurses.CBreak(false)
	goncurses.SetEscDelay(0)
	mainMenu(window)
}
