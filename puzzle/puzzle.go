package puzzle

import "bytes"

type Puzzles []Puzzle

type Puzzle struct {
	Board []Cells

	Dice Dice

	diceRow int8
	diceCol int8
}

// NewPuzzel constructs a new puzzle.Puzzle
func NewPuzzle() Puzzle {
	result := Puzzle{}
	result.Dice = NewDice()

	result.diceRow = 3
	result.diceCol = 5

	result.Board = make([]Cells, 12)
	for i := range result.Board {
		result.Board[i] = make(Cells, 12)
		result.Board[i][0].Value = -1
		result.Board[i][1].Value = -1
		result.Board[i][10].Value = -1
		result.Board[i][11].Value = -1

		if i == 0 || i == 1 || i == 10 || i == 11 {
			for j := range result.Board[i] {
				result.Board[i][j].Value = -1
			}
		}
	}

	return result
}

// String representation of the Puzzle
func (p Puzzle) String() string {
	var buffer bytes.Buffer
	// for _, cells := range p.Board[2:10] {
	// 	buffer.WriteString(cells[2:10].String())
	// 	buffer.WriteString("\n")
	// }

	for _, cells := range p.Board {
		buffer.WriteString(cells.String())
		buffer.WriteString("\n")
	}

	return buffer.String()
}
