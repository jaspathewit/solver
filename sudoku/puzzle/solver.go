package puzzle

import (
	"fmt"
	"solver/core/solver"
)

// SudokuSolver as solver for sudoku
type SudokuSolver struct{}

// Solve solves one step of the Sedoku
func (_ SudokuSolver) Solve(puzzle solver.Puzzle) (solver.Puzzles, solver.Puzzles, error) {
	// get the concrete type of the puzzle
	puz := puzzle.(*Puzzle)
	ps := make(solver.Puzzles, 0, 4)

	//fmt.Printf("All Possibles:\n")
	//for _, ref := range g.Cells {
	//	c, _ := g.Get(ref)
	//	fmt.Printf("Cell: %s\n", c)
	//}

	// eliminate all possibles
	for puz.EliminatePossibles() {
	}

	fmt.Printf("Possibles Eliminated:\n")
	//for _, ref := range g.Cells {
	//	c, _ := g.Get(ref)
	//	fmt.Printf("Cell: %s\n", c)
	//}

	// check that all cells without a value have at least 2 possibles
	if puz.ImpossibleSolution() {
		fmt.Printf("Impossible Solution\n")
		return nil, nil, nil
	}

	//fmt.Printf("Possibles Eliminated:\n")
	//for _, ref := range g.Cells {
	//	c, _ := g.Get(ref)
	//	fmt.Printf("Cell: %s\n", c)
	//}

	// check if the puzzle is solved
	if puz.Solved() {
		ps = append(ps, puz)
		return nil, ps, nil
	}

	// not solved yet
	// get the reference to the cell with the fewest possibles
	ref := puz.GetRefWithFewestPossibles()
	c, _ := puz.Get(ref)

	// get the possibles values for that cell
	possibles := c.PossibleValues()

	// loop through all the possible values
	for _, v := range possibles {
		// clone the puzzle
		pc := puz.Clone()

		// set the value on the cell
		c, _ := pc.Get(ref)
		c.SetValue(v)

		// add this puzzle to those to be solved
		ps = append(ps, pc)
	}

	// if we have any puzzles to be solved
	if len(ps) != 0 {
		return ps, nil, nil
	}

	// if the sudoku is not solved
	// no new puzzles, no solution and no errors
	return nil, nil, nil
}
