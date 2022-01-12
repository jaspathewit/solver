package puzzle

import (
	"bytes"
	"strconv"
)

// Cells a slice of cell
type Cells []Cell

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

//│─
// stringTop returns a string representation of the the top of the Cell
func (c Cell) stringTop() string {
	var buffer bytes.Buffer

	buffer.WriteString(" ")
	if c.Direction == DirectionNorth {
		buffer.WriteString("│")
	} else {
		buffer.WriteString(" ")
	}
	buffer.WriteString(" ")

	return buffer.String()
}

// stringMiddle returns a string representation of the the middle of the Cell
func (c Cell) stringMiddle() string {
	var buffer bytes.Buffer

	if c.Direction == DirectionWest {
		buffer.WriteString("─")
	} else {
		buffer.WriteString(" ")
	}

	buffer.WriteString(strconv.Itoa(int(c.Value)))

	if c.Direction == DirectionEast {
		buffer.WriteString("─")
	} else {
		buffer.WriteString(" ")
	}

	return buffer.String()
}

// stringBottom returns a string representation of the bottom of the Cell
func (c Cell) stringBottom() string {
	var buffer bytes.Buffer

	buffer.WriteString(" ")
	if c.Direction == DirectionSouth {
		buffer.WriteString("│")
	} else {
		buffer.WriteString(" ")
	}
	buffer.WriteString(" ")

	return buffer.String()
}
