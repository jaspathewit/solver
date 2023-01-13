package cmd

type TopologyType string

const (
	TopologyTypeNormal = TopologyType("Normal")
	TopologyTypeSamuri = TopologyType("Samuri")
)

// Sudoku struct contains a representation of a sudoku puzzle in a particlar topology
type Sudoku struct {
	Topology TopologyType `xml:",attr"`
	Rows     []string     `xml:"Row"`
}
