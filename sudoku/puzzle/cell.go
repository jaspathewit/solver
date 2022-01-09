package puzzle

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Values a hash of possible values that a cell could have
type Values map[string]struct{}

var _ fmt.Stringer = &Cell{}

// a map of pointers to cells indexed by string
type Cells map[string]*Cell

// Cell type, defines a single cell an a sudoku grid
type Cell struct {
	// the grid in which this cell exists
	Grid *Grid

	// the cell reference
	Ref string

	// the references of a cells peers
	NeighbourPeers []string
	RowPeers       []string
	ColPeers       []string

	// Value the current value of the cell
	value string

	// the possible values that this cell could have
	possibleValues Values

	label string
}

// Create a NewCell at the given row and col
func NewCell(ref string) *Cell {

	values := make(Values)
	// set the possible values
	for i := 1; i < 10; i++ {
		s := strconv.Itoa(i)
		values[s] = struct{}{}
	}

	result := Cell{Ref: ref,
		value:          " ",
		possibleValues: values}
	return &result
}

// Clone clones a cell
func (c *Cell) Clone() *Cell {
	result := NewCell(c.Ref)

	// clone the possible values
	result.possibleValues = make(Values)
	for v, _ := range c.possibleValues {
		result.possibleValues[v] = struct{}{}
	}

	// clone the current value
	result.value = c.value
	result.label = c.label

	// clone the neighbours
	result.NeighbourPeers = c.NeighbourPeers
	result.RowPeers = c.RowPeers
	result.ColPeers = c.ColPeers

	return result
}

// Value the current value in a cell
func (c *Cell) Value() string {
	return c.value
}

// SetValue sets the current value of this cell
// removes the value as a possible from all "neghbours"
func (c *Cell) SetValue(value string) {
	c.value = value

	c.Grid.EliminatePossibleValueFor(c.NeighbourPeers, value)
	c.Grid.EliminatePossibleValueFor(c.RowPeers, value)
	c.Grid.EliminatePossibleValueFor(c.ColPeers, value)

	// we have no more possible values
	c.possibleValues = make(Values)
}

// Label the current label of a cell
func (c *Cell) Label() string {
	return c.label
}

// SetValue sets the current value of this cell
// removes the value as a possible from all "neghbours"
func (c *Cell) SetLabel(label string) {
	c.label = label
}

// EliminatePossibleValue eliminates the given possible value from this cell
func (c *Cell) EliminatePossibleValue(value string) {
	delete(c.possibleValues, value)
}

// PossibleValues returns the current possible for this cell as a []string
func (c *Cell) PossibleValues() []string {
	if len(c.possibleValues) == 0 {
		return nil
	}

	result := make([]string, 0, 4)
	for k, _ := range c.possibleValues {
		result = append(result, k)
	}

	sort.Slice(result, func(i int, j int) bool { return (result[i] < result[j]) })
	return result
}

// returns the string representation of a cell
func (c *Cell) String() string {
	posstr := strings.Join(c.PossibleValues(), ",")
	result := fmt.Sprintf("Ref: %s, Value: %s Possibles; %s", c.Ref, c.Value(), posstr)
	return result
}
