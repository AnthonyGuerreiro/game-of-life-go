package display

import "../board"

type Display interface {
	ClearScreen()
	Display(b board.Interface)
	DisplayStats(iterations int)
}
