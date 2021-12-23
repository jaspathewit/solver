package solver

import (
	"question20/puzzle"
)

// TimestampDateFormat date format corresponding to "yyyymmddhhmmss000"
// const TimestampDateFormat = "20060102150405000"

// Solver interface implemented by Solvers that contain entities that can be dumped (files, zip, export or activeMQ)
type Solver interface {
	Solve(board puzzle.Board) (puzzle.Boards, puzzle.Boards, error)
}

// Task contains the data needed to solve an entity
type Task struct {
	Board  puzzle.Board
	Solver Solver
}
