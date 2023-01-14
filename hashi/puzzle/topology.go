package puzzle

// methods implemented by a topology
type Topology interface {
	// create a new grid in this topology
	// NewGrid() (*Grid, error)
	Rows() int
	Cols() int

	// get the gridrefs for the topology
	GridRefs() []string

	// the neigbour peers
	NeigbourPeers() PeerSet
	// the Row peers
	RowPeers() PeerSet
	// the Col Peers
	ColPeers() PeerSet
}

// a type for maintaining a slice of cell references.
type Peers []string

// a type for a slice of peers
type PeerSet []Peers

// function locates the peers for a cell reference
func (ps PeerSet) FindPeersFor(ref string) (Peers, error) {
	for _, peers := range ps {
		for _, r := range peers {
			if r == ref {
				return peers, nil
			}
		}
	}

	return nil, nil //fmt.Errorf("something wrong could not find peers for %s", ref)
}
