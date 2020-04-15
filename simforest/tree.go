package simforest

type Tree struct {
	Plant
}

func (t *Tree) Render() Marker {
	return Marker{
		DarkGreen,
		"T",
	}
}
