package game_of_life

import "time"
import "../board"
import "../display"

type GameOfLife struct {
	height                 int
	width                  int
	aliveProbability       float64
	randomCellsBornPerTurn int
	interval               int
	board                  board.Board
	nextBoard              board.Board
}

type Interface interface {
	IsAlive(i int, j int) bool
	Step()
	PlayAndShow(display display.Display)
}

func New(height int, width int, aliveProbability float64, randomCellsBornPerTurn int, interval int) Interface {
	var b = board.New(height, width, aliveProbability)
	var b2 = board.NewBoard(*b)
	return GameOfLife{height, width, aliveProbability,
			  randomCellsBornPerTurn, interval,
			  *b, *b2}
}

func (gol GameOfLife) IsAlive(i int, j int) bool {
	return gol.board.IsAlive(i, j)
}

func (gol GameOfLife) Step() {
	gol.compute()
	gol.commit()
}

func (gol GameOfLife) PlayAndShow(dis display.Display) {
	var iterations = 0;
	if dis == nil {
		dis = display.NoDisplay{}
	}

	dis.ClearScreen()
	dis.Display(board.Interface(gol.board))
	for !gol.board.IsFinished() {
		gol.Step()
		dis.ClearScreen()
		dis.Display(board.Interface(gol.board))
		iterations++
		time.Sleep(time.Duration(gol.interval) * time.Millisecond)
	}
	dis.DisplayStats(iterations)
}
