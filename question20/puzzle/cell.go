package puzzle

import (
	"bytes"
	"fmt"
)

// Direction of a cell (which direction did the dice go)
type Direction int8

const (
	DirectionNorth = Direction(1)
	DirectionEast  = Direction(2)
	DirectionSouth = Direction(3)
	DirectionWest  = Direction(4)
)

// Cell a location on the board
type Cell struct {
	Value     int8
	Direction Direction
}

// Cells a slice of cell
type Cells []Cell

// String representation of the slice of cells
func (cells Cells) String() string {
	var buffer bytes.Buffer

	// handle the top row
	for _, cell := range cells {
		buffer.WriteString(cell.stringTop())
	}
	buffer.WriteString("\n")
	// handle the middle row
	for _, cell := range cells {
		buffer.WriteString(cell.stringMiddle())

	}
	buffer.WriteString("\n")
	// handle the Bottom
	for _, cell := range cells {
		buffer.WriteString(cell.stringBottom())

	}
	return buffer.String()
}

// │─
// stringTop returns a string representation of the the top of the Cell
func (c Cell) stringTop() string {
	if c.Direction == DirectionNorth {
		return " | "
	}

	return "   "
}

// stringMiddle returns a string representation of the the middle of the Cell
func (c Cell) stringMiddle() string {
	if c.Direction == DirectionWest {
		return fmt.Sprintf("─%d ", c.Value)
	}

	if c.Direction == DirectionEast {
		return fmt.Sprintf(" %d─", c.Value)
	}

	return fmt.Sprintf(" %d ", c.Value)
}

// stringBottom returns a string representation of the bottom of the Cell
func (c Cell) stringBottom() string {
	if c.Direction == DirectionSouth {
		return " | "
	}
	return "   "
}
