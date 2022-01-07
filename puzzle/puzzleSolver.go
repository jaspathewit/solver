package puzzle

// Question20Solver as solver for question 20
type Question20Solver struct{}

// Solve solves one step of the puzzle
func (q Question20Solver) Solve(puzzle Puzzle) (Puzzles, Puzzles, error) {
	ps := make(Puzzles, 0, 4)

	// see if we can move the dice in any of the four directions
	for d := 1; d <= int(DirectionWest); d++ {
		// try and move in that direction
		p, ok := puzzle.Move(Direction(d))
		if ok {
			// we were able to move in that direction
			ps = append(ps, p)
		}
	}

	// if wo moved in any of the 4 directions then the puzzle is
	// not solved yet
	if len(ps) != 0 {
		return ps, nil, nil
	}

	// did we solve the puzzle ?
	if puzzle.Solved() {
		ps = append(ps, puzzle)
		return nil, ps, nil
	}

	// we could not move and the puzzle is not solved
	// no new puzzles, no solution and no errors
	return nil, nil, nil
}
