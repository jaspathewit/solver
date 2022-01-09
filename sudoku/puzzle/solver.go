package puzzle

import (
	"fmt"
)

// Solve solves a sudoku passed as a grid
func Solve(g *Grid, depth int) (bool, error) {

	fmt.Printf("Solving: %d\n", depth)
	fmt.Printf("%s", g)

	if depth > 500 {
		return false, fmt.Errorf("recursion depth exceeded")
	}

	//fmt.Printf("All Possibles:\n")
	//for _, ref := range g.Cells {
	//	c, _ := g.Get(ref)
	//	fmt.Printf("Cell: %s\n", c)
	//}

	// eliminate all possibles
	for g.EliminatePossibles() {
	}

	fmt.Printf("Possibles Eliminated:\n")
	//for _, ref := range g.Cells {
	//	c, _ := g.Get(ref)
	//	fmt.Printf("Cell: %s\n", c)
	//}

	// get that all cells without a value have at least 2 possibles
	if g.ImpossibleSolution() {
		fmt.Printf("Impossible Solution\n")
		return false, nil
	}

	//fmt.Printf("Possibles Eliminated:\n")
	//for _, ref := range g.Cells {
	//	c, _ := g.Get(ref)
	//	fmt.Printf("Cell: %s\n", c)
	//}

	// check if the grid is solved
	if g.Solved() {
		fmt.Printf("Solution Found\n%s", g)
		return true, nil
	}

	// not solved yet
	// get the reference to the cell with the fewest possibles
	ref := g.GetRefWithFewestPossibles()
	c, _ := g.Get(ref)

	// get the possibles values for that cell
	possibles := c.PossibleValues()

	// loop through the possible values
	for _, v := range possibles {
		// clone the grid
		gc, err := g.Clone()
		if err != nil {
			return false, err
		}

		// set the value on the cell
		c, _ := gc.Get(ref)
		c.SetValue(v)

		// solve this grid
		fmt.Printf("Solving for %s as %s\n", ref, v)
		solved, err := Solve(gc, depth+1)
		if solved || err != nil {
			return solved, err
		}
	}

	// if we got to here then there is no solution
	fmt.Printf("No Solution Found\n%s", g)
	return false, nil
}
