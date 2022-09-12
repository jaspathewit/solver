package solver

import "fmt"

// Puzzle interface implemented by puzzles
type Puzzle interface {
	fmt.Stringer
}

// Solver interface implemented by Solvers that solve a Puzzle
// of type PT
type Solver[PT Puzzle] interface {
	Solve(puzzle PT) ([]PT, []PT, error)
}
