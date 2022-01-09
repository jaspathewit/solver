package solver

// Solver interface implemented by Solvers that solve a question20
type Solver interface {
	Solve(puzzle Puzzle) (Puzzles, Puzzles, error)
}

// Puzzles slice of puzzles
type Puzzles []Puzzle

// Puzzle interface implemented by puzzles
type Puzzle interface {
}

// Task contains the data needed to solve an entity
type Task struct {
	Puzzle Puzzle
	Solver Solver
}
