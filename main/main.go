package main

import "../core"
import "../display"

func main() {
	var gameOfLife = game_of_life.New(4, 4, 0.3, 0, 500)
	gameOfLife.PlayAndShow(display.ConsoleDisplay{})
}
