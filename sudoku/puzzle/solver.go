package puzzle

import (
	"fmt"
	"solver/core/solver"
)

// SudokuSolver as solver for Sudokus
type SudokuSolver[PT solver.Puzzle] struct{}

// Solve solves one step of the Sudoku
func (_ SudokuSolver[PT]) Solve(puzzle Puzzle) ([]Puzzle, []Puzzle, error) {
	// get the concrete type of the puzzle
	// puz := puzzle.(*Puzzle)
	ps := make([]Puzzle, 0, 4)

	// fmt.Printf("All Possibles:\n")
	// for _, ref := range puzzle.Cells {
	// 	c, _ := puzzle.Get(ref)
	// 	fmt.Printf("Cell: %s\n", c)
	// }

	// eliminate all possibles
	for puzzle.EliminatePossibles() {
	}

	// fmt.Printf("Elim All Possibles:\n")
	// for _, ref := range puzzle.Cells {
	// 	c, _ := puzzle.Get(ref)
	// 	fmt.Printf("Cell: %s\n", c)
	// }

	// check that all cells without a value have at least 2 possibles
	// if puzzle.ImpossibleSolution() {
	// 	fmt.Printf("Impossible Solution\n")
	// 	return nil, nil, nil
	// }

	// check if the puzzle is solved
	if puzzle.Solved() {
		ps = append(ps, puzzle)
		return nil, ps, nil
	}

	// not solved yet
	// get the reference to the cell with the fewest possibles
	ref, ok := puzzle.GetRefWithFewestPossibles()
	if !ok {
		fmt.Printf("Impossible Solution\n")
		return nil, nil, nil
	}

	c, _ := puzzle.Get(ref)

	// get the possibles values for that cell
	possibles := c.PossibleValues()

	// loop through all the possible values
	for _, v := range possibles {
		// clone the puzzle
		pc := puzzle.Clone()

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
