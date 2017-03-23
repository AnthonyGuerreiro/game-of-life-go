package display

import "../board"

type NoDisplay struct{}

func (display NoDisplay) ClearScreen() {
	//nop
}
func (display NoDisplay) Display(b board.Interface) () {
	//nop
}
func (display NoDisplay) DisplayStats(iterations int) {
	//nop
}
