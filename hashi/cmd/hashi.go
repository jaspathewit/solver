package cmd

type TopologyType string

const (
	TopologyTypeNormal = TopologyType("Normal")
)

// Hashi struct contains a representation of a Hashi puzzle
type Hashi struct {
	Topology TopologyType `xml:",attr"`
	Rows     []string     `xml:"Row"`
}
