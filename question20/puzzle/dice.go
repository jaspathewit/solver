package puzzle

import (
	"bytes"
	"fmt"
)

// Dice representation of a Dice
type Dice struct {
	Top    int8
	North  int8
	East   int8
	South  int8
	West   int8
	Bottom int8

	Row int8
	Col int8
}

// NewDice construct a new dice
func NewDice() Dice {
	result := Dice{
		Top:    6,
		Bottom: 1,
		North:  5,
		East:   4,
		South:  2,
		West:   3,
	}
	return result
}

// String prints a string representation of a dice Top and 4 sides
//
//	N
//
// WTE
//
//	S
//
// ┌└┘┐│─
func (d Dice) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("┌──%d──┐\n", d.North))
	buffer.WriteString(fmt.Sprintf("%d |%d| %d\n", d.West, d.Top, d.East))
	buffer.WriteString(fmt.Sprintf("└──%d──┘\n", d.South))

	return buffer.String()
}

// Roll rolls the dice in the given direction
// alters the faces and the dices location in the question20
// corresponding to the direction of the role
func (d Dice) Roll(direction Direction) Dice {
	switch direction {
	case DirectionNorth:
		d.South, d.Top, d.North, d.Bottom = d.Bottom, d.South, d.Top, d.North
		d.Row -= 1
	case DirectionEast:
		d.East, d.Top, d.West, d.Bottom = d.Top, d.West, d.Bottom, d.East
		d.Col += 1
	case DirectionSouth:
		d.South, d.Top, d.North, d.Bottom = d.Top, d.North, d.Bottom, d.South
		d.Row += 1
	case DirectionWest:
		d.West, d.Top, d.East, d.Bottom = d.Top, d.East, d.Bottom, d.West
		d.Col -= 1
	}
	return d
}
