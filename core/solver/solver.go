package solver

import (
	"question20/puzzle"
)

// TimestampDateFormat date format corresponding to "yyyymmddhhmmss000"
// const TimestampDateFormat = "20060102150405000"

// Solver interface implemented by Solvers that solve a puzzle
type Solver interface {
	Solve(puzzle puzzle.Puzzle) (puzzle.Puzzles, puzzle.Puzzles, error)
}

// Task contains the data needed to solve an entity
type Task struct {
	Puzzle puzzle.Puzzle
	Solver Solver
}
