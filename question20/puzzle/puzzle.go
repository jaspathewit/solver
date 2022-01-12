package puzzle

import (
	"bytes"
)

type Puzzles []Puzzle

type Puzzle struct {
	Board  []Cells
	Labels Labels

	Dice Dice
}

// NewPuzzle constructs a new question20.Puzzle
// an "array" of 12 by 12 is used with the actual
// board running from 2:10 this simplifies checking
// for values in the knight move locations
func NewPuzzle() Puzzle {
	result := Puzzle{}
	result.Dice = NewDice()

	// put the dice in the starting position
	result.Dice.Row = 3
	result.Dice.Col = 5

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

	// record the first cell value from the top face of the dice
	result.Board[result.Dice.Row][result.Dice.Col].Value = result.Dice.Top

	return result
}

// Clone constructs a clone of the question20.Puzzle
// we need to clone the question20 because copies of the
// question20 will have the same slices for the board
// Labels is not cloned that is only added to the question20
// during the CCL algorithum
func (p Puzzle) Clone() Puzzle {

	result := NewPuzzle()
	// clone the dice
	result.Dice = p.Dice

	// fast copy for slices
	for row := range result.Board {
		copy(result.Board[row], p.Board[row])
	}

	return result
}

// String representation of the Puzzle
func (p Puzzle) String() string {
	var buffer bytes.Buffer
	for _, row := range p.Board[2:10] {
		buffer.WriteString(row[2:10].String())
		buffer.WriteString("\n")
	}

	return buffer.String()
}

// Total the total of all the Cells in the Puzzle
func (p Puzzle) Total() int64 {
	var result int64
	for _, row := range p.Board[2:10] {
		for _, cell := range row[2:10] {
			result += int64(cell.Value)
		}
	}
	return result
}

// Move "moves" the dice in a question20 in the given direction
// returns the question20 representation after rolling the dice
// in the direction given, and ok if the Move in the the
// given direction was "possible".
func (p Puzzle) Move(direction Direction) (Puzzle, bool) {
	// take a clone of the given Puzzle
	c := p.Clone()

	// record that the dice is moving in the given direction from the current cell
	c.Board[c.Dice.Row][c.Dice.Col].Direction = direction

	// roll the dice in the that direction
	dice := c.Dice.Roll(direction)

	// check if the roll was valid direction
	// there was no value other than 0 recorded in the cell being moved to
	if c.Board[dice.Row][dice.Col].Value != 0 {
		return Puzzle{}, false
	}

	// record the current dice after the roll
	c.Dice = dice
	// record the value of the face of the dice
	c.Board[c.Dice.Row][c.Dice.Col].Value = c.Dice.Top

	// check if the after the roll of the dice the question20 is still valid
	// Knights move constraint has not been violated
	if !c.Valid() {
		return Puzzle{}, false
	}

	// check if after the roll of the dice the question20 is still solvable
	// the question20 has not become partitioned (2 or more areas of 0) that
	// cannot be reached by roling the dice
	if c.Partitioned() {
		return Puzzle{}, false
	}

	return c, true
}

// Solved checks if the given question20 is solved
// Dice has ended in a position a knights move away from the
// starting position and all cells have been visited
func (p Puzzle) Solved() bool {
	result := p.isDiceLocationCorrect() && p.noZeroCells()

	return result
}

// isDiceLocation checks if the dice is in one of the correct locations
func (p Puzzle) isDiceLocationCorrect() bool {
	if p.Dice.Row == 2 && p.Dice.Col == 3 {
		return true
	}

	if p.Dice.Row == 2 && p.Dice.Col == 7 {
		return true
	}

	if p.Dice.Row == 4 && p.Dice.Col == 3 {
		return true
	}

	if p.Dice.Row == 4 && p.Dice.Col == 7 {
		return true
	}

	if p.Dice.Row == 5 && p.Dice.Col == 4 {
		return true
	}

	if p.Dice.Row == 5 && p.Dice.Col == 6 {
		return true
	}

	return false
}

// noZeroCells checks if there are any zero valued cells
func (p Puzzle) noZeroCells() bool {
	for _, row := range p.Board[2:10] {
		for _, cell := range row[2:10]  {
			if cell.Value == 0 {
				return false
			}
		}
	}

	return true
}

// Valid checks if the given question20 is valid
// the current value of the cell containing the dice
// must not be the same as any cell reachable from current
// dice location via the knights move
// .......
// ..0.0..
// .0...0.
// ...X...
// .0...0.
// ..0.0..
// .......
func (p Puzzle) Valid() bool {
	value := p.Board[p.Dice.Row][p.Dice.Col].Value

	if value == p.Board[p.Dice.Row-2][p.Dice.Col-1].Value {
		return false
	}

	if value == p.Board[p.Dice.Row-2][p.Dice.Col+1].Value {
		return false
	}

	if value == p.Board[p.Dice.Row-1][p.Dice.Col+2].Value {
		return false
	}

	if value == p.Board[p.Dice.Row+1][p.Dice.Col+2].Value {
		return false
	}

	if value == p.Board[p.Dice.Row+2][p.Dice.Col+1].Value {
		return false
	}

	if value == p.Board[p.Dice.Row+2][p.Dice.Col-1].Value {
		return false
	}

	if value == p.Board[p.Dice.Row+1][p.Dice.Col-2].Value {
		return false
	}

	if value == p.Board[p.Dice.Row-1][p.Dice.Col-2].Value {
		return false
	}

	return true
}

// Partitioned determines if all the unassigned cells are still all four way connected
func (p Puzzle) Partitioned() bool {
	// create the labels for the cells and add them to the question20
	p.Labels = NewLabels(len(p.Board))
	// remove the Labels when we are done here
	defer func() { p.Labels = nil }()

	currentLabel := int8(0)
	// apply the ccl algorithum
	// only loop through the actual cells
	for rowI := range p.Board[2:10] {
		row := rowI + 2
		for colI := range p.Board[row][2:10] {
			col := colI + 2
			// check if this cell is not labeled and does not have a value
			if p.Labels[row][col] == 0 && p.Board[row][col].Value == 0 {
				// add one to the current label
				currentLabel += 1

				// check if we have found a new partition
				if currentLabel >= 2 {
					return true
				}

				// perform the Depth First Search
				p.dfs(row, col, currentLabel)
			}
		}
	}

	return false
}

// dfs perform depth first search on the question20
func (p Puzzle) dfs(row int, col int, currentLabel int8) {
	// check if Cell is not interesting or already labeled
	if p.Board[row][col].Value != 0 || p.Labels[row][col] != 0 {
		return
	}

	// mark the current cell
	p.Labels[row][col] = currentLabel

	// recursively mark the 4 neighbors
	p.dfs(row-1, col, currentLabel)
	p.dfs(row, col+1, currentLabel)
	p.dfs(row+1, col, currentLabel)
	p.dfs(row, col-1, currentLabel)
}
