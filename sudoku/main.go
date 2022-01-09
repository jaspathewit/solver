package sudoku

import (
	"log"
	"question20/sudoku/puzzle"
)

func main() {
	g, err := puzzelLibelle()
	if err != nil {
		log.Fatalf("failed to create grid for puzzel %s", err)
	}
	//fmt.Printf("Grid \n%s", g)

	//for _, ref := range g.Cells {
	//	c, _ := g.Get(ref)
	//	fmt.Printf("Cell: %s\n", c)
	//}

	cg, err := g.Clone()
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	//for _, ref := range cg.Cells {
	//	c, _ := cg.Get(ref)
	//	fmt.Printf("Cell: %s\n", c)
	//}

	_, err = puzzle.Solve(cg, 0)
	if err != nil {
		log.Fatalf("solve: %s", err)
	}

	//fmt.Printf("Grid \n%s", g)

}

// puzzelLibelle
func puzzelLibelle() (*puzzle.Grid, error) {

	topology := puzzle.Normal{}

	g, err := puzzle.NewGrid(topology)
	if err != nil {
		return nil, err
	}

	// set up the starting values
	g.Set("1_2", "5")
	g.Set("1_3", "2")
	g.Set("1_5", "6")
	g.Set("1_6", "8")
	g.Set("1_8", "3")
	g.Set("2_2", "7")
	g.Set("2_5", "5")
	g.Set("2_7", "9")
	g.Set("2_8", "2")
	g.Set("3_2", "3")
	g.Set("3_5", "1")
	g.Set("3_9", "6")
	g.Set("5_3", "4")
	g.Set("5_6", "5")
	g.Set("5_7", "6")
	g.Set("6_3", "8")
	g.Set("6_5", "4")
	g.Set("6_7", "2")
	g.Set("7_1", "1")
	g.Set("7_2", "9")
	g.Set("7_6", "2")
	g.Set("7_8", "7")
	g.Set("8_6", "6")
	g.Set("8_9", "2")
	g.Set("9_5", "8")
	g.Set("9_7", "1")

	return g, nil
}

