package board

import "math/rand"
import "time"

type BoardImpl struct {
	height     int
	width      int
	cellsAlive int
	cells      [][]bool
	random     *rand.Rand
}

type Interface interface {
	GetHeight() int
	GetWidth() int
	IsAlive(i int, j int) bool
}

type Board interface {
	GetHeight() int
	GetWidth() int
	GetCellsAlive() int
	GetCells() [][]bool

	IsFinished() bool
	IsAlive(i int, j int) bool
	Spawn(i int, j int)
	Kill(i int, j int)
	SpawnRandom()
}

func (b *BoardImpl) IsFinished() bool {
	return b.cellsAlive == 0
}

func New(height int, width int, aliveProbability float64) *Board {
	var cells = createCells(height, width)
	var b = BoardImpl{height, width, 0, cells, createRand()}
	var i int
	var j int
	for i = 0; i < height; i++ {
		for j = 0; j < width; j++ {
			if b.random.Float64() < aliveProbability {
				b.Spawn(i, j)
			}
		}
	}
	var newBoard = Board(&b)
	return &newBoard
}

func NewBoard(b Board) *Board {
	var height = b.GetHeight()
	var width = b.GetWidth()
	var cellsAlive = b.GetCellsAlive()
	var cells = copyCells(height, width, b.GetCells())

	var newBoardImpl = BoardImpl{height, width, cellsAlive, cells, createRand()}
	var newBoard = Board(&newBoardImpl)
	return &newBoard
}

func createRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
}

func (b *BoardImpl) GetWidth() int {
	return b.width
}

func (b *BoardImpl) GetHeight() int {
	return b.height
}

func (b *BoardImpl) IsAlive(i int, j int) bool {
	return b.cells[i][j]
}

func (b *BoardImpl) Spawn(i int, j int) {
	var alive = b.IsAlive(i, j)
	if !alive {
		b.cells[i][j] = true
		b.cellsAlive++
	}
}

func (b *BoardImpl) Kill(i int, j int) {
	var alive = b.IsAlive(i, j)
	if alive {
		b.cells[i][j] = false
		b.cellsAlive--
	}
}

func (b *BoardImpl) SpawnRandom() {
	if b.cellsAlive == b.height*b.width {
		return
	}

	var i = b.random.Intn(b.height)
	var j = b.random.Intn(b.width)

	if !b.IsAlive(i, j) {
		b.Spawn(i, j)
	} else {
		b.SpawnRandom()
	}
}

func copyCells(height int, width int, cells [][]bool) [][]bool {
	var i int
	var j int
	var copy = createCells(height, width)
	for i = 0; i < height; i++ {
		for j = 0; j < width; j++ {
			copy[i][j] = cells[i][j]
		}
	}
	return copy
}

func (b *BoardImpl) GetCells() [][]bool {
	return b.cells
}

func (b *BoardImpl) GetCellsAlive() int {
	return b.cellsAlive
}

func createCells(height int, width int) [][]bool {
	var cells = make([][]bool, height)
	var i int
	for i = 0; i < height; i++ {
		cells[i] = make([]bool, width)
	}
	return cells
}
