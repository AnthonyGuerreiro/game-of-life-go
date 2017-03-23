package display

import "fmt"
import "../board"

type ConsoleDisplay struct{}

func (display ConsoleDisplay) ClearScreen() {
	fmt.Print("\033[H\033[2J");
}

func (display ConsoleDisplay) Display(b board.Interface) () {
	var i int
	var j int
	var height = b.GetHeight()
	var width = b.GetWidth()

	print("-", width+2);
	print("\n", 1);
	for i = 0; i < height; i++ {
		print("|", 1)
		for j = 0; j < width; j++ {
			fmt.Print(getChar(b.IsAlive(i, j)))
		}
		print("|", 1)
		print("\n", 1)
	}
	print("-", width+2);
	fmt.Println()
}

func print(s string, times int) {
	var i int
	for i = 0; i < times; i++ {
		fmt.Print(s)
	}
}

func getChar(alive bool) string {
	if alive {
		return "*"
	}
	return " "
}

func (display ConsoleDisplay) DisplayStats(iterations int) {
	fmt.Println("Game lasted for ", iterations, "iterations")
}
