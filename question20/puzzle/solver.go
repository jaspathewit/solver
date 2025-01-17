package puzzle

import "solver/core/solver"

// Question20Solver as solver for question 20
type Question20Solver[PT solver.Puzzle] struct{}

// Solve solves one step of the question20
func (q Question20Solver[PT]) Solve(puzzle Puzzle) ([]Puzzle, []Puzzle, error) {
	ps := make([]Puzzle, 0, 4)

	// see if we can move the dice in any of the four directions
	for d := 1; d <= int(DirectionWest); d++ {
		// try and move in that direction
		p := puzzle.Clone()
		ok := p.Move(Direction(d))
		if ok {
			// we were able to move in that direction
			ps = append(ps, p)
		}
	}

	// if we moved in any of the 4 directions then the question20 is
	// not solved yet
	if len(ps) != 0 {
		return ps, nil, nil
	}

	// did we solve the question20 ?
	if puzzle.Solved() {
		ps = append(ps, puzzle)
		return nil, ps, nil
	}

	// we could not move and the question20 is not solved
	// no new puzzles, no solution and no errors
	return nil, nil, nil
}
