package puzzle

type Question20Solver struct{}

func (q Question20Solver) Solve(puzzle Puzzle) (Puzzles, Puzzles, error) {
	ps := make(Puzzles, 0, 4)

	// see if we can move the dice in any of the four directions
	for d := 1; d <= int(DirectionWest); d++ {
		p, ok := puzzle.Move(Direction(d))
		if ok {
			ps = append(ps, p)
		}
	}

	// did we move in any direction
	if len(ps) != 0 {
		return ps, nil, nil
	}

	// did we solve the puzzle ?
	if puzzle.IsSolved() {
		ps = append(ps, puzzle)
		return nil, ps, nil
	}

	// we could not move and the pzzle is not solved
	return nil, nil, nil
}
