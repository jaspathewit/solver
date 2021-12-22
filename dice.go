package main

import (
	"bytes"
	"fmt"
)

type Dice struct {
	Top    int
	North  int
	East   int
	South  int
	West   int
	Bottom int
}

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
//  N
// WTE
//  S
// ┌└┘┐│─

func (d Dice) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("┌──%d──┐\n", d.North))
	buffer.WriteString(fmt.Sprintf("%d |%d| %d\n", d.West, d.Top, d.East))
	buffer.WriteString(fmt.Sprintf("└──%d──┘\n", d.South))

	return buffer.String()
}

// Clone clones a dice creating and returning a new dice
func (d Dice) Clone() Dice {
	result := d
	return result
}

// Roll North
func (d *Dice) RollNorth() {
	d.South, d.Top, d.North, d.Bottom = d.Bottom, d.South, d.Top, d.North
}
