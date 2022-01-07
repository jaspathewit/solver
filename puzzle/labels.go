package puzzle

import (
	"bytes"
	"strconv"
)

// Labels 2d "array" of label
type Labels [][]int8

func NewLabels(size int) Labels {
	// create the labels for the cells for the ccl algorithum
	result := make([][]int8, size)
	for row := range result {
		result[row] = make([]int8, size)
	}

	return result
}

// String representation of Labels
func (ls Labels) String() string {
	var buffer bytes.Buffer
	for _, row := range ls {
		for _, label := range row {
			buffer.WriteString(strconv.Itoa(int(label)))
		}
		buffer.WriteString("\n")
	}

	return buffer.String()
}
