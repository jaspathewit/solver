package puzzle

import (
	"fmt"
	"log"
	"strings"
)

// Contains the types and operations that can be performed on a "grid"
// a grid is in essence just a slice of cells, each cell defines three
// slices of peer cells which are it's neighbours. The surrounding grid
// row and column

// Grid the sudoku grid
type Grid struct {
	Topology  Topology
	cellIndex map[string]*Cell
	Cells     []string
}

// create a new grid of the given topology
func NewGrid(topology Topology) (*Grid, error) {
	result := Grid{Topology: topology}
	result.cellIndex = make(map[string]*Cell)
	result.Cells = make([]string, 0)

	for _, r := range topology.GridRefs() {
		c := NewCell(r)

		neighbourPeers, err := topology.NeigbourPeers().FindPeersFor(c.Ref)
		if err != nil {
			return nil, fmt.Errorf("neighbour: %s", err)
		}
		c.NeighbourPeers = neighbourPeers

		rowPeers, err := topology.RowPeers().FindPeersFor(c.Ref)
		if err != nil {
			return nil, fmt.Errorf("row: %s", err)
		}
		c.RowPeers = rowPeers

		colPeers, err := topology.ColPeers().FindPeersFor(c.Ref)
		if err != nil {
			return nil, fmt.Errorf("col: %s", err)
		}
		c.ColPeers = colPeers

		// add the cell
		result.Add(c)
	}
	return &result, nil
}

// Clone clones a grid
func (g *Grid) Clone() (*Grid, error) {
	result, err := NewGrid(g.Topology)
	if err != nil {
		return nil, err
	}

	// loop through the references
	for _, ref := range g.Cells {
		c, _ := g.Get(ref)
		// clone the cell
		clone := c.Clone()
		result.Add(clone)
	}

	return result, nil
}

// Add a cell to the grid
func (g *Grid) Add(cell *Cell) {
	cell.Grid = g
	g.Cells = append(g.Cells, cell.Ref)
	g.cellIndex[cell.Ref] = cell

}

// Get function gets a cell by its reference
func (g *Grid) Get(ref string) (*Cell, bool) {
	cell, ok := g.cellIndex[ref]
	return cell, ok
}

// function puts a cell into the grid by reference
func (g *Grid) Put(cell *Cell) {
	g.cellIndex[cell.Ref] = cell
}

// function sets a cell by its reference to the given fixed value
func (g *Grid) Set(ref string, value string) error {
	c, ok := g.Get(ref)
	if !ok {
		return fmt.Errorf("set: cannot find cell %s", ref)
	}

	// set the value adjusts possible values of neighbour cells
	c.SetValue(value)

	g.cellIndex[c.Ref] = c

	return nil
}

// function sets a cells label
func (g *Grid) SetLabel(ref string, label string) error {
	c, ok := g.Get(ref)
	if !ok {
		return fmt.Errorf("set: cannot find cell %s", ref)
	}

	c.SetLabel(label)

	g.cellIndex[c.Ref] = c

	return nil
}

// eliminate possible value for
func (g *Grid) EliminatePossibleValueFor(refs []string, value string) {
	// go through the cell references
	for _, ref := range refs {
		c, ok := g.Get(ref)
		if !ok {
			log.Fatalf("cannot find cell %s", ref)
		}
		// remove it as a possible value
		c.EliminatePossibleValue(value)
	}

}

// eliminate possible returns true if there was at least one cell that could be set
func (g *Grid) EliminatePossibles() bool {
	// go through the cells
	result := false
	for _, ref := range g.Cells {
		c, _ := g.Get(ref)
		possibleValues := c.PossibleValues()
		if len(possibleValues) == 1 {
			value := possibleValues[0]
			c.SetValue(value)
			result = true
		}
	}
	return result
}

// Solved tests if the grid is solved no cell has any possibles
func (g *Grid) Solved() bool {
	// go through the cells
	for _, ref := range g.Cells {
		c, _ := g.Get(ref)
		possibleValues := c.PossibleValues()
		if len(possibleValues) != 0 {
			return false
		}
	}
	return true
}

// ImpossibleSolution tests if the grid so for is an impossible solution
func (g *Grid) ImpossibleSolution() bool {
	// go through the cells
	for _, ref := range g.Cells {
		c, _ := g.Get(ref)
		possibleValues := c.PossibleValues()
		if c.Value() == " " && len(possibleValues) == 0 {
			return true
		}
	}
	return false
}

// GetRefWithFewestPossibles returns the cell reference with the fewest possbile values
func (g *Grid) GetRefWithFewestPossibles() string {
	// go through the cells
	result := ""
	minPossibles := 9
	for _, ref := range g.Cells {
		c, _ := g.Get(ref)
		numberOfPossibles := len(c.PossibleValues())
		if numberOfPossibles != 0 && numberOfPossibles < minPossibles {
			result = ref
			minPossibles = numberOfPossibles
		}
	}
	return result
}

// return the grid as a printable string
func (g *Grid) String() string {
	sb := strings.Builder{}

	sb.WriteString("|~~~:~~~;~~~|~~~:~~~;~~~|~~~:~~~;~~~|~~~:~~~;~~~|~~~:~~~;~~~|~~~:~~~;~~~|~~~:~~~;~~~|~~~:~~~;~~~|~~~:~~~;~~~|\n")

	for row := 1; row < g.Topology.Rows()+1; row++ {
		sb.WriteString("|")

		for col := 1; col < g.Topology.Cols()+1; col++ {
			ref := fmt.Sprintf("%d_%d", row, col)

			cell, ok := g.Get(ref)
			if !ok {
				sb.WriteString("   ")
			} else {
				sb.WriteString(cell.Value())
				if cell.Label() == "" {
					sb.WriteString(" ")
				} else {
					sb.WriteString(cell.Label())
				}
				sb.WriteString(" ")
			}

			if col%3 == 0 {
				sb.WriteString("|")
			} else {
				sb.WriteString(":")
			}

		}
		if row%3 == 0 {
			sb.WriteString("\n|~~~:~~~;~~~|~~~:~~~;~~~|~~~:~~~;~~~|~~~:~~~;~~~|~~~:~~~;~~~|~~~:~~~;~~~|~~~:~~~;~~~|~~~:~~~;~~~|~~~:~~~;~~~|\n")
		} else {
			sb.WriteString("\n|---:---;---|---:---;---|---:---;---|---:---;---|---:---;---|---:---;---|---:---;---|---:---;---|---:---;---|\n")
		}
	}
	return sb.String()
}
