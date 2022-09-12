package puzzle

import (
	"fmt"
	"log"
	"strings"
)

// Contains the types and operations that can be performed on a "puzzle"
// that contains a sudoku puzzle
// a grid is in essence just a slice of cells, each cell defines three
// slices of peer cells which are it's neighbours. The surrounding grid
// row and column

// Puzzle the sudoku puzzle
type Puzzle struct {
	Topology  Topology
	cellIndex map[string]*Cell
	Cells     []string
}

// NewPuzzle create a new puzzle of the given topology
func NewPuzzle(topology Topology) (Puzzle, error) {
	result := Puzzle{Topology: topology}
	result.cellIndex = make(map[string]*Cell)
	result.Cells = make([]string, 0)

	for _, r := range topology.GridRefs() {
		c := NewCell(r)

		neighbourPeers, err := topology.NeigbourPeers().FindPeersFor(c.Ref)
		if err != nil {
			return Puzzle{}, fmt.Errorf("neighbour: %s", err)
		}
		c.NeighbourPeers = neighbourPeers

		rowPeers, err := topology.RowPeers().FindPeersFor(c.Ref)
		if err != nil {
			return Puzzle{}, fmt.Errorf("row: %s", err)
		}
		c.RowPeers = rowPeers

		colPeers, err := topology.ColPeers().FindPeersFor(c.Ref)
		if err != nil {
			return Puzzle{}, fmt.Errorf("col: %s", err)
		}
		c.ColPeers = colPeers

		// add the cell
		result.Add(c)
	}
	return result, nil
}

// Clone clones a puzzle
func (p Puzzle) Clone() Puzzle {
	result, err := NewPuzzle(p.Topology)
	if err != nil {
		log.Panic("could not clone puzzle")
	}

	// loop through the references
	for _, ref := range p.Cells {
		c, _ := p.Get(ref)
		// clone the cell
		clone := c.Clone()
		result.Add(clone)
	}

	return result
}

// Add a cell to the puzzle
func (p *Puzzle) Add(cell *Cell) {
	cell.Grid = p
	p.Cells = append(p.Cells, cell.Ref)
	p.cellIndex[cell.Ref] = cell

}

// Get function gets a cell by its reference
func (p Puzzle) Get(ref string) (*Cell, bool) {
	cell, ok := p.cellIndex[ref]
	return cell, ok
}

// Put puts a cell into the grid by reference
func (p *Puzzle) Put(cell *Cell) {
	p.cellIndex[cell.Ref] = cell
}

// Set sets a cell by its reference to the given fixed value
func (p *Puzzle) Set(ref string, value string) error {
	c, ok := p.Get(ref)
	if !ok {
		return fmt.Errorf("set: cannot find cell %s", ref)
	}

	// set the value adjusts possible values of neighbour cells
	c.SetValue(value)

	p.cellIndex[c.Ref] = c

	return nil
}

// function sets a cells label
func (p *Puzzle) SetLabel(ref string, label string) error {
	c, ok := p.Get(ref)
	if !ok {
		return fmt.Errorf("set: cannot find cell %s", ref)
	}

	c.SetLabel(label)

	p.cellIndex[c.Ref] = c

	return nil
}

// eliminate possible value for
func (p *Puzzle) EliminatePossibleValueFor(refs []string, value string) {
	// go through the cell references
	for _, ref := range refs {
		c, ok := p.Get(ref)
		if !ok {
			log.Fatalf("cannot find cell %s", ref)
		}
		// remove it as a possible value
		c.EliminatePossibleValue(value)
	}

}

// eliminate possible returns true if there was at least one cell that could be set
func (p *Puzzle) EliminatePossibles() bool {
	// go through the cells
	result := false
	for _, ref := range p.Cells {
		c, _ := p.Get(ref)
		possibleValues := c.PossibleValues()
		if len(possibleValues) == 1 {
			value := possibleValues[0]
			c.SetValue(value)
			result = true
		}
	}
	return result
}

// Solved tests if the puzzle is solved no cell has any possibles
func (p Puzzle) Solved() bool {
	// go through the cells
	for _, ref := range p.Cells {
		c, _ := p.Get(ref)
		possibleValues := c.PossibleValues()
		if len(possibleValues) != 0 {
			return false
		}
	}
	return true
}

// ImpossibleSolution tests if the grid so far is an impossible solution
func (p Puzzle) ImpossibleSolution() bool {
	// go through the cells
	for _, ref := range p.Cells {
		c, _ := p.Get(ref)
		possibleValues := c.PossibleValues()
		if c.Value() == " " && len(possibleValues) == 0 {
			return true
		}
	}
	return false
}

// GetRefWithFewestPossibles returns the cell reference with the fewest possible values
func (p Puzzle) GetRefWithFewestPossibles() string {
	// go through the cells
	result := ""
	minPossibles := 9
	for _, ref := range p.Cells {
		c, _ := p.Get(ref)
		numberOfPossibles := len(c.PossibleValues())
		if numberOfPossibles != 0 && numberOfPossibles < minPossibles {
			result = ref
			minPossibles = numberOfPossibles
		}
	}
	return result
}

// String return the puzzle as a printable string
func (p Puzzle) String() string {
	sb := strings.Builder{}

	sb.WriteString("|~~~:~~~;~~~|~~~:~~~;~~~|~~~:~~~;~~~|~~~:~~~;~~~|~~~:~~~;~~~|~~~:~~~;~~~|~~~:~~~;~~~|~~~:~~~;~~~|~~~:~~~;~~~|\n")

	for row := 1; row < p.Topology.Rows()+1; row++ {
		sb.WriteString("|")

		for col := 1; col < p.Topology.Cols()+1; col++ {
			ref := fmt.Sprintf("%d_%d", row, col)

			cell, ok := p.Get(ref)
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
