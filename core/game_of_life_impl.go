package game_of_life

import "../board"

func (gol GameOfLife) commit() {
	gol.board = gol.nextBoard
	gol.nextBoard = *board.NewBoard(gol.board)
}

func (gol GameOfLife) compute() {
	var height = gol.height
	var width = gol.width
	var i int
	var j int

	for i = 0; i < height; i++ {
		for j = 0; j < width; j++ {
			var neighbors = gol.getNeighborCount(i, j)
			gol.computeCell(i, j, neighbors)
		}
	}
}
func (gol GameOfLife) computeCell(i int, j int, neighbors int) {
	var underpopulated = neighbors < 2
	var overpopulated = neighbors > 3
	var optimal = neighbors == 3

	if underpopulated || overpopulated {
		gol.board.Kill(i, j)
	} else if optimal {
		gol.board.Spawn(i, j)
	}
}

func (gol GameOfLife) getNeighborCount(i int, j int) int {
	var neighborsCount = 0;

	//0, height is top-left
	var top = gol.isTop(i);
	var bottom = gol.isBottom(i);
	var left = gol.isLeft(j);
	var right = gol.isRight(j);
	if !bottom && !left && gol.board.IsAlive(i-1, j-1) {
		neighborsCount++;
	}
	if !bottom && gol.board.IsAlive(i-1, j) {
		neighborsCount++;
	}
	if !bottom && !right && gol.board.IsAlive(i-1, j+1) {
		neighborsCount++;
	}
	if !left && gol.board.IsAlive(i, j-1) {
		neighborsCount++;
	}
	if !right && gol.board.IsAlive(i, j+1) {
		neighborsCount++;
	}
	if !top && !left && gol.board.IsAlive(i+1, j-1) {
		neighborsCount++;
	}
	if !top && gol.board.IsAlive(i+1, j) {
		neighborsCount++;
	}
	if !top && !right && gol.board.IsAlive(i+1, j+1) {
		neighborsCount++;
	}
	return neighborsCount;
}

func (gol GameOfLife) isTop(i int) bool {
	return i == gol.height-1;
}

func (gol GameOfLife) isBottom(i int) bool {
	return i == 0;
}

func (gol GameOfLife) isLeft(j int) bool {
	return j == 0;
}

func (gol GameOfLife) isRight(j int) bool {
	return j == gol.width-1;
}
