package puzzle

type Boards []Board

type Board struct {
	Count int
}

// NewBoard constructs a new puzzle.Board
func NewBoard() Board {
	return Board{}
}
